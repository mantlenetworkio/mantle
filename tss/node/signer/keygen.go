package signer

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
)

const (
	setGroupPublicKeyMethodName     = "setGroupPublicKey"
	errMaxPriorityFeePerGasNotFound = "Method eth_maxPriorityFeePerGas not found"
)

var (
	FallbackGasTipCap = big.NewInt(1500000000)
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
	//new tss group manager contract parsed
	parsed, err := abi.JSON(strings.NewReader(
		tgm.TssGroupManagerABI,
	))
	if err != nil {
		p.logger.Err(err).Msg("Unable to new parsed from tss group manager contract abi")
		return err
	}
	//get tss group manager abi
	tgmABI, err := tgm.TssGroupManagerMetaData.GetAbi()
	if err != nil {
		p.logger.Err(err).Msg("Unable to get tss group manager contract abi")
		return err
	}

	rawTgmContract := bind.NewBoundContract(address, parsed, p.l1Client, p.l1Client, p.l1Client)

	dataBytes, err := tsscommon.SetGroupPubKeyBytes(localKey, poolPubkey)
	if err != nil {
		p.logger.Err(err).Msg("failed to pack set group pub key params")
		return err
	}
	setGroupPublicKeuID := tgmABI.Methods[setGroupPublicKeyMethodName].ID
	calldata := append(setGroupPublicKeuID, dataBytes...)

	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, p.chainId)
	if err != nil {
		p.logger.Err(err).Msg("Unable to new option from chainId ")
		return err

	}
	if opts.Context == nil {
		opts.Context = p.ctx
	}
	nonce64, err := p.l1Client.NonceAt(p.ctx, p.address, nil)
	if err != nil {
		p.logger.Err(err).Msgf("%s unable to get current nonce",
			p.address)
		return err
	}
	p.logger.Info().Msgf("Current nonce is %d", nonce64)
	nonce := new(big.Int).SetUint64(nonce64)
	opts.Nonce = nonce
	opts.NoSend = true
	tx, err := rawTgmContract.RawTransact(opts, calldata)
	if err != nil {
		if strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound) {
			opts.GasTipCap = FallbackGasTipCap
			tx, err = rawTgmContract.RawTransact(opts, calldata)
			if err != nil {
				p.logger.Err(err).Msg("Unable to new transaction with raw contract")
				return err
			}
		} else {
			p.logger.Err(err).Msg("Unable to new transaction with raw contract")
			return err
		}
	}

	newTx, err := p.EstimateGas(p.ctx, tx, rawTgmContract, address)

	if err != nil {
		p.logger.Err(err).Msg("got failed in estimate gas function ")
		return err
	}

	if err := p.l1Client.SendTransaction(p.ctx, newTx); err != nil {
		p.logger.Err(err).Msg("Unable to send transaction to l1 chain, need to retry ")
		for i := 0; i < 3; i++ {
			p.logger.Info().Msgf("commit transaction retry %d times", i)
			newTx, err = p.RetryTransaction(newTx, rawTgmContract)
			if err == nil {
				err = p.l1Client.SendTransaction(p.ctx, newTx)
				if err == nil {
					break
				}
			}
		}
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
	go confirmTxReceipt(newTx.Hash())
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

func (p *Processor) EstimateGas(ctx context.Context, tx *etht.Transaction, rawContract *bind.BoundContract, address ethc.Address) (*etht.Transaction, error) {
	header, err := p.l1Client.HeaderByNumber(ctx, nil)
	if err != nil {
		p.logger.Err(err).Msg("failed to get header by l1client")
		return nil, err
	}
	var gasPrice *big.Int
	var gasTipCap *big.Int
	var gasFeeCap *big.Int
	if header.BaseFee == nil {
		gasPrice, err = p.l1Client.SuggestGasPrice(ctx)
		if err != nil {
			p.logger.Err(err).Msg("cannot fetch gas price")
			return nil, err
		}
	} else {
		gasTipCap, err = p.l1Client.SuggestGasTipCap(ctx)
		if err != nil {
			p.logger.Warn().Msg("failed to SuggestGasTipCap, FallbackGasTipCap = big.NewInt(1500000000) ")
			gasTipCap = big.NewInt(1500000000)
		}
		gasFeeCap = new(big.Int).Add(
			gasTipCap,
			new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
		)
	}

	msg := ethereum.CallMsg{
		From:      p.address,
		To:        &address,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Value:     nil,
		Data:      tx.Data(),
	}

	gasLimit, err := p.l1Client.EstimateGas(ctx, msg)
	if err != nil {
		p.logger.Err(err).Msg("failed to EstimateGas")
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, p.chainId)
	if err != nil {
		p.logger.Err(err).Msg("failed to new ops in estimate gas function")
		return nil, err
	}
	opts.Context = ctx
	opts.NoSend = true
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())

	opts.GasTipCap = gasTipCap
	opts.GasFeeCap = gasFeeCap
	opts.GasLimit = 25 * gasLimit //add 20% buffer to gas limit

	return rawContract.RawTransact(opts, tx.Data())

}

func (p *Processor) RetryTransaction(tx *etht.Transaction, rawContract *bind.BoundContract) (*etht.Transaction, error) {
	p.logger.Info().Msg("start to retry commit transaction to l1chain")
	nonce64, err := p.l1Client.NonceAt(p.ctx, p.address, nil)
	if err != nil {
		p.logger.Err(err).Msgf("%s unable to get current nonce",
			p.address)
		return nil, err
	}
	p.logger.Info().Msgf("Current nonce is %d", nonce64)

	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, p.chainId)
	if err != nil {
		p.logger.Err(err).Msg("failed to new ops in estimate gas function")
		return nil, err
	}
	opts.Context = context.Background()
	opts.NoSend = true
	opts.Nonce = new(big.Int).SetUint64(nonce64)
	opts.GasTipCap = tx.GasTipCap()
	opts.GasFeeCap = tx.GasFeeCap()
	opts.GasLimit = tx.Gas()
	return rawContract.RawTransact(opts, tx.Data())

}
