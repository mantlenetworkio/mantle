package signer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/bitdao-io/bitnetwork/tss/bindings/tgm"
	tsscommon "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/common"
	"github.com/bitdao-io/bitnetwork/tss/node/tsslib/keygen"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	etht "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"time"
)

func (p *Processor) Keygen() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "keygen").Logger()

	logger.Info().Msg("start to keygen ")

	go func() {
		defer func() {
			logger.Info().Msg("exit keygen process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.keygenRequestChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Msgf("dealing resId (%s) ", resId)

				var keyR tsscommon.KeygenRequest
				if err := json.Unmarshal(req.Params, &keyR); err != nil {
					logger.Error().Msg("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				var keygenReq = keygen.Request{
					Keys:      keyR.Nodes,
					ThresHold: keyR.Threshold,
				}
				resp, err := p.tssServer.Keygen(keygenReq)

				if err != nil {
					logger.Err(err).Msg("failed to keygen !")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 202, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
				} else {
					if resp.Status == common.Success {
						keygenResponse := tsscommon.KeygenResponse{
							ClusterPublicKey: resp.PubKey,
						}
						RpcResponse := tdtypes.NewRPCSuccessResponse(tdtypes.JSONRPCStringID(resId), keygenResponse)
						p.wsClient.SendMsg(RpcResponse)
						err := p.setGroupPublicKey(p.localPubkey, resp.PubKey)
						if err != nil {
							logger.Err(err).Msg("failed to send tss group manager transactionx")
						}
					} else {
						RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 202, "failed", resp.FailReason)
						p.wsClient.SendMsg(RpcResponse)
					}
				}

			}
		}
	}()
}

func (p *Processor) setGroupPublicKey(localKey, poolPubkey string) error {
	p.logger.Info().Msg("connecting to layer one")
	if err := ensureConnection(p.l1Client); err != nil {
		p.logger.Err(err).Msg("Unable to connect to layer one")
		return err
	}
	if len(p.tssGroupManagerAddress) == 0 {
		p.logger.Error().Msg("tss group manager address is empty ")
		return errors.New("tss group manager address is empty")
	}
	address := ethc.HexToAddress(p.tssGroupManagerAddress)
	chainId, err := p.l1Client.ChainID(p.ctx)
	if err != nil {
		return err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, chainId)
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	opts.NoSend = true
	contract, err := tgm.NewTssGroupManager(address, p.l1Client)
	if err != nil {
		return err
	}
	gasPrice, err := p.l1Client.SuggestGasPrice(context.Background())
	if err != nil {
		p.logger.Err(err).Msg("cannot fetch gas price")
		return err
	}
	opts.GasPrice = gasPrice
	tx, err := contract.SetGroupPublicKey(opts, []byte(localKey), []byte(poolPubkey))
	if err != nil {
		return err
	}
	if err := p.l1Client.SendTransaction(p.ctx, tx); err != nil {
		return err
	}

	receipt, err := waitForReceipt(p.l1Client, tx)
	if err != nil {
		return err
	}
	p.logger.Info().Msgf("tss group tss transaction confirmed , hash %s, gas-used %s,block number  %s", tx.Hash().Hex(), receipt.GasUsed, receipt.BlockNumber)
	return nil

}

//Ensure we can actually connect l1
func ensureConnection(client *ethclient.Client) error {
	t := time.NewTicker(1 * time.Second)
	retries := 0
	defer t.Stop()
	for ; true; <-t.C {
		_, err := client.ChainID(context.Background())
		if err == nil {
			break
		} else {
			retries += 1
			if retries > 90 {
				return err
			}
		}
	}
	return nil
}

// Wait for the receipt by polling the backend
func waitForReceipt(backend *ethclient.Client, tx *etht.Transaction) (*etht.Receipt, error) {
	t := time.NewTicker(300 * time.Millisecond)
	receipt := new(etht.Receipt)
	var err error
	for range t.C {
		receipt, err = backend.TransactionReceipt(context.Background(), tx.Hash())
		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if receipt != nil {
			t.Stop()
			break
		}
	}
	return receipt, nil
}
