package signer

import (
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rs/zerolog"

	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/mt-tss/bindings/tsh"
	tsscommon "github.com/mantlenetworkio/mantle/mt-tss/common"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
)

const (
	slashingMethodName = "slashing"
)

func (p *Processor) SignSlash() {
	defer p.wg.Done()
	logger := p.logger.With().Str("step", "sign Slash Message").Logger()

	logger.Info().Msg("start to sign Slash message ")

	go func() {
		defer func() {
			logger.Info().Msg("exit sign process")
		}()
		for {
			select {
			case <-p.stopChan:
				return
			case req := <-p.signSlashChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				logger.Info().Msgf("dealing resId (%s) ", resId)

				var nodeSignRequest tsscommon.NodeSignRequest
				rawMsg := json.RawMessage{}
				nodeSignRequest.RequestBody = &rawMsg

				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					logger.Error().Msg("failed to unmarshal node sign request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				var requestBody tsscommon.SlashRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					logger.Error().Msg("failed to umarshal slash params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}
				nodeSignRequest.RequestBody = requestBody

				err := p.checkSlashMessages(requestBody)
				if err != nil {
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())

					p.wsClient.SendMsg(RpcResponse)
					logger.Err(err).Msg("check event failed")
					continue
				}

				nodesaddrs := make([]ethc.Address, len(nodeSignRequest.Nodes))
				for i, node := range nodeSignRequest.Nodes {
					addr, _ := tsscommon.NodeToAddress(node)
					nodesaddrs[i] = addr
				}
				hashTx, err := tsscommon.SlashMsgHash(requestBody.BatchIndex, requestBody.Address, nodesaddrs, requestBody.SignType)
				mesTx, err := tsscommon.SlashMsgBytes(requestBody.BatchIndex, requestBody.Address, nodesaddrs, requestBody.SignType)
				if err != nil {
					logger.Err(err).Msg("failed to encode SlashMsg")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					p.wsClient.SendMsg(RpcResponse)
					continue
				}

				data, culprits, err := p.handleSign(nodeSignRequest, hashTx, logger)

				if err != nil {
					logger.Error().Msgf("slash %s sign failed ", requestBody.Address)
					var errorRes tdtypes.RPCResponse
					if len(culprits) > 0 {
						respData := strings.Join(culprits, ",")
						errorRes = tdtypes.NewRPCErrorResponse(req.ID, 100, err.Error(), respData)
						p.nodeStore.AddCulprits(culprits)
					} else {
						errorRes = tdtypes.NewRPCErrorResponse(req.ID, 201, "sign failed", err.Error())
					}
					er := p.wsClient.SendMsg(errorRes)
					if er != nil {
						logger.Err(er).Msg("failed to send msg to tss manager")
					} else {
						p.removeWaitSlashMsg(requestBody)
					}
					continue
				}
				txData, gasPrice, err := p.txBuilder(mesTx, data, logger)
				if err != nil {
					logger.Err(err).Msg("failed to txbuilder slash tranction")
					errorRes := tdtypes.NewRPCErrorResponse(req.ID, 201, "sign failed", err.Error())
					p.wsClient.SendMsg(errorRes)
					continue
				}

				signResponse := tsscommon.SignResponse{
					Signature:       data,
					SlashTxBytes:    txData,
					SlashTxGasPrice: gasPrice.String(),
				}

				RpcResponse := tdtypes.NewRPCSuccessResponse(req.ID, signResponse)
				err = p.wsClient.SendMsg(RpcResponse)
				if err != nil {
					logger.Err(err).Msg("failed to sendMsg to bridge ")
				} else {
					logger.Info().Msg("send slash sign response successfully")
					p.removeWaitSlashMsg(requestBody)
				}
			}
		}
	}()
}

