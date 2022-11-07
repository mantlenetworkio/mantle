package eth

import (
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique"
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
		// A batch of block bodies arrived to one of our previous requests
		var producers clique.Producers
		if err := msg.Decode(&producers); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}

		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			eg.SetProducers(producers)
		}

	case msg.Code == GetProducersMsg:
		var producers clique.Producers
		var getProducers clique.GetProducers
		if err := msg.Decode(&getProducers); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			producers = eg.GetProducers(getProducers)
		}

		return p.SendProducers(producers)
	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

func (p *peer) AsyncSendProducers(prs *clique.Producers) {
	// todo add producers to peer: check? index
	select {
	case p.queuedPrs <- prs:
		p.knowPrs.Add(prs.Index)
		// Mark all the producers as known, but ensure we don't overflow our limits
		for p.knowPrs.Cardinality() >= maxKnownPrs {
			p.knowPrs.Pop()
		}

	default:
		p.Log().Debug("Dropping producers propagation", "block number", prs.Number)
	}
}

// SendProducers sends a batch of transaction receipts, corresponding to the
// ones requested from an already RLP encoded format.
func (p *peer) SendProducers(producers clique.Producers) error {
	// todo send producers with signature
	return p2p.Send(p.rw, ProducersMsg, producers)
}

func (p *peer) RequestProducers(producers clique.Producers) error {
	return p2p.Send(p.rw, GetProducersMsg, producers)
}
