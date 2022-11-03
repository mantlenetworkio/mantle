package eth

import (
	"fmt"

	"github.com/mantlenetworkio/mantle/l2geth/consensus/clique"
	"github.com/mantlenetworkio/mantle/l2geth/log"
	"github.com/mantlenetworkio/mantle/l2geth/p2p"
	"github.com/mantlenetworkio/mantle/l2geth/p2p/enode"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
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
		var request clique.Producers
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		var producers []rlp.RawValue
		if encoded, err := rlp.EncodeToBytes(request); err != nil {
			log.Error("Failed to encode receipt", "err", err)
		} else {
			producers = append(producers, encoded)
		}

		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			eg.SetProducers(request)
		}

	case msg.Code == GetProducersMsg:
		// Decode the retrieval message
		msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
		if _, err := msgStream.List(); err != nil {
			return err
		}
		// Gather state data until the fetch or network limits is reached
		var producers []rlp.RawValue
		var results clique.Producers

		var getProducers clique.GetProducers

		if eg, ok := pm.blockchain.Engine().(*clique.Clique); ok {
			results = eg.GetProducers(getProducers)
		}

		if encoded, err := rlp.EncodeToBytes(results); err != nil {
			log.Error("Failed to encode receipt", "err", err)
		} else {
			producers = append(producers, encoded)
		}
		return p.SendProducerRLP(producers)
	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

// SendProducerRLP sends a batch of transaction receipts, corresponding to the
// ones requested from an already RLP encoded format.
func (p *peer) SendProducerRLP(producers []rlp.RawValue) error {
	return p2p.Send(p.rw, ProducersMsg, producers)
}

func (p *peer) RequestProducerRLP(producers []rlp.RawValue) error {
	return p2p.Send(p.rw, GetProducersMsg, producers)
}
