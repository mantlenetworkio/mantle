package server

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
	"github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

const (
	defaultWSWriteChanCapacity = 100
	defaultWSWriteWait         = 10 * time.Second
	defaultWSReadWait          = 30 * time.Second
	defaultWSPingPeriod        = (defaultWSReadWait * 9) / 10
)

// WebsocketManager provides a WS handler for incoming connections and passes a
// map of functions along with any additional params to new connections.
// NOTE: The websocket path is defined externally, e.g. in node/node.go
type WebsocketManager struct {
	websocket.Upgrader

	logger        log.Logger
	wsConnOptions []func(*wsConnection)

	recvChanMap map[string]chan ResponseMsg
	rcRWLock    *sync.RWMutex

	sendChan   map[string]chan types.RPCRequest // node -> send channel
	aliveNodes map[string]struct{}              // node -> struct{}{}
	scRWLock   *sync.RWMutex
}

// NewWebsocketManager returns a new WebsocketManager that passes a map of
// functions, connection options and logger to new WS connections.
func NewWebsocketManager(
	wsConnOptions ...func(*wsConnection),
) *WebsocketManager {
	return &WebsocketManager{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// TODO ???
				//
				// The default behaviour would be relevant to browser-based clients,
				// afaik. I suppose having a pass-through is a workaround for allowing
				// for more complex security schemes, shifting the burden of
				// AuthN/AuthZ outside the Tendermint RPC.
				// I can't think of other uses right now that would warrant a TODO
				// though. The real backstory of this TODO shall remain shrouded in
				// mystery
				return true
			},
		},
		logger:        log.NewNopLogger(),
		wsConnOptions: wsConnOptions,

		recvChanMap: make(map[string]chan ResponseMsg),
		rcRWLock:    &sync.RWMutex{},

		sendChan:   make(map[string]chan types.RPCRequest),
		aliveNodes: make(map[string]struct{}),
		scRWLock:   &sync.RWMutex{},
	}
}

func (wm *WebsocketManager) SetWsConnOptions(wsConnOptions ...func(*wsConnection)) {
	wm.wsConnOptions = wsConnOptions
}

// SetLogger sets the logger.
func (wm *WebsocketManager) SetLogger(l log.Logger) {
	wm.logger = l
}

func (wm *WebsocketManager) AliveNodes() []string {
	ret := make([]string, 0)
	for node := range wm.aliveNodes {
		ret = append(ret, node)
	}
	return ret
}

func (wm *WebsocketManager) RegisterResChannel(requestId string, recvChan chan ResponseMsg, stopChan chan struct{}) error {
	wm.rcRWLock.Lock()
	defer wm.rcRWLock.Unlock()
	wm.recvChanMap[requestId] = recvChan

	go func() {
		<-stopChan // block util stop
		wm.unregisterRecvChan(requestId)
	}()

	return nil
}

func (wm *WebsocketManager) SendMsg(msg RequestMsg) error {
	wm.scRWLock.RLock()
	defer wm.scRWLock.RUnlock()
	sendChan, ok := wm.sendChan[msg.TargetNode]
	if !ok {
		return errors.New(fmt.Sprintf("the node(%s) is lost", msg.TargetNode))
	}
	go func() {
		sendChan <- msg.RpcRequest
	}()
	return nil
}

func (wm *WebsocketManager) unregisterRecvChan(requestId string) {
	wm.rcRWLock.Lock()
	defer wm.rcRWLock.Unlock()
	delete(wm.recvChanMap, requestId)
}

func (wm *WebsocketManager) clientConnected(pubkey string, channel chan types.RPCRequest) {
	wm.scRWLock.Lock()
	defer wm.scRWLock.Unlock()
	wm.sendChan[pubkey] = channel
	if wm.aliveNodes == nil {
		wm.aliveNodes = make(map[string]struct{})
	}
	wm.aliveNodes[pubkey] = struct{}{}
	wm.logger.Info("new node connected", "public key", pubkey)
}

