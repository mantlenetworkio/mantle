package tm

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
	"github.com/rcrowley/go-metrics"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/libs/service"
	tmsync "github.com/tendermint/tendermint/libs/sync"
	"github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

const (
	defaultMaxReconnectAttempts = 25
	defaultWriteWait            = 0
	defaultReadWait             = 0
	defaultPingPeriod           = 0
)

// WSClient is a JSON-RPC client, which uses WebSocket for communication with
// the remote server.
//
// WSClient is safe for concurrent use by multiple goroutines.
type WSClient struct { // nolint: maligned
	conn *websocket.Conn

	Address  string // IP:PORT or /path/to/socket
	Endpoint string // /websocket/url/endpoint
	Dialer   func(string, string) (net.Conn, error)

	PubKey string
	PriKey *ecdsa.PrivateKey

	// Single user facing channel to read RPCRequest from, closed only when the
	// client is being stopped.
	RequestsCh chan types.RPCRequest

	// Callback, which will be called each time after successful reconnect.
	onReconnect func()

	// internal channels
	send            chan types.RPCResponse // user requests
	backlog         chan types.RPCResponse // stores a single user request received during a conn failure
	reconnectAfter  chan error             // reconnect requests
	readRoutineQuit chan struct{}          // a way for readRoutine to close writeRoutine

	// Maximum reconnect attempts (0 or greater; default: 25).
	maxReconnectAttempts int

	// Support both ws and wss protocols
	protocol string

	wg sync.WaitGroup

	mtx            tmsync.RWMutex
	sentLastPingAt time.Time
	reconnecting   bool
	nextReqID      int
	// sentIDs        map[types.JSONRPCIntID]bool // IDs of the requests currently in flight

	// Time allowed to write a message to the server. 0 means block until operation succeeds.
	writeWait time.Duration

	// Time allowed to read the next message from the server. 0 means block until operation succeeds.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait. If 0, no pings will be sent.
	pingPeriod time.Duration

	service.BaseService

	// Time between sending a ping and receiving a pong. See
	// https://godoc.org/github.com/rcrowley/go-metrics#Timer.
	PingPongLatencyTimer metrics.Timer
}

// NewWS returns a new client. See the commentary on the func(*WSClient)
// functions for a detailed description of how to configure ping period and
// pong wait time. The endpoint argument must begin with a `/`.
// An error is returned on invalid remote. The function panics when remote is nil.
func NewWS(remoteAddr, endpoint string, options ...func(*WSClient)) (*WSClient, error) {
	parsedURL, err := newParsedURL(remoteAddr)
	if err != nil {
		return nil, err
	}
	// default to ws protocol, unless wss is explicitly specified
	if parsedURL.Scheme != protoWSS {
		parsedURL.Scheme = protoWS
	}

	dialFn, err := makeHTTPDialer(remoteAddr)
	if err != nil {
		return nil, err
	}

	c := &WSClient{
		Address:              parsedURL.GetTrimmedHostWithPath(),
		Dialer:               dialFn,
		Endpoint:             endpoint,
		PingPongLatencyTimer: metrics.NewTimer(),

		maxReconnectAttempts: defaultMaxReconnectAttempts,
		readWait:             defaultReadWait,
		writeWait:            defaultWriteWait,
		pingPeriod:           defaultPingPeriod,
		protocol:             parsedURL.Scheme,

		// sentIDs: make(map[types.JSONRPCIntID]bool),
	}
	c.BaseService = *service.NewBaseService(nil, "WSClient", c)
	for _, option := range options {
		option(c)
	}
	return c, nil
}

// MaxReconnectAttempts sets the maximum number of reconnect attempts before returning an error.
// It should only be used in the constructor and is not Goroutine-safe.
func MaxReconnectAttempts(max int) func(*WSClient) {
	return func(c *WSClient) {
		c.maxReconnectAttempts = max
	}
}

// ReadWait sets the amount of time to wait before a websocket read times out.
// It should only be used in the constructor and is not Goroutine-safe.
func ReadWait(readWait time.Duration) func(*WSClient) {
	return func(c *WSClient) {
		c.readWait = readWait
	}
}

// WriteWait sets the amount of time to wait before a websocket write times out.
// It should only be used in the constructor and is not Goroutine-safe.
func WriteWait(writeWait time.Duration) func(*WSClient) {
	return func(c *WSClient) {
		c.writeWait = writeWait
	}
}

// PingPeriod sets the duration for sending websocket pings.
// It should only be used in the constructor - not Goroutine-safe.
func PingPeriod(pingPeriod time.Duration) func(*WSClient) {
	return func(c *WSClient) {
		c.pingPeriod = pingPeriod
	}
}

