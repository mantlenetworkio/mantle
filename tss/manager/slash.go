package manager

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
	"github.com/mantlenetworkio/mantle/tss/slash"
)

var sendState = SendState{
	states: make(map[[28]byte]string, 0),
	lock:   &sync.Mutex{},
}

func (m Manager) slashing() {
	queryTicker := time.NewTicker(m.taskInterval)
	for {
		signingInfos := m.store.ListSlashingInfo()
		for _, si := range signingInfos {
			m.handleSlashing(si)
		}
		select {
		case <-m.stopChan:
			return
		case <-queryTicker.C:
		}
	}
}

func (m Manager) handleSlashing(si slash.SlashingInfo) {
	log.Info("start to handleSlashing", "address", si.Address.String(), "batch_index", si.BatchIndex, "slash_type", si.SlashType, "election id", si.ElectionId)
	currentBlockNumber, err := m.l1Cli.BlockNumber(context.Background())
	if err != nil {
		log.Error("failed to query block number", "err", err)
		return
	}
	found, err := m.tssStakingSlashingCaller.GetSlashRecord(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber)}, new(big.Int).SetUint64(si.BatchIndex), si.Address)
	if err != nil {
		log.Error("failed to GetSlashRecord", "err", err)
		return
	}
	if found { // is submitted to ethereum
		found, err = m.tssStakingSlashingCaller.GetSlashRecord(&bind.CallOpts{BlockNumber: new(big.Int).SetUint64(currentBlockNumber - uint64(m.l1ConfirmBlocks))}, new(big.Int).SetUint64(si.BatchIndex), si.Address)
		if err != nil {
			log.Error("failed to GetSlashRecord", "err", err)
			return
		}
		if found { // this slashing is confirmed on ethereum
			m.store.RemoveSlashingInfo(si.Address, si.BatchIndex)
		}
		return
	}

	unJailMembers, err := m.tssGroupManagerCaller.GetTssGroupUnJailMembers(nil)
	if err != nil {
		log.Error("failed to GetTssGroupUnJailMembers", "err", err)
		return
	}
	if !isAddrExist(unJailMembers, si.Address) {
		log.Warn("can not slash the address are not unJailed", "address", si.Address.String())
		m.store.RemoveSlashingInfo(si.Address, si.BatchIndex)
		return
	}

	currentTssInfo, err := m.tssQueryService.QueryActiveInfo()
	if err != nil {
		log.Error("failed to query active tss info", "err", err)
		return
	}

	if si.ElectionId != currentTssInfo.ElectionId {
		log.Error("the election which this node supposed to be slashed is expired, ignore the slash",
			"node", si.Address.String(), "electionId", si.ElectionId, "batch index", si.BatchIndex)
		m.store.RemoveSlashingInfo(si.Address, si.BatchIndex)
		return
	}

	sendStatus := sendState.get(si.Address, si.BatchIndex)
	if sendStatus != "" {
		log.Info("the slashing is in progress", "status", sendStatus, "address", si.Address.String(), "index", si.BatchIndex)
		return
	}

	availableNodes := m.availableNodes(currentTssInfo.TssMembers)
	if len(availableNodes) < currentTssInfo.Threshold+1 {
		log.Error("not enough available nodes to sign slashing")
		return
	}

	for i, node := range availableNodes {
		address, err := tss.NodeToAddress(node)
		if err != nil {
			log.Error("wrong public key format", "err", err)
			return
		}
		if address == si.Address {
			availableNodes = append(availableNodes[:i], availableNodes[i+1:]...)
			break
		}
	}
	if len(availableNodes) < currentTssInfo.Threshold+1 {
		log.Error("not enough available nodes to sign slashing")
		return
	}
	ctx := types.NewContext().
		WithTssInfo(currentTssInfo).
		WithElectionId(currentTssInfo.ElectionId).
		WithRequestId(randomRequestId()).
		WithAvailableNodes(availableNodes)

	request := tss.SlashRequest{
		Address:    si.Address,
		BatchIndex: si.BatchIndex,
		SignType:   si.SlashType,
	}
	ctx, err = m.agreement(ctx, request, tss.AskSlash)
	if err != nil {
		log.Error("failed to achieve agreement to sign slashing", "address", si.Address.String(), "index", si.BatchIndex)
		return
	}
	if len(ctx.Approvers()) < ctx.TssInfos().Threshold+1 {
		log.Error("not enough available nodes to sign slashing")
		return
	}

	approversAddress := make([]common.Address, len(ctx.Approvers()), len(ctx.Approvers()))
	for i, node := range ctx.Approvers() {
		addr, _ := tss.NodeToAddress(node)
		approversAddress[i] = addr
	}
	digestBz, err := tss.SlashMsgHash(request.BatchIndex, request.Address, approversAddress, request.SignType)
	if err != nil {
		log.Error("failed to encode SlashMsg", "err", err)
		return
	}
	// store the si with the related transaction bytes
	signResp, _, err := m.sign(ctx, request, digestBz, tss.SignSlash)
	if err != nil {
		log.Error("failed to sign slashing", "error", err)
		return
	}

	if err = m.submitSlashing(signResp, si); err != nil {
		log.Error("failed to submit slashing transaction", "error", err)
	}
	return
}