func (wm *WebsocketManager) clientDisconnected(pubkey string) {
	wm.scRWLock.Lock()
	defer wm.scRWLock.Unlock()

	delete(wm.aliveNodes, pubkey)
	delete(wm.sendChan, pubkey)
	wm.logger.Info("node disconnected", "public key", pubkey)
}

// WebsocketHandler upgrades the request/response (via http.Hijack) and starts
// the wsConnection.
func (wm *WebsocketManager) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := wm.Upgrade(w, r, nil)
	if err != nil {
		// TODO - return http error
		wm.logger.Error("Failed to upgrade connection", "err", err)
		return
	}
	defer func() {
		if err := wsConn.Close(); err != nil {
			wm.logger.Error("Failed to close connection", "err", err)
		}
	}()

	pubKey := r.Header.Get("pubKey")
	timeStr := r.Header.Get("time")
	sig := r.Header.Get("sig")

	if len(pubKey) == 0 || len(timeStr) == 0 || len(sig) == 0 {
		wm.logger.Error("Failed to establish connection", "err", errors.New("invalid header"))
		return
	}

	timeInt64, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil || time.Now().Unix()-timeInt64 > 5 {
		wm.logger.Error("illegal timestamp", "err", err)
		return
	}

	pubKeyBytes, pubErr := hex.DecodeString(pubKey)
	sigBytes, sigErr := hex.DecodeString(sig)
	if pubErr != nil || sigErr != nil {
		wm.logger.Error("hex decode error for pubkey or sig", "err", err)
		return
	}
	digestBz := crypto.Keccak256Hash([]byte(timeStr)).Bytes()
	if !crypto.VerifySignature(pubKeyBytes, digestBz, sigBytes[:64]) {
		wm.logger.Error("illegal signature", "publicKey", pubKey, "time", timeStr, "signature", sig)
		return
	}

	// register connection
	con := newWSConnection(wsConn, pubKey, wm.wsConnOptions...)
	con.SetLogger(wm.logger.With("remote", wsConn.RemoteAddr()))
	wm.logger.Info("New websocket connection", "remote", con.remoteAddr)

	err = con.Start() // BLOCKING
	if err != nil {
		wm.logger.Error("Failed to start connection", "err", err)
		return
	}
	if err := con.Stop(); err != nil {
		wm.logger.Error("error while stopping connection", "error", err)
	}

}

// WebSocket connection

// A single websocket connection contains listener id, underlying ws
// connection, and the event switch for subscribing to events.
//
// In case of an error, the connection is stopped.
type wsConnection struct {
	service.BaseService

	remoteAddr string
	baseConn   *websocket.Conn

	nodePublicKey string
	responseChan  chan types.RPCResponse
	requestChan   chan types.RPCRequest

	// chan, which is closed when/if readRoutine errors
	// used to abort writeRoutine
	readRoutineQuit chan struct{}

	// write channel capacity
	writeChanCapacity int

	// each write times out after this.
	writeWait time.Duration

	// Connection times out if we haven't received *anything* in this long, not even pings.
	readWait time.Duration

	// Send pings to server with this period. Must be less than readWait, but greater than zero.
	pingPeriod time.Duration

	// Maximum message size.
	readLimit int64

	// callback which is called upon disconnect
	onDisconnect func(remoteAddr, pubKey string)

	ctx    context.Context
	cancel context.CancelFunc
}

// NewWSConnection wraps websocket.Conn.
//
// See the commentary on the func(*wsConnection) functions for a detailed
// description of how to configure ping period and pong wait time. NOTE: if the
// write buffer is full, pongs may be dropped, which may cause clients to
// disconnect. see https://github.com/gorilla/websocket/issues/97
func newWSConnection(
	baseConn *websocket.Conn,
	publicKey string,
	options ...func(*wsConnection),
) *wsConnection {
	wsc := &wsConnection{
		remoteAddr:        baseConn.RemoteAddr().String(),
		baseConn:          baseConn,
		nodePublicKey:     publicKey,
		writeWait:         defaultWSWriteWait,
		writeChanCapacity: defaultWSWriteChanCapacity,
		readWait:          defaultWSReadWait,
		pingPeriod:        defaultWSPingPeriod,
		readRoutineQuit:   make(chan struct{}),
	}
	wsc.responseChan = make(chan types.RPCResponse, wsc.writeChanCapacity)
	wsc.requestChan = make(chan types.RPCRequest, wsc.writeChanCapacity)
	for _, option := range options {
		option(wsc)
	}
	wsc.baseConn.SetReadLimit(wsc.readLimit)
	wsc.BaseService = *service.NewBaseService(nil, "wsConnection", wsc)
	return wsc
}

