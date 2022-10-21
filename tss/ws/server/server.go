package server

import (
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"net"
	"net/http"
	"os"

	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/service"
	rpcserver "github.com/tendermint/tendermint/rpc/jsonrpc/server"
)

type WSServer struct {
	service.BaseService

	Config *rpcserver.Config
	Logger log.Logger

	Listener net.Listener
	Handler  http.Handler
	WM       *WebsocketManager
}

func NewWSServer(localAddr string) (*WebsocketManager, error) {
	wsServer := &WSServer{}
	var err error

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	wsServer.Config = rpcserver.DefaultConfig()
	wsServer.Handler = http.NewServeMux()

	mux := http.NewServeMux()
	wmLogger := logger.With("protocol", "ws")
	wsServer.WM = NewWebsocketManager()

	wsServer.WM.SetWsConnOptions(OnConnect(wsServer.WM),
		OnDisconnect(func(remoteAddr, pubKey string) {
			wsServer.WM.clientDisconnected(pubKey)
		}),
	)

	wsServer.WM.SetLogger(wmLogger)
	wsServer.Logger = wmLogger
	mux.HandleFunc("/ws", wsServer.WM.WebsocketHandler)
	wsServer.Handler = mux

	wsServer.Listener, err = rpcserver.Listen(localAddr, wsServer.Config)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := rpcserver.Serve(
			wsServer.Listener,
			wsServer.Handler,
			wsServer.Logger,
			wsServer.Config,
		); err != nil {
			panic(err)
		}
	}()

	return wsServer.WM, nil
}

func OnConnect(wm *WebsocketManager) func(wsc *wsConnection) {
	return func(wsc *wsConnection) {
		wm.clientConnected(wsc.nodePublicKey, wsc.requestChan)
		go func() {
			for {
				select {
				case res := <-wsc.Output():
					wsc.Logger.Info("received response", res.String())
					recvChanMap := wm.recvChanMap
					if len(recvChanMap) > 0 {
						id := res.ID.(tmtypes.JSONRPCStringID).String()
						recvChan, ok := recvChanMap[id]
						if ok {
							recvChan <- ResponseMsg{
								RpcResponse: res,
								SourceNode:  wsc.nodePublicKey,
							}
							continue
						}
					}
					wsc.Logger.Info("unrecognized response Id", res.ID.(tmtypes.JSONRPCStringID).String())
				case <-wsc.readRoutineQuit:
					return
				}
			}
		}()
	}
}
