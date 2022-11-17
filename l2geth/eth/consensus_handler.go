package eth

import (
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique"
	"github.com/mantlenetworkio/mantle/l2geth/core/forkid"
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

func (pm *ProtocolManager) consensusHandler(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	p := pm.newPeer(int(eth64), peer, rw)
	select {
	case pm.newPeerCh <- p:
		pm.wg.Add(1)
		defer pm.wg.Done()
		// Ignore maxPeers if this is a trusted peer
		if pm.peersTmp.Len() >= pm.maxPeers && !p.Peer.Info().Network.Trusted {
			return p2p.DiscTooManyPeers
		}
		p.Log().Debug("Ethereum peer connected", "name", p.Name())
		// Execute the Ethereum handshake
		var (
			genesis = pm.blockchain.Genesis()
			head    = pm.blockchain.CurrentHeader()
			hash    = head.Hash()
			number  = head.Number.Uint64()
			td      = pm.blockchain.GetTd(hash, number)
		)
		if err := p.Handshake(pm.networkID, td, hash, genesis.Hash(), forkid.NewID(pm.blockchain), pm.forkFilter); err != nil {
			p.Log().Debug("Ethereum handshake failed", "err", err)
			return err
		}
		if rw, ok := p.rw.(*meteredMsgReadWriter); ok {
			rw.Init(p.version)
		}
		// Register the peer locally
		if err := pm.peersTmp.Register(p); err != nil {
			p.Log().Error("Ethereum peer registration failed", "err", err)
			return err
		}
		defer pm.removePeer(p.id)

		// Handle incoming messages until the connection is torn down
		for {
			if err := pm.handleConsensusMsg(p); err != nil {
				p.Log().Debug("Ethereum message handling failed", "err", err)
				return err
			}
		}
	case <-pm.quitSync:
		return p2p.DiscQuitting
	}
}

func (pm *ProtocolManager) handleConsensusMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on its contents
	switch {
	case msg.Code == ProducersMsg:
		log.Debug(fmt.Sprintf("Get ProducersMsg from %v", p.id))
		// A batch of block bodies arrived to one of our previous requests
		var tmp []byte
		var proUpdate clique.ProducerUpdate
		if err := msg.Decode(&tmp); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		proUpdate.Deserialize(tmp)
		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			eg.SetProducers(proUpdate)
		}

	case msg.Code == GetProducersMsg:
		var proupdate clique.ProducerUpdate
		var getProducers clique.GetProducers
		if err := msg.Decode(&getProducers); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			proupdate = eg.GetProducers(getProducers)
		}

		return p.SendProducers(proupdate)
	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

func (p *peer) AsyncSendProducers(prs *clique.ProducerUpdate) {
	select {
	case p.queuedPrs <- prs:
		log.Info(fmt.Sprintf("in len : %v", len(p.queuedPrs)))
		p.knowPrs.Add(prs.Producers.Index)
		// Mark all the producers as known, but ensure we don't overflow our limits
		for p.knowPrs.Cardinality() >= maxKnownPrs {
			p.knowPrs.Pop()
		}

	default:
		p.Log().Debug("Dropping producers propagation", "block number", prs.Producers.Number)
	}
}

// SendProducers sends a batch of transaction receipts, corresponding to the
// ones requested from an already RLP encoded format.
func (p *peer) SendProducers(proUpdate clique.ProducerUpdate) error {
	p.knowPrs.Add(proUpdate.Producers.Index)
	// Mark all the producers as known, but ensure we don't overflow our limits
	for p.knowPrs.Cardinality() >= maxKnownPrs {
		p.knowPrs.Pop()
	}
	return p2p.Send(p.rw, ProducersMsg, proUpdate.Serialize())
}

func (p *peer) RequestProducers(producers clique.Producers) error {
	return p2p.Send(p.rw, GetProducersMsg, producers)
}
