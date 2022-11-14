package client

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/ws/client/tm"
	tmsync "github.com/tendermint/tendermint/libs/sync"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

type WSClients struct {
	mtx tmsync.RWMutex

	ReqChan  chan tmtypes.RPCRequest
	StopChan chan struct{}

	//This is used to subscribe the msg from server
	Cli *tm.WSClient
}

func NewWSClient(remoteAddr, endpoint string, privKey *ecdsa.PrivateKey, pubkey string) (*WSClients, error) {
	if client, err := tm.NewWS(remoteAddr, endpoint); err != nil {
		return nil, err
	} else {
		client.PubKey = pubkey
		client.PriKey = privKey

		wsc := &WSClients{
			Cli: client,
		}
		if err := wsc.Cli.Start(); err != nil {
			return nil, err
		}
		log.Info("auth success!")
		return wsc, nil
	}
}

func (wsc *WSClients) RegisterResChannel(requestMsg chan tmtypes.RPCRequest, stopChan chan struct{}) error {
	wsc.mtx.Lock()
	defer wsc.mtx.Unlock()

	wsc.ReqChan = requestMsg
	wsc.StopChan = stopChan

	wsc.Cli.Logger.Info("register-res-channel")

	//subscribe the message from the server
	go wsc.rspListener()

	return nil
}

func (wsc *WSClients) SendMsg(rsp tmtypes.RPCResponse) error {
	if err := wsc.Cli.Send(context.Background(), rsp); err != nil {
		log.Error("send rsp failed!")
		return err
	}
	log.Info("send rsp success!")
	return nil
}

func (wsc *WSClients) rspListener() {
	ticker := time.NewTicker(100 * time.Second)
	for {
		select {
		case msg := <-wsc.Cli.RequestsCh:
			wsc.ReqChan <- msg
		case <-wsc.StopChan:
			wsc.Cli.Logger.Info("we are stopping channel")
			wsc.StopChan = nil
			return
		case <-ticker.C:
			wsc.Cli.Logger.Info("rsp goroutine is alive")
		}
	}
}