func (p *Processor) checkSlashMessages(sign tsscommon.SlashRequest) error {
	p.waitSignSlashLock.RLock()
	defer p.waitSignSlashLock.RUnlock()
	v, ok := p.waitSignSlashMsgs[sign.Address.String()]
	if !ok {
		return errors.New("slash sign request has not been verified")
	}
	_, ok = v[sign.BatchIndex]
	if !ok {
		return errors.New("slash sign request has not been verified")
	}

	return nil
}
func (p *Processor) removeWaitSlashMsg(msg tsscommon.SlashRequest) {
	p.waitSignSlashLock.Lock()
	defer p.waitSignSlashLock.Unlock()
	v, ok := p.waitSignSlashMsgs[msg.Address.String()]
	if ok {
		_, sok := v[msg.BatchIndex]
		if sok {
			delete(v, msg.BatchIndex)
		}
		if len(v) == 0 {
			delete(p.waitSignSlashMsgs, msg.Address.String())
		}
	}
}

func (p *Processor) txBuilder(txData, sig []byte, logger zerolog.Logger) ([]byte, *big.Int, error) {
	logger.Info().Msg("connecting to layer one")
	if err := ensureConnection(p.l1Client); err != nil {
		logger.Err(err).Msg("Unable to connect to layer one")
		return nil, nil, err
	}
	if len(p.tssStakingSlashingAddress) == 0 {
		logger.Error().Msg("tss staking slashing address is empty ")
		return nil, nil, errors.New("tss staking slashing address is empty")
	}
	address := ethc.HexToAddress(p.tssStakingSlashingAddress)

	//new raw contract
	parsed, err := abi.JSON(strings.NewReader(tsh.TssStakingSlashingABI))
	if err != nil {
		logger.Err(err).Msg("Unable to new parsed from slash contract abi")
		return nil, nil, err
	}
	//get staking slash contract abi
	tshABI, err := tsh.TssStakingSlashingMetaData.GetAbi()
	if err != nil {
		logger.Err(err).Msg("Unable to get tss staking slashing ABI")
		return nil, nil, err
	}

	rawSlashContract := bind.NewBoundContract(address, parsed, p.l1Client, p.l1Client, p.l1Client)
	dataBytes, err := tsscommon.SlashBytes(txData, sig)
	if err != nil {
		logger.Err(err).Msg("failed to pack slash bytes")
		return nil, nil, err
	}
	slashingID := tshABI.Methods[slashingMethodName].ID
	calldata := append(slashingID, dataBytes...)

	opts, err := bind.NewKeyedTransactorWithChainID(p.privateKey, p.chainId)
	if err != nil {
		p.logger.Err(err).Msg("failed to new keyed transactor")
		return nil, nil, err
	}
	if opts.Context == nil {
		opts.Context = p.ctx
	}

	nonce64, err := p.l1Client.NonceAt(p.ctx, p.address, nil)
	if err != nil {
		p.logger.Err(err).Msgf("%s unable to get current nonce",
			p.address)
		return nil, nil, err
	}
	p.logger.Info().Msgf("Current nonce is %d", nonce64)
	nonce := new(big.Int).SetUint64(nonce64)
	opts.Nonce = nonce
	opts.NoSend = true

	tx, err := rawSlashContract.RawTransact(opts, calldata)
	if err != nil {
		if strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound) {
			opts.GasTipCap = FallbackGasTipCap
			tx, err = rawSlashContract.RawTransact(opts, calldata)
			if err != nil {
				logger.Err(err).Msg("failed to build slashing transaction tx!")
				return nil, nil, err
			}
		} else {
			logger.Err(err).Msg("failed to build slashing transaction tx!")
			return nil, nil, err
		}
	}

	newTx, err := p.EstimateGas(p.ctx, tx, rawSlashContract, address)

	txBinary, err := newTx.MarshalBinary()
	if err != nil {
		logger.Err(err).Msg("failed to get marshal binary from transaction tx")
		return nil, nil, err
	}
	return txBinary, newTx.GasPrice(), nil
}
