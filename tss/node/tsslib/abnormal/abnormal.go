package abnormal

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func NewNode(pk string, data, sig []byte) Node {
	return Node{
		Pubkey:    pk,
		Data:      data,
		Signature: sig,
	}
}

func (n *Node) Equal(node Node) bool {
	if n.Pubkey == node.Pubkey && bytes.Equal(n.Signature, node.Signature) {
		return true
	}
	return false
}

func NewAbnormal(reason string, nodes []Node) Abnormal {
	return Abnormal{
		FailReason:   reason,
		Nodes:        nodes,
		AbnormalLock: &sync.RWMutex{},
	}
}

func (a Abnormal) String() string {
	sb := strings.Builder{}
	sb.WriteString("reason:" + a.FailReason + " is_unicast:" + strconv.FormatBool(a.IsUnicast) + "\n")
	sb.WriteString(fmt.Sprintf("nodes:%+v\n", a.Nodes))
	return sb.String()
}

func (a *Abnormal) SetAbnormal(reason string, nodes []Node, isUnicast bool) {
	a.AbnormalLock.Lock()
	defer a.AbnormalLock.Unlock()
	a.FailReason = reason
	a.IsUnicast = isUnicast
	a.Nodes = append(a.Nodes, nodes...)
}