// OnReconnect sets the callback, which will be called every time after
// successful reconnect.
func OnReconnect(cb func()) func(*WSClient) {
	return func(c *WSClient) {
		c.onReconnect = cb
	}
}

// String returns WS client full address.
func (c *WSClient) String() string {
	return fmt.Sprintf("WSClient{%s (%s)}", c.Address, c.Endpoint)
}

// OnStart implements service.Service by dialing a server and creating read and
// write routines.
func (c *WSClient) OnStart() error {
	err := c.dial()
	if err != nil {
		return err
	}

	c.RequestsCh = make(chan types.RPCRequest)

	c.send = make(chan types.RPCResponse)
	// 1 additional error may come from the read/write
	// goroutine depending on which failed first.
	c.reconnectAfter = make(chan error, 1)
	// capacity for 1 request. a user won't be able to send more because the send
	// channel is unbuffered.
	c.backlog = make(chan types.RPCResponse, 1)

	c.startReadWriteRoutines()
	go c.reconnectRoutine()

	return nil
}

// Stop overrides service.Service#Stop. There is no other way to wait until Quit
// channel is closed.
func (c *WSClient) Stop() error {
	if err := c.BaseService.Stop(); err != nil {
		return err
	}
	// only close user-facing channels when we can't write to them
	c.wg.Wait()
	close(c.RequestsCh)

	return nil
}

// IsReconnecting returns true if the client is reconnecting right now.
func (c *WSClient) IsReconnecting() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.reconnecting
}

// IsActive returns true if the client is running and not reconnecting.
func (c *WSClient) IsActive() bool {
	return c.IsRunning() && !c.IsReconnecting()
}

