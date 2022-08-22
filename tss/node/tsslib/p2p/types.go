package p2p

import (
	maddr "github.com/multiformats/go-multiaddr"
	"strings"
)

type AddrList []maddr.Multiaddr

func (al *AddrList) Type() string {
	//TODO implement me
	panic("implement me")
}

type Config struct {
	RendezvousString string
	Port             int
	BootstrapPeers   []string
	ExternalIP       string
}

func (al *AddrList) String() string {
	addresses := make([]string, len(*al))
	for i, addr := range *al {
		addresses[i] = addr.String()
	}
	return strings.Join(addresses, ",")
}

func (al *AddrList) Set(value string) error {
	addr, err := maddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}
