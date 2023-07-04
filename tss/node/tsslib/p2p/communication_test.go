package p2p

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	maddr "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"

	"github.com/mantlenetworkio/mantle/tss/node/store"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/messages"
)

type CommunicationTestSuite struct {
}

var _ = Suite(&CommunicationTestSuite{})

func (CommunicationTestSuite) TestBasicCommunication(c *C) {
	//new level db storage
	store, err := store.NewStorage("/Users/jayliu/Projects/rde/local/data/test/db")
	c.Assert(err, IsNil)
	comm, err := NewCommunication(nil, 6668, "", false, store)
	c.Assert(err, IsNil)
	c.Assert(comm, NotNil)
	comm.SetSubscribe(messages.TSSKeyGenMsg, "hello", make(chan *Message))
	c.Assert(comm.getSubscriber(messages.TSSKeySignMsg, "hello"), IsNil)
	c.Assert(comm.getSubscriber(messages.TSSKeyGenMsg, "hello"), NotNil)
	comm.CancelSubscribe(messages.TSSKeyGenMsg, "hello")
	comm.CancelSubscribe(messages.TSSKeyGenMsg, "whatever")
	comm.CancelSubscribe(messages.TSSKeySignMsg, "asdsdf")
}

func checkExist(a []maddr.Multiaddr, b string) bool {
	for _, el := range a {
		if el.String() == b {
			return true
		}
	}
	return false
}

func (CommunicationTestSuite) TestEstablishP2pCommunication(c *C) {
	bootstrapPeer := "/ip4/127.0.0.1/tcp/17551/p2p/..."
	bootstrapPrivKey := "..."
	fakeExternalIP := "127.0.0.1"
	fakeExternalMultiAddr := "/ip4/127.0.0.1/tcp/12220"
	privKey, err := crypto.HexToECDSA(bootstrapPrivKey)
	c.Assert(err, IsNil)
	//new level db storage
	store, err := store.NewStorage("/Users/jayliu/Projects/rde/local/data/test/db")
	c.Assert(err, IsNil)
	var bootstrapPeers AddrList
	bootstrapPeers.Set(bootstrapPeer)
	comm, err := NewCommunication(bootstrapPeers, 12220, fakeExternalIP, false, store)
	c.Assert(err, IsNil)
	c.Assert(comm.Start(crypto.FromECDSA(privKey)), IsNil)

	// we add test for external ip advertising
	c.Assert(checkExist(comm.host.Addrs(), fakeExternalMultiAddr), Equals, true)

}

func TestDos(t *testing.T) {
	bootstrapPeer := "/ip4/127.0.0.1/tcp/17551/p2p/..."
	bootstrapPrivKey := "..."
	fakeExternalIP := "127.0.0.1"
	fakeExternalMultiAddr := "/ip4/127.0.0.1/tcp/12220"
	privKey, err := crypto.HexToECDSA(bootstrapPrivKey)
	require.NoError(t, err)
	//new level db storage
	store, err := store.NewStorage("/Users/jayliu/Projects/rde/local/data/test/db")
	require.NoError(t, err)
	var bootstrapPeers AddrList
	bootstrapPeers.Set(bootstrapPeer)
	comm, err := NewCommunication(bootstrapPeers, 12220, fakeExternalIP, false, store)
	require.NoError(t, err)
	err = comm.Start(crypto.FromECDSA(privKey))
	require.NoError(t, err)

	// we add test for external ip advertising
	require.EqualValues(t, checkExist(comm.host.Addrs(), fakeExternalMultiAddr), true)

}