func (m Manager) submitSlashing(signResp tss.SignResponse, si slash.SlashingInfo) error {
	tx := new(eth.Transaction)
	if err := tx.UnmarshalBinary(signResp.SlashTxBytes); err != nil {
		return err
	}
	if err := m.l1Cli.SendTransaction(context.Background(), tx); err != nil {
		log.Error("failed to send transaction", "err", err)
		return err
	}
	confirmTxReceipt := func(txHash common.Hash, info slash.SlashingInfo) *eth.Receipt {
		sendState.set(info.Address, info.BatchIndex, "has not minted")
		ctx, cancel := context.WithTimeout(context.Background(), m.confirmReceiptTimeout)
		queryTicker := time.NewTicker(m.taskInterval)
		defer func() {
			cancel()
			queryTicker.Stop()
			sendState.remove(info.Address, info.BatchIndex)
		}()
		for {
			receipt, err := m.l1Cli.TransactionReceipt(context.Background(), txHash)
			switch {
			case receipt != nil:
				txHeight := receipt.BlockNumber.Uint64()
				tipHeight, err := m.l1Cli.BlockNumber(context.Background())
				if err != nil {
					log.Error("Unable to fetch block number", "err", err)
					break
				}
				log.Info("Transaction mined, checking confirmations",
					"txHash", txHash, "txHeight", txHeight,
					"tipHeight", tipHeight,
					"numConfirmations", m.l1ConfirmBlocks)
				sendState.set(info.Address, info.BatchIndex, "minted, wait for confirming")
				if txHeight+uint64(m.l1ConfirmBlocks) <= tipHeight {
					reverted := receipt.Status == 0
					log.Info("Transaction confirmed",
						"txHash", txHash,
						"reverted", reverted)
					// remove submitted slashing info
					m.store.RemoveSlashingInfo(si.Address, si.BatchIndex)
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
	go confirmTxReceipt(tx.Hash(), si)
	return nil
}

type SendState struct {
	states map[[28]byte]string
	lock   *sync.Mutex
}

func (ss SendState) set(address common.Address, batchIndex uint64, status string) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, batchIndex)
	var key [28]byte
	copy(key[:], append(address.Hash().Bytes(), indexBz...))
	ss.states[key] = status
}

func (ss SendState) get(address common.Address, batchIndex uint64) string {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, batchIndex)
	var key [28]byte
	copy(key[:], append(address.Hash().Bytes(), indexBz...))
	return ss.states[key]
}

func (ss SendState) remove(address common.Address, batchIndex uint64) {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	indexBz := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBz, batchIndex)
	var key [28]byte
	copy(key[:], append(address.Hash().Bytes(), indexBz...))
	delete(ss.states, key)
}

func isAddrExist(set []common.Address, find common.Address) bool {
	for _, s := range set {
		if bytes.Compare(s.Bytes(), find.Bytes()) == 0 {
			return true
		}
	}
	return false
}
