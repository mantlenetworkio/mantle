package server

import (
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type IWebsocketManager interface {
	AliveNodes() []string
	RegisterResChannel(id string, responseMsg chan ResponseMsg, stopChan chan struct{}) error
	SendMsg(request RequestMsg) error
}

type ResponseMsg struct {
	RpcResponse tmtypes.RPCResponse
	SourceNode  string
}

type RequestMsg struct {
	RpcRequest tmtypes.RPCRequest
	TargetNode string
}