// OnDisconnect sets a callback which is used upon disconnect - not
// Goroutine-safe. Nop by default.
func OnDisconnect(onDisconnect func(remoteAddr, pubKey string)) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.onDisconnect = onDisconnect
	}
}

// WriteWait sets the amount of time to wait before a websocket write times out.
// It should only be used in the constructor - not Goroutine-safe.
func WriteWait(writeWait time.Duration) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.writeWait = writeWait
	}
}

// WriteChanCapacity sets the capacity of the websocket write channel.
// It should only be used in the constructor - not Goroutine-safe.
func WriteChanCapacity(cap int) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.writeChanCapacity = cap
	}
}

// ReadWait sets the amount of time to wait before a websocket read times out.
// It should only be used in the constructor - not Goroutine-safe.
func ReadWait(readWait time.Duration) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.readWait = readWait
	}
}

// PingPeriod sets the duration for sending websocket pings.
// It should only be used in the constructor - not Goroutine-safe.
func PingPeriod(pingPeriod time.Duration) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.pingPeriod = pingPeriod
	}
}

// ReadLimit sets the maximum size for reading message.
// It should only be used in the constructor - not Goroutine-safe.
func ReadLimit(readLimit int64) func(*wsConnection) {
	return func(wsc *wsConnection) {
		wsc.readLimit = readLimit
	}
}

// OnStart implements service.Service by starting the read and write routines. It
// blocks until there's some error.
func (wsc *wsConnection) OnStart() error {

	// Read subscriptions/unsubscriptions to events
	go wsc.readRoutine()
	// Write responses, BLOCKING.
	wsc.writeRoutine()

	return nil
}

// OnStop implements service.Service by unsubscribing remoteAddr from all
// subscriptions.
func (wsc *wsConnection) OnStop() {
	if wsc.onDisconnect != nil {
		wsc.onDisconnect(wsc.remoteAddr, wsc.nodePublicKey)
	}

	if wsc.ctx != nil {
		wsc.cancel()
	}
}

func (wsc *wsConnection) Output() chan types.RPCResponse {
	return wsc.responseChan
}

// GetRemoteAddr returns the remote address of the underlying connection.
// It implements WSRPCConnection
func (wsc *wsConnection) GetRemoteAddr() string {
	return wsc.remoteAddr
}

// WriteRPCResponse pushes a response to the writeChan, and blocks until it is
// accepted.
// It implements WSRPCConnection. It is Goroutine-safe.
func (wsc *wsConnection) WriteRPCResponse(ctx context.Context, resp types.RPCResponse) error {
	select {
	case <-wsc.Quit():
		return errors.New("connection was stopped")
	case <-ctx.Done():
		return ctx.Err()
	case wsc.responseChan <- resp:
		return nil
	}
}

// TryWriteRPCResponse attempts to push a response to the writeChan, but does
// not block.
// It implements WSRPCConnection. It is Goroutine-safe
func (wsc *wsConnection) TryWriteRPCResponse(resp types.RPCResponse) bool {
	select {
	case <-wsc.Quit():
		return false
	case wsc.responseChan <- resp:
		return true
	default:
		return false
	}
}

// Context returns the connection's context.
// The context is canceled when the client's connection closes.
func (wsc *wsConnection) Context() context.Context {
	if wsc.ctx != nil {
		return wsc.ctx
	}
	wsc.ctx, wsc.cancel = context.WithCancel(context.Background())
	return wsc.ctx
}

