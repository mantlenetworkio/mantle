package eth

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/core"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/p2p"
	"github.com/mantlenetworkio/mantle/l2geth/p2p/enode"
)

func (pm *ProtocolManager) makeConsensusProtocol(version uint) p2p.Protocol {
	length := consensusProtocolLength

	return p2p.Protocol{
		Name:    consensusProtocolName,
		Version: version,
		Length:  length,
		Run:     pm.consensusHandler,
		NodeInfo: func() interface{} {
			return pm.NodeInfo()
		},
		PeerInfo: func(id enode.ID) interface{} {
			if p := pm.peers.Peer(fmt.Sprintf("%x", id[:8])); p != nil {
				return p.Info()
			}
			return nil
		},
	}
}

func (pm *ProtocolManager) removePeerTmp(id string) {
	// Short circuit if the peer was already removed
	peer := pm.consensusPeers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing Ethereum consensus peer", "peer", id)

	if err := pm.consensusPeers.Unregister(id); err != nil {
		log.Error("Consensus Peer removal failed", "peer", id, "err", err)
	}
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (pm *ProtocolManager) consensusHandler(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	p := pm.newPeer(int(eth64), peer, rw)
	select {
	case pm.newPeerCh <- p:
		pm.wg.Add(1)
		defer pm.wg.Done()
		// Ignore maxPeers if this is a trusted peer
		if pm.consensusPeers.Len() >= pm.maxPeers && !p.Peer.Info().Network.Trusted {
			return p2p.DiscTooManyPeers
		}
		p.Log().Debug("Ethereum consensus peer connected", "name", p.Name())
		if rw, ok := p.rw.(*meteredMsgReadWriter); ok {
			rw.Init(p.version)
		}
		// Register the peer locally
		if err := pm.consensusPeers.Register(p); err != nil {
			p.Log().Error("Ethereum peer registration failed", "err", err)
			return err
		}
		defer pm.removePeerTmp(p.id)

		// Handle incoming messages until the connection is torn down
		for {
			if err := pm.checkPeer(p); err != nil {
				return err
			}

			if err := pm.handleConsensusMsg(p); err != nil {
				p.Log().Debug("Ethereum consensus message handling failed", "err", err)
				return err
			}
		}
	case <-pm.quitSync:
		return p2p.DiscQuitting
	}
}

func (pm *ProtocolManager) checkPeer(p *peer) error {
	if !pm.schedulerInst.IsRunning() {
		return nil
	}
	if bytes.Equal(pm.etherbase.Bytes(), pm.schedulerInst.Scheduler().Bytes()) {
		has := make(chan bool)
		if err := pm.eventMux.Post(core.PeerAddEvent{
			PeerId: p.ID().Bytes(),
			Has:    has,
		}); err != nil {
			return err
		}

		p.Log().Debug("wait for peer id check", "ID", p.ID().String())
		select {
		case find := <-has:
			if !find {
				p.Log().Debug("Have not find peer ", "ID", p.ID().String())
				return errors.New("have not find peer")
			} else {
				p.Log().Debug("find peer ", "ID", p.ID().String())
			}
		}
	}
	return nil
}

func (pm *ProtocolManager) handleConsensusMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > consensusMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on its contents
	switch {
	case msg.Code == BatchPeriodStartMsg:
		var bs *types.BatchPeriodStartMsg
		if err := msg.Decode(&bs); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		log.Info("Batch Period Start RollbackStates", "batch_index", bs.BatchIndex, "start_height", bs.StartHeight, "max_height", bs.MaxHeight, "expire_time", bs.ExpireTime)
		if !types.VerifySigner(bs, pm.schedulerInst.Scheduler()) {
			return nil
		}

		p.knowBatchPeriodStartMsg.Add(bs.BatchIndex)
		erCh := make(chan error, 1)
		pm.eventMux.Post(core.BatchPeriodStartEvent{
			Msg:   bs,
			ErrCh: erCh,
		})
	case msg.Code == BatchPeriodAnswerMsg:
		var bpa *types.BatchPeriodAnswerMsg
		if err := msg.Decode(&bpa); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		log.Info("Batch Period Answer RollbackStates", "batchIndex", bpa.BatchIndex, "start_index", bpa.StartIndex, "tx_len", len(bpa.Txs))
		if !pm.schedulerInst.IsRunning() {
			log.Debug("not scheduler")
			return nil
		}
		if !types.VerifySigner(bpa, pm.schedulerInst.CurrentStartMsg().Sequencer) {
			return nil
		}
		p.knowBatchPeriodAnswerMsg.Add(bpa.Hash())
		erCh := make(chan error, 1)
		pm.eventMux.Post(core.BatchPeriodAnswerEvent{
			Msg:   bpa,
			ErrCh: erCh,
		})
	case msg.Code == RollbackMsg:
		var rm *types.RollbackMsg
		if err := msg.Decode(&rm); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		log.Info("Fraud Proof Reorg RollbackStates")
		p.knowRollbackMsg.Add(rm.Hash())
		erCh := make(chan error, 1)
		pm.eventMux.Post(core.RollbackStartEvent{
			Msg:   rm,
			ErrCh: erCh,
		})

	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

// ---------------------------- Consensus Control Messages ----------------------------

// BatchPeriodStartMsg will
func (pm *ProtocolManager) batchPeriodStartMsgBroadcastLoop() {
	log.Info("Start batchPeriodStartMsg broadcast routine")
	// automatically stops if unsubscribe
	for obj := range pm.batchStartMsgSub.Chan() {
		if se, ok := obj.Data.(core.BatchPeriodStartEvent); ok {
			log.Debug("Got BatchPeriodStartEvent, broadcast it",
				"batch_index", se.Msg.BatchIndex,
				"start_height", se.Msg.StartHeight,
				"max_height", se.Msg.MaxHeight)
			pm.BroadcastBatchPeriodStartMsg(se.Msg) // First propagate block to peers
		}
	}
}

func (pm *ProtocolManager) BroadcastBatchPeriodStartMsg(msg *types.BatchPeriodStartMsg) {
	peers := pm.consensusPeers.PeersWithoutStartMsg(msg.BatchIndex)
	for _, p := range peers {
		p.AsyncSendBatchPeriodStartMsg(msg)
	}
	log.Trace("Broadcast batch period start msg")
}

func (p *peer) AsyncSendBatchPeriodStartMsg(msg *types.BatchPeriodStartMsg) {
	select {
	case p.queuedBatchStartMsg <- msg:
		p.knowBatchPeriodStartMsg.Add(msg.BatchIndex)
		for p.knowBatchPeriodStartMsg.Cardinality() >= maxKnownStartMsg {
			p.knowBatchPeriodStartMsg.Pop()
		}

	default:
		p.Log().Debug("Dropping batch period start msg propagation", "batch_index", msg.BatchIndex)
	}
}

// BatchPeriodAnswerMsg
func (pm *ProtocolManager) batchPeriodAnswerMsgBroadcastLoop() {
	log.Info("Start batchPeriodAnswerMsg broadcast routine")
	for obj := range pm.batchAnswerMsgSub.Chan() {
		if ee, ok := obj.Data.(core.BatchPeriodAnswerEvent); ok {
			log.Debug("Broadcast BatchPeriodAnswerEvent", "sequencer", ee.Msg.Sequencer, "tx_len", len(ee.Msg.Txs))
			pm.BroadcastBatchPeriodAnswerMsg(ee.Msg) // First propagate block to peers
		}
	}
}

func (pm *ProtocolManager) BroadcastBatchPeriodAnswerMsg(msg *types.BatchPeriodAnswerMsg) {
	peers := pm.consensusPeers.PeersWithoutEndMsg(msg.Hash())
	for _, p := range peers {
		p.AsyncSendBatchPeriodAnswerMsg(msg)
	}
	log.Trace("Broadcast batch period answer msg")
}

func (p *peer) AsyncSendBatchPeriodAnswerMsg(msg *types.BatchPeriodAnswerMsg) {
	select {
	case p.queuedBatchAnswerMsg <- msg:
		p.knowBatchPeriodAnswerMsg.Add(msg.Hash())
		for p.knowBatchPeriodAnswerMsg.Cardinality() >= maxKnownEndMsg {
			p.knowBatchPeriodAnswerMsg.Pop()
		}

	default:
		p.Log().Debug("Dropping batch period end msg propagation", "start_index", msg.StartIndex)
	}
}

// RollbackMsg
func (pm *ProtocolManager) rollbackMsgBroadcastLoop() {
	log.Info("Start rollbackMsg broadcast routine")
	// automatically stops if unsubscribe
	for obj := range pm.rollbackMsgSub.Chan() {
		if fe, ok := obj.Data.(core.RollbackStartEvent); ok {
			log.Debug("Got BatchPeriodAnswerEvent, broadcast it")

			pm.BroadcastRollbackMsg(fe.Msg) // First propagate block to peers
		}
	}
}

func (pm *ProtocolManager) BroadcastRollbackMsg(rollbackMsg *types.RollbackMsg) {
	peers := pm.consensusPeers.PeersWithoutRollbackMsg(rollbackMsg.Hash())
	for _, p := range peers {
		p.AsyncSendRollbackMsg(rollbackMsg)
	}
	log.Trace("Broadcast fraud proof reorg msg")
}

func (p *peer) AsyncSendRollbackMsg(rollbackMsg *types.RollbackMsg) {
	select {
	case p.queuedRollback <- rollbackMsg:
		p.knowRollbackMsg.Add(rollbackMsg.Hash())
		for p.knowRollbackMsg.Cardinality() >= maxKnownRollbackMsg {
			p.knowRollbackMsg.Pop()
		}

	default:
		p.Log().Debug("Dropping producers propagation")
	}
}

// ---------------------------- Proposers ----------------------------

// SendBatchPeriodStart sends a batch of transaction receipts, corresponding to the
// ones requested from an already RLP encoded format.
func (p *peer) SendBatchPeriodStart(bps *types.BatchPeriodStartMsg) error {
	p.knowBatchPeriodStartMsg.Add(bps.BatchIndex)
	// Mark all the producers as known, but ensure we don't overflow our limits
	for p.knowBatchPeriodStartMsg.Cardinality() >= maxKnownStartMsg {
		p.knowBatchPeriodStartMsg.Pop()
	}
	return p2p.Send(p.rw, BatchPeriodStartMsg, bps)
}

func (p *peer) SendBatchPeriodAnswer(bpa *types.BatchPeriodAnswerMsg) error {
	p.knowBatchPeriodAnswerMsg.Add(bpa.Hash())
	// Mark all the producers as known, but ensure we don't overflow our limits
	for p.knowBatchPeriodAnswerMsg.Cardinality() >= maxKnownEndMsg {
		p.knowBatchPeriodAnswerMsg.Pop()
	}
	return p2p.Send(p.rw, BatchPeriodAnswerMsg, bpa)
}

func (p *peer) SendRollback(fpr *types.RollbackMsg) error {
	p.knowRollbackMsg.Add(fpr.Hash())
	// Mark all the producers as known, but ensure we don't overflow our limits
	for p.knowRollbackMsg.Cardinality() >= maxKnownRollbackMsg {
		p.knowRollbackMsg.Pop()
	}
	return p2p.Send(p.rw, RollbackMsg, fpr)
}