// Send the given RPC request to the server. Results will be available on
// ResponsesCh, errors, if any, on ErrorsCh. Will block until send succeeds or
// ctx.Done is closed.
func (c *WSClient) Send(ctx context.Context, request types.RPCResponse) error {
	select {
	case c.send <- request:
		c.Logger.Info("sent a request", "reqId", request.ID)
		// c.mtx.Lock()
		// c.sentIDs[request.ID.(types.JSONRPCIntID)] = true
		// c.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Private methods

func (c *WSClient) nextRequestID() types.JSONRPCIntID {
	c.mtx.Lock()
	id := c.nextReqID
	c.nextReqID++
	c.mtx.Unlock()
	return types.JSONRPCIntID(id)
}

func (c *WSClient) dial() error {
	dialer := &websocket.Dialer{
		NetDial: c.Dialer,
		Proxy:   http.ProxyFromEnvironment,
	}
	rHeader := http.Header{}
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)

	digestBz := crypto.Keccak256Hash([]byte(timeStr)).Bytes()
	sig, err := crypto.Sign(digestBz, c.PriKey)
	if err != nil {
		return err
	}

	rHeader.Set("pubKey", c.PubKey)
	rHeader.Set("sig", hex.EncodeToString(sig))
	rHeader.Set("time", timeStr)
	conn, _, err := dialer.Dial(c.protocol+"://"+c.Address+c.Endpoint, rHeader) // nolint:bodyclose
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// reconnect tries to redial up to maxReconnectAttempts with exponential
// backoff.
func (c *WSClient) reconnect() error {
	attempt := 0

	c.mtx.Lock()
	c.reconnecting = true
	c.mtx.Unlock()
	defer func() {
		c.mtx.Lock()
		c.reconnecting = false
		c.mtx.Unlock()
	}()

	for {
		jitter := time.Duration(tmrand.Float64() * float64(time.Second)) // 1s == (1e9 ns)
		backoffDuration := jitter + ((1 << uint(attempt)) * time.Second)

		c.Logger.Info("reconnecting", "attempt", attempt+1, "backoff_duration", backoffDuration)
		time.Sleep(backoffDuration)

		err := c.dial()
		if err != nil {
			c.Logger.Error("failed to redial", "err", err)
		} else {
			c.Logger.Info("reconnected")
			if c.onReconnect != nil {
				go c.onReconnect()
			}
			return nil
		}

		attempt++

		if attempt > c.maxReconnectAttempts {
			return fmt.Errorf("reached maximum reconnect attempts: %w", err)
		}
	}
}

func (c *WSClient) startReadWriteRoutines() {
	c.wg.Add(2)
	c.readRoutineQuit = make(chan struct{})
	go c.readRoutine()
	go c.writeRoutine()
}

func (c *WSClient) processBacklog() error {
	select {
	case request := <-c.backlog:
		if c.writeWait > 0 {
			if err := c.conn.SetWriteDeadline(time.Now().Add(c.writeWait)); err != nil {
				c.Logger.Error("failed to set write deadline", "err", err)
			}
		}
		if err := c.conn.WriteJSON(request); err != nil {
			c.Logger.Error("failed to resend request", "err", err)
			c.reconnectAfter <- err
			// requeue request
			c.backlog <- request
			return err
		}
		c.Logger.Info("resend a request", "req", request)
	default:
	}
	return nil
}

func (c *WSClient) reconnectRoutine() {
	for {
		select {
		case originalError := <-c.reconnectAfter:
			// wait until writeRoutine and readRoutine finish
			c.wg.Wait()
			if err := c.reconnect(); err != nil {
				c.Logger.Error("failed to reconnect", "err", err, "original_err", originalError)
				if err = c.Stop(); err != nil {
					c.Logger.Error("failed to stop conn", "error", err)
				}

				return
			}
			// drain reconnectAfter
		LOOP:
			for {
				select {
				case <-c.reconnectAfter:
				default:
					break LOOP
				}
			}
			err := c.processBacklog()
			if err == nil {
				c.startReadWriteRoutines()
			}

		case <-c.Quit():
			return
		}
	}
}

// The client ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *WSClient) writeRoutine() {
	var ticker *time.Ticker
	if c.pingPeriod > 0 {
		// ticker with a predefined period
		ticker = time.NewTicker(c.pingPeriod)
	} else {
		// ticker that never fires
		ticker = &time.Ticker{C: make(<-chan time.Time)}
	}

	defer func() {
		ticker.Stop()
		c.conn.Close()
		// err != nil {
		// ignore error; it will trigger in tests
		// likely because it's closing an already closed connection
		// }
		c.wg.Done()
	}()

	for {
		select {
		case request := <-c.send:
			if c.writeWait > 0 {
				if err := c.conn.SetWriteDeadline(time.Now().Add(c.writeWait)); err != nil {
					c.Logger.Error("failed to set write deadline", "err", err)
				}
			}
			if err := c.conn.WriteJSON(request); err != nil {
				c.Logger.Error("failed to send request", "err", err)
				c.reconnectAfter <- err
				// add request to the backlog, so we don't lose it
				c.backlog <- request
				return
			}
		case <-ticker.C:
			if c.writeWait > 0 {
				if err := c.conn.SetWriteDeadline(time.Now().Add(c.writeWait)); err != nil {
					c.Logger.Error("failed to set write deadline", "err", err)
				}
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				c.Logger.Error("failed to write ping", "err", err)
				c.reconnectAfter <- err
				return
			}
			c.mtx.Lock()
			c.sentLastPingAt = time.Now()
			c.mtx.Unlock()
			c.Logger.Debug("sent ping")
		case <-c.readRoutineQuit:
			return
		case <-c.Quit():
			if err := c.conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			); err != nil {
				c.Logger.Error("failed to write message", "err", err)
			}
			return
		}
	}
}

// The client ensures that there is at most one reader to a connection by
// executing all reads from this goroutine.
func (c *WSClient) readRoutine() {
	defer func() {
		c.conn.Close()
		// err != nil {
		// ignore error; it will trigger in tests
		// likely because it's closing an already closed connection
		// }
		c.wg.Done()
	}()

	c.conn.SetPongHandler(func(string) error {
		// gather latency stats
		c.mtx.RLock()
		t := c.sentLastPingAt
		c.mtx.RUnlock()
		c.PingPongLatencyTimer.UpdateSince(t)

		c.Logger.Debug("got pong")
		return nil
	})

	for {
		// reset deadline for every message type (control or data)
		if c.readWait > 0 {
			if err := c.conn.SetReadDeadline(time.Now().Add(c.readWait)); err != nil {
				c.Logger.Error("failed to set read deadline", "err", err)
			}
		}
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				return
			}

			c.Logger.Error("failed to read response", "err", err)
			close(c.readRoutineQuit)
			c.reconnectAfter <- err
			return
		}

		var request types.RPCRequest
		err = json.Unmarshal(data, &request)
		if err != nil {
			c.Logger.Error("failed to parse response", "err", err, "data", string(data))
			continue
		}

		if err = validateResponseID(request.ID); err != nil {
			c.Logger.Error("error in response ID", "id", request.ID, "err", err)
			continue
		}

		c.Logger.Info("got request", "id", request.ID)

		select {
		case <-c.Quit():
		case c.RequestsCh <- request:
		}
	}
}