// Read from the socket and subscribe to or unsubscribe from events
func (wsc *wsConnection) readRoutine() {
	// readRoutine will block until response is written or WS connection is closed
	writeCtx := context.Background()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("WSJSONRPC: %v", r)
			}
			wsc.Logger.Error("Panic in WSJSONRPC handler", "err", err, "stack", string(debug.Stack()))
			if err := wsc.WriteRPCResponse(writeCtx, types.RPCInternalError(types.JSONRPCIntID(-1), err)); err != nil {
				wsc.Logger.Error("Error writing RPC response", "err", err)
			}
			go wsc.readRoutine()
		}
	}()

	wsc.baseConn.SetPongHandler(func(m string) error {
		return wsc.baseConn.SetReadDeadline(time.Now().Add(wsc.readWait))
	})

	for {
		select {
		case <-wsc.Quit():
			return
		default:
			// reset deadline for every type of message (control or data)
			if err := wsc.baseConn.SetReadDeadline(time.Now().Add(wsc.readWait)); err != nil {
				wsc.Logger.Error("failed to set read deadline", "err", err)
			}

			_, r, err := wsc.baseConn.NextReader()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					wsc.Logger.Info("Client closed the connection")
				} else {
					wsc.Logger.Error("Failed to read request", "err", err)
				}
				if err := wsc.Stop(); err != nil {
					wsc.Logger.Error("Error closing websocket connection", "err", err)
				}
				close(wsc.readRoutineQuit)
				return
			}

			dec := json.NewDecoder(r)
			var response types.RPCResponse
			err = dec.Decode(&response)
			if err != nil {
				wsc.Logger.Error("error unmarshaling response", "err", err)
				continue
			}

			if response.ID == nil {
				wsc.Logger.Info("[WS]received response with no ID, drop it")
				continue
			}
			if err := wsc.WriteRPCResponse(writeCtx, response); err != nil {
				wsc.Logger.Error("Error writing RPC response", "err", err)
			}

		}
	}
}

// receives on a write channel and writes out on the socket
func (wsc *wsConnection) writeRoutine() {
	pingTicker := time.NewTicker(wsc.pingPeriod)
	defer pingTicker.Stop()

	// https://github.com/gorilla/websocket/issues/97
	pongs := make(chan string, 1)
	wsc.baseConn.SetPingHandler(func(m string) error {
		select {
		case pongs <- m:
		default:
		}
		return nil
	})

	for {
		select {
		case <-wsc.Quit():
			return
		case <-wsc.readRoutineQuit: // error in readRoutine
			return
		case m := <-pongs:
			err := wsc.writeMessageWithDeadline(websocket.PongMessage, []byte(m))
			if err != nil {
				wsc.Logger.Info("Failed to write pong (client may disconnect)", "err", err)
			}
		case <-pingTicker.C:
			err := wsc.writeMessageWithDeadline(websocket.PingMessage, []byte{})
			if err != nil {
				wsc.Logger.Error("Failed to write ping", "err", err)
				return
			}
		case msg := <-wsc.requestChan:
			wsc.Logger.Info("send msg from requestChan to target client", "method", msg.Method)
			jsonBytes, err := json.MarshalIndent(msg, "", "  ")
			if err != nil {
				wsc.Logger.Error("Failed to marshal RPCRequest to JSON", "err", err)
				continue
			}
			if err = wsc.writeMessageWithDeadline(websocket.TextMessage, jsonBytes); err != nil {
				wsc.Logger.Error("Failed to write request", "err", err, "msg", msg)
				return
			}
		}
	}
}

// All writes to the websocket must (re)set the write deadline.
// If some writes don't set it while others do, they may timeout incorrectly
// (https://github.com/tendermint/tendermint/issues/553)
func (wsc *wsConnection) writeMessageWithDeadline(msgType int, msg []byte) error {
	if err := wsc.baseConn.SetWriteDeadline(time.Now().Add(wsc.writeWait)); err != nil {
		return err
	}
	return wsc.baseConn.WriteMessage(msgType, msg)
}
