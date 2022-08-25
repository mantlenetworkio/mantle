package manager

import (
	"context"
	"encoding/binary"
	"sync"
	"time"

	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/log"
	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/manager/types"
	"github.com/bitdao-io/bitnetwork/tss/slash"
	ethc "github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

const numConfirmations = 15

var (
	queryInterval = 10 * time.Second
	sendState     SendState
)

func init() {
	sendState = SendState{
		states: make(map[[28]byte]string, 0),
		lock:   &sync.Mutex{},
	}
}

func (m Manager) slashing() {
	queryTicker := time.NewTicker(queryInterval)
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
	currentTssInfo := m.tssQueryService.QueryInfo()
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

	availableNodes := m.availableNodes(currentTssInfo.PartyPubKeys)
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
	ctx, err := m.agreement(ctx, request, tss.AskSlash)
	if err != nil {
		log.Error("failed to achieve agreement to sign slashing", "address", si.Address.String(), "index", si.BatchIndex)
		return
	}
	if len(ctx.Approvers()) < ctx.TssInfos().Threshold+1 {
		log.Error("not enough available nodes to sign slashing")
		return
	}

	// store the si with the related transaction bytes
	signResp, _, err := m.sign(ctx, request, nil, tss.SignSlash)
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
	err := m.l1Cli.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Error("failed to send transaction", "err", err)
		return err
	}
	confirmTxReceipt := func(txHash ethc.Hash, info slash.SlashingInfo) *eth.Receipt {
		sendState.set(info.Address, info.BatchIndex, "has not minted")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		queryTicker := time.NewTicker(queryInterval)
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
					"numConfirmations", numConfirmations)
				sendState.set(info.Address, info.BatchIndex, "minted, wait for confirming")
				if txHeight+numConfirmations < tipHeight {
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
