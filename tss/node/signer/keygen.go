package signer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethc "github.com/ethereum/go-ethereum/common"
	etht "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/tss/bindings/tgm"
	tsscommon "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keygen"
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
						logger.Info().Msgf("keygen start to set group publickey for l1 contract")
						err := p.setGroupPublicKey(p.localPubKeyByte, resp.PubKeyByte)
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

func (p *Processor) setGroupPublicKey(localKey, poolPubkey []byte) error {
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

	contract, err := tgm.NewTssGroupManager(address, p.l1Client)
	if err != nil {
		p.logger.Err(err).Msg("Unable to new tss group manager contract")
		return err
	}
	groupPubKeyBytes, err := tsscommon.SetGroupPubKeyBytes(localKey, poolPubkey)
	if err != nil {
		p.logger.Err(err).Msg("failed to abi encode group public key")
		return err
	}

	opts, err := p.EstimateGas(groupPubKeyBytes, address)
	if err != nil {
		p.logger.Err(err).Msg("failed to create opts ")
		return err
	}

	tx, err := contract.SetGroupPublicKey(opts, localKey, poolPubkey)
	if err != nil {
		p.logger.Err(err).Msg("Unable to set group public key with contract")
		return err
	}
	if err := p.l1Client.SendTransaction(p.ctx, tx); err != nil {
		p.logger.Err(err).Msg("Unable to send transaction to l1 chain")
		return err
	}

	confirmTxReceipt := func(txHash ethc.Hash) *etht.Receipt {
		ctx, cancel := context.WithTimeout(context.Background(), p.confirmReceiptTimeout)
		queryTicker := time.NewTicker(p.taskInterval)
		defer func() {
			cancel()
			queryTicker.Stop()
		}()
		for {
			receipt, err := p.l1Client.TransactionReceipt(context.Background(), txHash)
			switch {
			case receipt != nil:
				txHeight := receipt.BlockNumber.Uint64()
				tipHeight, err := p.l1Client.BlockNumber(context.Background())
				if err != nil {
					log.Error("Unable to fetch block number", "err", err)
					break
				}
				log.Info("Transaction mined, checking confirmations",
					"txHash", txHash, "txHeight", txHeight,
					"tipHeight", tipHeight,
					"numConfirmations", p.l1ConfirmBlocks)
				if txHeight+uint64(p.l1ConfirmBlocks) < tipHeight {
					reverted := receipt.Status == 0
					log.Info("Transaction confirmed",
						"txHash", txHash,
						"reverted", reverted)
					// remove submitted slashing info
					return receipt
				}
			case err != nil:
				log.Error("failed to query receipt for transaction", "txHash", txHash.String())
			default:
			}
			select {
			case <-ctx.Done():
				return nil
			case <-queryTicker.C:
			}
		}
	}
	go confirmTxReceipt(tx.Hash())
	if err != nil {
		return err
	}
	return nil

}

// Ensure we can actually connect l1
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

func (p *Processor) EstimateGas(inputData []byte, toAddress ethc.Address) (*bind.TransactOpts, error) {
	header, err := p.l1Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		p.logger.Err(err).Msg("failed to get header by l1client")
		return nil, err
	}
	var gasPrice *big.Int
	var gasTipCap *big.Int
	var gasFeeCap *big.Int
	if header.BaseFee == nil {
		gasPrice, err = p.l1Client.SuggestGasPrice(context.Background())
		if err != nil {
			p.logger.Err(err).Msg("cannot fetch gas price")
			return nil, err
		}
	} else {
		gasTipCap, err = p.l1Client.SuggestGasTipCap(context.Background())
		if err != nil {
			p.logger.Err(err).Msg("failed to SuggestGasTipCap, FallbackGasTipCap = big.NewInt(1500000000) ")
			gasTipCap = big.NewInt(1500000000)
		}
		gasFeeCap = new(big.Int).Add(
			gasTipCap,
			new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
		)
	}

	msg := ethereum.CallMsg{
		From:      p.address,
		To:        &toAddress,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      inputData,
	}

	gasLimit, err := p.l1Client.EstimateGas(context.Background(), msg)
	if err != nil {
		p.logger.Err(err).Msg("failed to EstimateGas")
		return nil, err
	}
	gasLimit = uint64(float64(gasLimit) * 1.2) // add 20% buffer to prevent outOfGas error

	chainId, err := p.l1Client.ChainID(p.ctx)
	if err != nil {
		p.logger.Err(err).Msg("Unable to get chainId on l1 chain")
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, chainId)
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	nonce64, err := p.l1Client.NonceAt(p.ctx, p.address, nil)
	if err != nil {
		p.logger.Err(err).Msgf("%s unable to get current nonce",
			p.address)
		return nil, err
	}
	p.logger.Info().Msgf("Current nonce is %d", nonce64)
	nonce := new(big.Int).SetUint64(nonce64)
	opts.Nonce = nonce
	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = gasLimit
	opts.NoSend = true
	return opts, err

}
