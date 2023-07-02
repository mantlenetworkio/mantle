package abnormal

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	maxAbnormalNodeLength = 100
)

func NewNode(pk string, data, sig []byte) *Node {
	return &Node{
		Pubkey:    pk,
		Data:      data,
		Signature: sig,
	}
}

func (n *Node) Equal(node *Node) bool {
	if node == nil {
		return false
	}
	if n.Pubkey == node.Pubkey && bytes.Equal(n.Signature, node.Signature) {
		return true
	}
	return false
}

func NewAbnormal(reason string, nodes []*Node) *Abnormal {
	abnormal := &Abnormal{
		FailReason:   reason,
		Nodes:        make([]*Node, 0, maxAbnormalNodeLength),
		AbnormalLock: sync.RWMutex{},
	}
	abnormal.appendNewNodes(nodes)
	return abnormal
}

func (a *Abnormal) String() string {
	sb := strings.Builder{}
	sb.WriteString("reason:" + a.FailReason + " is_unicast:" + strconv.FormatBool(a.IsUnicast) + "\n")
	sb.WriteString(fmt.Sprintf("nodes:%+v\n", a.Nodes))
	return sb.String()
}

func (a *Abnormal) appendNewNodes(newNodes []*Node) {
	if len(newNodes) > maxAbnormalNodeLength {
		a.Nodes = newNodes[len(newNodes)-maxAbnormalNodeLength:]
	} else if len(newNodes)+len(a.Nodes) > maxAbnormalNodeLength {
		exceedAmount := len(newNodes) + len(a.Nodes) - maxAbnormalNodeLength
		a.Nodes = a.Nodes[exceedAmount:]
		a.Nodes = append(a.Nodes, newNodes...)
	} else {
		a.Nodes = append(a.Nodes, newNodes...)
	}
}

func (a *Abnormal) SetAbnormal(reason string, nodes []*Node, isUnicast bool) {
	a.AbnormalLock.Lock()
	defer a.AbnormalLock.Unlock()
	a.FailReason = reason
	a.IsUnicast = isUnicast
	a.appendNewNodes(nodes)
}
