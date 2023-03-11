package derive

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mantlenetworkio/mantle/mt-node/eth"
)

func TestPayloadsByNumber(t *testing.T) {
	p := payloadsByNumber{}
	mk := func(i uint64) payloadAndSize {
		return payloadAndSize{
			payload: &eth.ExecutionPayload{
				BlockNumber: eth.Uint64Quantity(i),
			},
		}
	}
	// add payload A, check it was added
	a := mk(123)
	heap.Push(&p, a)
	require.Equal(t, p.Len(), 1)
	require.Equal(t, p[0], a)

	// add payload B, check it was added in top-priority spot
	b := mk(100)
	heap.Push(&p, b)
	require.Equal(t, p.Len(), 2)
	require.Equal(t, p[0], b)

	// add payload C, check it did not get first like B, since block num is higher
	c := mk(150)
	heap.Push(&p, c)
	require.Equal(t, p.Len(), 3)
	require.Equal(t, p[0], b) // still b

	// pop b
	heap.Pop(&p)
	require.Equal(t, p.Len(), 2)
	require.Equal(t, p[0], a)

	// pop a
	heap.Pop(&p)
	require.Equal(t, p.Len(), 1)
	require.Equal(t, p[0], c)

	// pop c
	heap.Pop(&p)
	require.Equal(t, p.Len(), 0)

	// duplicate entry
	heap.Push(&p, b)
	require.Equal(t, p.Len(), 1)
	heap.Push(&p, b)
	require.Equal(t, p.Len(), 2)
	heap.Pop(&p)
	require.Equal(t, p.Len(), 1)
}

func TestPayloadMemSize(t *testing.T) {
	require.Equal(t, payloadMemFixedCost, payloadMemSize(nil), "nil is same fixed cost")
	require.Equal(t, payloadMemFixedCost, payloadMemSize(&eth.ExecutionPayload{}), "empty payload fixed cost")
	require.Equal(t, payloadMemFixedCost+payloadTxMemOverhead, payloadMemSize(&eth.ExecutionPayload{Transactions: []eth.Data{nil}}), "nil tx counts")
	require.Equal(t, payloadMemFixedCost+payloadTxMemOverhead, payloadMemSize(&eth.ExecutionPayload{Transactions: []eth.Data{make([]byte, 0)}}), "empty tx counts")
	require.Equal(t, payloadMemFixedCost+4*payloadTxMemOverhead+42+1337+0+1,
		payloadMemSize(&eth.ExecutionPayload{Transactions: []eth.Data{
			make([]byte, 42),
			make([]byte, 1337),
			make([]byte, 0),
			make([]byte, 1),
		}}), "mixed txs")
}

func TestPayloadsQueue(t *testing.T) {
	pq := PayloadsQueue{
		MaxSize: payloadMemFixedCost * 3,
		SizeFn:  payloadMemSize,
	}
	require.Equal(t, 0, pq.Len())
	require.Equal(t, (*eth.ExecutionPayload)(nil), pq.Peek())
	require.Equal(t, (*eth.ExecutionPayload)(nil), pq.Pop())

	a := &eth.ExecutionPayload{BlockNumber: 3}
	b := &eth.ExecutionPayload{BlockNumber: 4}
	c := &eth.ExecutionPayload{BlockNumber: 5}
	bAlt := &eth.ExecutionPayload{BlockNumber: 4}
	require.NoError(t, pq.Push(b))
	require.Equal(t, pq.Len(), 1)
	require.Equal(t, pq.Peek(), b)

	require.Error(t, pq.Push(nil), "cannot add nil payloads")

	require.NoError(t, pq.Push(c))
	require.Equal(t, pq.Len(), 2)
	require.Equal(t, pq.MemSize(), 2*payloadMemFixedCost)
	require.Equal(t, pq.Peek(), b, "expecting b to still be the lowest number payload")

	require.NoError(t, pq.Push(a))
	require.Equal(t, pq.Len(), 3)
	require.Equal(t, pq.MemSize(), 3*payloadMemFixedCost)
	require.Equal(t, pq.Peek(), a, "expecting a to be new lowest number")

	require.Equal(t, pq.Pop(), a)
	require.Equal(t, pq.Len(), 2, "expecting to pop the lowest")

	require.NoError(t, pq.Push(bAlt))
	require.Equal(t, pq.Len(), 3)
	require.Equal(t, pq.Peek(), b, "expecting b to be lowest, compared to bAlt and c")

	require.Equal(t, pq.Pop(), b)
	require.Equal(t, pq.Len(), 2)
	require.Equal(t, pq.MemSize(), 2*payloadMemFixedCost)

	require.Equal(t, pq.Pop(), bAlt)
	require.Equal(t, pq.Len(), 1)
	require.Equal(t, pq.Peek(), c, "expecting c to only remain")

	d := &eth.ExecutionPayload{BlockNumber: 5, Transactions: []eth.Data{make([]byte, payloadMemFixedCost*3+1)}}
	require.Error(t, pq.Push(d), "cannot add payloads that are too large")

	require.NoError(t, pq.Push(b))
	require.Equal(t, pq.Len(), 2, "expecting b, c")
	require.Equal(t, pq.Peek(), b)
	require.NoError(t, pq.Push(a))
	require.Equal(t, pq.Len(), 3, "expecting a, b, c")
	require.Equal(t, pq.Peek(), a)
	require.NoError(t, pq.Push(bAlt))
	require.Equal(t, pq.Len(), 3, "expecting b, bAlt, c")
	require.NotContainsf(t, pq.pq[:], a, "a should be dropped after 3 items already exist under max size constraint")
}
