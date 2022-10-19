package clique

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	crand "crypto/rand"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"testing/quick"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/stretchr/testify/assert"
)

func TestSequencerSetBasic(t *testing.T) {
	// empty or nil seqidator lists are allowed,
	// but attempting to IncrementProducerPriority on them will panic.
	sset := NewSequencerSet([]*Sequencer{})
	assert.Panics(t, func() { sset.IncrementProducerPriority(1) })

	sset = NewSequencerSet(nil)
	assert.Panics(t, func() { sset.IncrementProducerPriority(1) })

	assert.EqualValues(t, sset, sset.Copy())
	var addr common.Address
	copy(addr[:], "some seq")
	assert.False(t, sset.HasAddress(addr))
	idx, seq := sset.GetByAddress(addr)
	assert.EqualValues(t, -1, idx)
	assert.Nil(t, seq)
	addr, seq = sset.GetByIndex(-100)
	assert.EqualValues(t, [20]byte{}, addr)
	assert.Nil(t, seq)
	addr, seq = sset.GetByIndex(0)
	assert.EqualValues(t, [20]byte{}, addr)
	assert.Nil(t, seq)
	addr, seq = sset.GetByIndex(100)
	assert.EqualValues(t, [20]byte{}, addr)
	assert.Nil(t, seq)
	assert.Zero(t, sset.Size())
	assert.Equal(t, int64(0), sset.TotalPower())
	assert.Nil(t, sset.GetProducer())
	// assert.Equal(t, []byte{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4,
	//	0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95,
	//	0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}, sset.Hash())
	// add
	seq = randSequencer(sset.TotalPower())

	assert.NoError(t, sset.UpdateWithChangeSet([]*Sequencer{seq}))

	assert.True(t, sset.HasAddress(seq.Address))
	idx, _ = sset.GetByAddress(seq.Address)
	assert.EqualValues(t, 0, idx)
	addr, _ = sset.GetByIndex(0)
	assert.Equal(t, seq.Address, addr)
	assert.Equal(t, 1, sset.Size())
	assert.Equal(t, seq.Power, sset.TotalPower())
	assert.NotNil(t, sset.Hash())
	assert.NotPanics(t, func() { sset.IncrementProducerPriority(1) })
	assert.Equal(t, seq.Address, sset.GetProducer().Address)

	// update
	seq = randSequencer(sset.TotalPower())
	assert.NoError(t, sset.UpdateWithChangeSet([]*Sequencer{seq}))
	_, seq = sset.GetByAddress(seq.Address)
	seq.Power += 100
	proposerPriority := seq.ProducerPriority

	seq.ProducerPriority = 0
	assert.NoError(t, sset.UpdateWithChangeSet([]*Sequencer{seq}))
	_, seq = sset.GetByAddress(seq.Address)
	assert.Equal(t, proposerPriority, seq.ProducerPriority)
}

func TestSequencerSetSequencerBasic(t *testing.T) {
	seq := RandSequencer(false, 1)
	badVal := &Sequencer{}

	testCases := []struct {
		seqs SequencerSet
		err  bool
		msg  string
	}{
		{
			seqs: SequencerSet{},
			err:  true,
			msg:  "sequencer set is nil or empty",
		},
		{
			seqs: SequencerSet{
				Sequencers: []*Sequencer{},
			},
			err: true,
			msg: "sequencer set is nil or empty",
		},
		{
			seqs: SequencerSet{
				Sequencers: []*Sequencer{seq},
			},
			err: true,
			msg: "producer failed validate basic, error: nil sequencer",
		},
		{
			seqs: SequencerSet{
				Sequencers: []*Sequencer{badVal},
			},
			err: true,
			msg: "invalid sequencer #0: sequencer does not have a public key",
		},
		{
			seqs: SequencerSet{
				Sequencers: []*Sequencer{seq},
				Producer:   seq,
			},
			err: false,
			msg: "",
		},
	}

	for _, tc := range testCases {
		err := tc.seqs.ValidateBasic()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestCopy(t *testing.T) {
	sset := randSequencerSet(10)
	ssetHash := sset.Hash()
	if len(ssetHash) == 0 {
		t.Fatalf("SequencerSet had unexpected zero hash")
	}

	ssetCopy := sset.Copy()
	ssetCopyHash := ssetCopy.Hash()

	if !bytes.Equal(ssetHash.Bytes(), ssetCopyHash.Bytes()) {
		t.Fatalf("SequencerSet copy had wrong hash. Orig: %X, Copy: %X", ssetHash, ssetCopyHash)
	}
}

// Test that IncrementProducerPriority requires positive times.
func TestIncrementProducerPriorityPositiveTimes(t *testing.T) {
	sset := NewSequencerSet([]*Sequencer{
		newSequencer(byteToCommonAddr([]byte("foo")), randPubKey(), 1000),
		newSequencer(byteToCommonAddr([]byte("bar")), randPubKey(), 300),
		newSequencer(byteToCommonAddr([]byte("baz")), randPubKey(), 330),
	})

	assert.Panics(t, func() { sset.IncrementProducerPriority(-1) })
	assert.Panics(t, func() { sset.IncrementProducerPriority(0) })
	sset.IncrementProducerPriority(1)
}

func BenchmarkSequencerSetCopy(b *testing.B) {
	b.StopTimer()
	sset := NewSequencerSet([]*Sequencer{})
	for i := 0; i < 1000; i++ {
		privKey, _ := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
		pubKey := privKey.PublicKey
		seq := NewSequencer(randAddress(), pubKey, 10)
		err := sset.UpdateWithChangeSet([]*Sequencer{seq})
		if err != nil {
			panic("Failed to add sequencer")
		}
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sset.Copy()
	}
}

//-------------------------------------------------------------------

func TestProducerSelection1(t *testing.T) {
	sset := NewSequencerSet([]*Sequencer{
		newSequencer(byteToCommonAddr([]byte("foo")), randPubKey(), 1000),
		newSequencer(byteToCommonAddr([]byte("bar")), randPubKey(), 300),
		newSequencer(byteToCommonAddr([]byte("baz")), randPubKey(), 330),
	})
	var producers []string
	for i := 0; i < 5; i++ {
		seq := sset.GetProducer()
		producers = append(producers, seq.Address.String())
		sset.IncrementProducerPriority(1)
	}
	expected := byteToCommonAddr([]byte("foo")).String() + byteToCommonAddr([]byte("baz")).String() + byteToCommonAddr([]byte("foo")).String() + byteToCommonAddr([]byte("bar")).String() + byteToCommonAddr([]byte("foo")).String()
	if expected != strings.Join(producers, "") {
		t.Errorf("expected sequence of proposers was\n%v\nbut got \n%v", expected, strings.Join(producers, ""))
	}
}

func TestProducerSelection2(t *testing.T) {
	addr0 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	addr1 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	addr2 := []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	// when all voting power is same, we go in order of addresses
	seq0, seq1, seq2 := newSequencer(byteToCommonAddr(addr0), randPubKey(), 100), newSequencer(byteToCommonAddr(addr1), randPubKey(), 100), newSequencer(byteToCommonAddr(addr2), randPubKey(), 100)
	seqList := []*Sequencer{seq0, seq1, seq2}
	seqs := NewSequencerSet(seqList)
	for i := 0; i < len(seqList)*5; i++ {
		ii := (i) % len(seqList)
		prop := seqs.GetProducer()
		if !bytes.Equal(prop.Address.Bytes(), seqList[ii].Address.Bytes()) {
			t.Fatalf("(%d): Expected %X. Got %X", i, seqList[ii].Address, prop.Address)
		}
		seqs.IncrementProducerPriority(1)
	}

	// One sequencer has more than the others, but not enough to produce twice in a row
	*seq2 = *newSequencer(byteToCommonAddr(addr2), randPubKey(), 400)
	seqs = NewSequencerSet(seqList)
	// seqs.IncrementProducerPriority(1)
	prop := seqs.GetProducer()
	if !bytes.Equal(prop.Address.Bytes(), addr2) {
		t.Fatalf("Expected address with highest voting power to be first proposer. Got %X", prop.Address)
	}
	seqs.IncrementProducerPriority(1)
	prop = seqs.GetProducer()
	if !bytes.Equal(prop.Address.Bytes(), addr0) {
		t.Fatalf("Expected smallest address to be seqidator. Got %X", prop.Address)
	}

	// One sequencer has more than the others, and enough to be produce twice in a row
	*seq2 = *newSequencer(byteToCommonAddr(addr2), randPubKey(), 401)
	seqs = NewSequencerSet(seqList)
	prop = seqs.GetProducer()
	if !bytes.Equal(prop.Address.Bytes(), addr2) {
		t.Fatalf("Expected address with highest voting power to be first proposer. Got %X", prop.Address)
	}
	seqs.IncrementProducerPriority(1)
	prop = seqs.GetProducer()
	if !bytes.Equal(prop.Address.Bytes(), addr2) {
		t.Fatalf("Expected address with highest voting power to be second proposer. Got %X", prop.Address)
	}
	seqs.IncrementProducerPriority(1)
	prop = seqs.GetProducer()
	if !bytes.Equal(prop.Address.Bytes(), addr0) {
		t.Fatalf("Expected smallest address to be seqidator. Got %X", prop.Address)
	}

	// each sequencer should be the producer a proportional number of times
	seq0, seq1, seq2 = newSequencer(byteToCommonAddr(addr0), randPubKey(), 4), newSequencer(byteToCommonAddr(addr1), randPubKey(), 5), newSequencer(byteToCommonAddr(addr2), randPubKey(), 3)
	seqList = []*Sequencer{seq0, seq1, seq2}
	propCount := make([]int, 3)
	seqs = NewSequencerSet(seqList)
	N := 1
	for i := 0; i < 120*N; i++ {
		prop := seqs.GetProducer()
		ii := prop.Address[19]
		propCount[ii]++
		seqs.IncrementProducerPriority(1)
	}

	if propCount[0] != 40*N {
		t.Fatalf(
			"Expected prop count for seqidator with 4/12 of voting power to be %d/%d. Got %d/%d",
			40*N,
			120*N,
			propCount[0],
			120*N,
		)
	}
	if propCount[1] != 50*N {
		t.Fatalf(
			"Expected prop count for seqidator with 5/12 of voting power to be %d/%d. Got %d/%d",
			50*N,
			120*N,
			propCount[1],
			120*N,
		)
	}
	if propCount[2] != 30*N {
		t.Fatalf(
			"Expected prop count for seqidator with 3/12 of voting power to be %d/%d. Got %d/%d",
			30*N,
			120*N,
			propCount[2],
			120*N,
		)
	}
}

func TestProducerSelection3(t *testing.T) {
	sset := NewSequencerSet([]*Sequencer{
		newSequencer(byteToCommonAddr([]byte("asequencer_address12")), randPubKey(), 1),
		newSequencer(byteToCommonAddr([]byte("bsequencer_address12")), randPubKey(), 1),
		newSequencer(byteToCommonAddr([]byte("csequencer_address12")), randPubKey(), 1),
		newSequencer(byteToCommonAddr([]byte("dsequencer_address12")), randPubKey(), 1),
	})

	producerOrder := make([]*Sequencer, 4)
	for i := 0; i < 4; i++ {
		// need to give all sequencer to have keys
		privKey, _ := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
		pk := privKey.PublicKey
		sset.Sequencers[i].PubKey = pk
		producerOrder[i] = sset.GetProducer()
		sset.IncrementProducerPriority(1)
	}

	// i for the loop
	// j for the times
	// we should go in order for ever, despite some IncrementProducerPriority with times > 1
	var (
		i int
		j int32
	)
	for ; i < 10000; i++ {
		got := sset.GetProducer().Address
		expected := producerOrder[j%4].Address
		if !bytes.Equal(got.Bytes(), expected.Bytes()) {
			t.Fatalf(fmt.Sprintf("sset.Producer (%X) does not match expected producer (%X) for (%d, %d)", got, expected, i, j))
		}

		// serialize, deserialize, check producer
		// b := sset.toBytes()
		// sset = sset.fromBytes(b)

		computed := sset.GetProducer() // findGetProducer()
		if i != 0 {
			if !bytes.Equal(got.Bytes(), computed.Address.Bytes()) {
				t.Fatalf(
					fmt.Sprintf(
						"sset.Producer (%X) does not match computed proposer (%X) for (%d, %d)",
						got,
						computed.Address,
						i,
						j,
					),
				)
			}
		}

		// times is usually 1
		times := int32(1)
		mod := (rand.Int() % 5) + 1
		if rand.Int()%mod > 0 {
			// sometimes its up to 5
			times = (rand.Int31() % 4) + 1
		}
		sset.IncrementProducerPriority(times)

		j += times
	}
}

func byteToCommonAddr(bytes []byte) common.Address {
	return common.BytesToAddress(bytes)
}

func newSequencer(address common.Address, pubKey ecdsa.PublicKey, power int64) *Sequencer {
	return &Sequencer{Address: address, PubKey: pubKey, Power: power}
}

func randPubKey() ecdsa.PublicKey {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), crand.Reader)
	return privKey.PublicKey
}

func randAddress() common.Address {
	seed := make([]byte, 20)
	rand.Read(seed)
	var addr common.Address
	copy(addr[:], seed)
	return addr
}

func randSequencer(totalVotingPower int64) *Sequencer {
	// this modulo limits the ProducerPriority/Power to stay in the
	// bounds of MaxTotalPower minus the already existing voting power:
	seq := NewSequencer(randAddress(), randPubKey(), int64(rand.Uint64()%uint64(MaxTotalPower-totalVotingPower)))
	seq.ProducerPriority = rand.Int63() % (MaxTotalPower - totalVotingPower)
	return seq
}

func randSequencerSet(numSequencers int) *SequencerSet {
	sequencers := make([]*Sequencer, numSequencers)
	totalVotingPower := int64(0)
	for i := 0; i < numSequencers; i++ {
		sequencers[i] = randSequencer(totalVotingPower)
		totalVotingPower += sequencers[i].Power
	}
	return NewSequencerSet(sequencers)
}

//func (seqs *SequencerSet) toBytes() []byte {
//
//}
//
//func (seqs *SequencerSet) fromBytes(b []byte) *SequencerSet {
//
//}

//-------------------------------------------------------------------

func TestSequencerSetTotalVotingPowerPanicsOnOverflow(t *testing.T) {
	// NewSequencerSet calls IncrementProducerPriority which calls TotalPower()
	// which should panic on overflows:
	shouldPanic := func() {
		NewSequencerSet([]*Sequencer{
			{Address: byteToCommonAddr([]byte("a")), Power: math.MaxInt64, ProducerPriority: 0},
			{Address: byteToCommonAddr([]byte("b")), Power: math.MaxInt64, ProducerPriority: 0},
			{Address: byteToCommonAddr([]byte("c")), Power: math.MaxInt64, ProducerPriority: 0},
		})
	}

	assert.Panics(t, shouldPanic)
}

func TestAvgProducerPriority(t *testing.T) {
	// Create Sequencer set without calling IncrementProducerPriority:
	tcs := []struct {
		vs   SequencerSet
		want int64
	}{
		0: {SequencerSet{Sequencers: []*Sequencer{{ProducerPriority: 0}, {ProducerPriority: 0}, {ProducerPriority: 0}}}, 0},
		1: {
			SequencerSet{
				Sequencers: []*Sequencer{{ProducerPriority: math.MaxInt64}, {ProducerPriority: 0}, {ProducerPriority: 0}},
			}, math.MaxInt64 / 3,
		},
		2: {
			SequencerSet{
				Sequencers: []*Sequencer{{ProducerPriority: math.MaxInt64}, {ProducerPriority: 0}},
			}, math.MaxInt64 / 2,
		},
		3: {
			SequencerSet{
				Sequencers: []*Sequencer{{ProducerPriority: math.MaxInt64}, {ProducerPriority: math.MaxInt64}},
			}, math.MaxInt64,
		},
		4: {
			SequencerSet{
				Sequencers: []*Sequencer{{ProducerPriority: math.MinInt64}, {ProducerPriority: math.MinInt64}},
			}, math.MinInt64,
		},
	}
	for i, tc := range tcs {
		got := tc.vs.computeAvgProducerPriority()
		assert.Equal(t, tc.want, got, "test case: %v", i)
	}
}

func TestAveragingInIncrementProducerPriority(t *testing.T) {
	// Test that the averaging works as expected inside of IncrementProducerPriority.
	// Each seqidator comes with zero voting power which simplifies reasoning about
	// the expected ProducerPriority.
	tcs := []struct {
		vs    SequencerSet
		times int32
		avg   int64
	}{
		0: {
			SequencerSet{
				Sequencers: []*Sequencer{
					{Address: byteToCommonAddr([]byte("a")), ProducerPriority: 1},
					{Address: byteToCommonAddr([]byte("b")), ProducerPriority: 2},
					{Address: byteToCommonAddr([]byte("c")), ProducerPriority: 3},
				},
			},
			1, 2,
		},
		1: {
			SequencerSet{
				Sequencers: []*Sequencer{
					{Address: byteToCommonAddr([]byte("a")), ProducerPriority: 10},
					{Address: byteToCommonAddr([]byte("b")), ProducerPriority: -10},
					{Address: byteToCommonAddr([]byte("c")), ProducerPriority: 1},
				},
			},
			// this should average twice but the average should be 0 after the first iteration
			// (voting power is 0 -> no changes)
			11,
			0, // 1 / 3
		},
		2: {
			SequencerSet{
				Sequencers: []*Sequencer{
					{Address: byteToCommonAddr([]byte("a")), ProducerPriority: 100},
					{Address: byteToCommonAddr([]byte("b")), ProducerPriority: -10},
					{Address: byteToCommonAddr([]byte("c")), ProducerPriority: 1},
				},
			},
			1, 91 / 3,
		},
	}
	for i, tc := range tcs {
		// work on copy to have the old ProducerPriorities:
		newVset := tc.vs.CopyIncrementProducerPriority(tc.times)
		for _, seq := range tc.vs.Sequencers {
			_, updatedVal := newVset.GetByAddress(seq.Address)
			assert.Equal(t, updatedVal.ProducerPriority, seq.ProducerPriority-tc.avg, "test case: %v", i)
		}
	}
}

func TestAveragingInIncrementProducerPriorityWithVotingPower(t *testing.T) {
	// Other than TestAveragingInIncrementProducerPriority this is a more complete test showing
	// how each ProducerPriority changes in relation to the seqidator's voting power respectively.
	// average is zero in each round:
	vp0 := int64(10)
	vp1 := int64(1)
	vp2 := int64(1)
	total := vp0 + vp1 + vp2
	avg := (vp0 + vp1 + vp2 - total) / 3
	seqs := SequencerSet{Sequencers: []*Sequencer{
		{Address: byteToCommonAddr([]byte{0}), ProducerPriority: 0, Power: vp0},
		{Address: byteToCommonAddr([]byte{1}), ProducerPriority: 0, Power: vp1},
		{Address: byteToCommonAddr([]byte{2}), ProducerPriority: 0, Power: vp2},
	}}
	tcs := []struct {
		seqs                  *SequencerSet
		wantProducerPrioritys []int64
		times                 int32
		wantProducer          *Sequencer
	}{
		0: {
			seqs.Copy(),
			[]int64{
				// Acumm+Power-Avg:
				0 + vp0 - total - avg, // mostest will be subtracted by total voting power (12)
				0 + vp1,
				0 + vp2,
			},
			1,
			seqs.Sequencers[0],
		},
		1: {
			seqs.Copy(),
			[]int64{
				(0 + vp0 - total) + vp0 - total - avg, // this will be mostest on 2nd iter, too
				(0 + vp1) + vp1,
				(0 + vp2) + vp2,
			},
			2,
			seqs.Sequencers[0],
		}, // increment twice -> expect average to be subtracted twice
		2: {
			seqs.Copy(),
			[]int64{
				0 + 3*(vp0-total) - avg, // still mostest
				0 + 3*vp1,
				0 + 3*vp2,
			},
			3,
			seqs.Sequencers[0],
		},
		3: {
			seqs.Copy(),
			[]int64{
				0 + 4*(vp0-total), // still mostest
				0 + 4*vp1,
				0 + 4*vp2,
			},
			4,
			seqs.Sequencers[0],
		},
		4: {
			seqs.Copy(),
			[]int64{
				0 + 4*(vp0-total) + vp0, // 4 iters was mostest
				0 + 5*vp1 - total,       // now this seq is mostest for the 1st time (hence -12==totalPower)
				0 + 5*vp2,
			},
			5,
			seqs.Sequencers[1],
		},
		5: {
			seqs.Copy(),
			[]int64{
				0 + 6*vp0 - 5*total, // mostest again
				0 + 6*vp1 - total,   // mostest once up to here
				0 + 6*vp2,
			},
			6,
			seqs.Sequencers[0],
		},
		6: {
			seqs.Copy(),
			[]int64{
				0 + 7*vp0 - 6*total, // in 7 iters this seq is mostest 6 times
				0 + 7*vp1 - total,   // in 7 iters this seq is mostest 1 time
				0 + 7*vp2,
			},
			7,
			seqs.Sequencers[0],
		},
		7: {
			seqs.Copy(),
			[]int64{
				0 + 8*vp0 - 7*total, // mostest again
				0 + 8*vp1 - total,
				0 + 8*vp2,
			},
			8,
			seqs.Sequencers[0],
		},
		8: {
			seqs.Copy(),
			[]int64{
				0 + 9*vp0 - 7*total,
				0 + 9*vp1 - total,
				0 + 9*vp2 - total,
			}, // mostest
			9,
			seqs.Sequencers[2],
		},
		9: {
			seqs.Copy(),
			[]int64{
				0 + 10*vp0 - 8*total, // after 10 iters this is mostest again
				0 + 10*vp1 - total,   // after 6 iters this seq is "mostest" once and not in between
				0 + 10*vp2 - total,
			}, // in between 10 iters this seq is "mostest" once
			10,
			seqs.Sequencers[0],
		},
		10: {
			seqs.Copy(),
			[]int64{
				0 + 11*vp0 - 9*total,
				0 + 11*vp1 - total, // after 6 iters this seq is "mostest" once and not in between
				0 + 11*vp2 - total,
			}, // after 10 iters this seq is "mostest" once
			11,
			seqs.Sequencers[0],
		},
	}
	for i, tc := range tcs {
		tc.seqs.IncrementProducerPriority(tc.times)

		assert.Equal(t, tc.wantProducer.Address, tc.seqs.GetProducer().Address,
			"test case: %v",
			i)

		for seqIdx, seq := range tc.seqs.Sequencers {
			assert.Equal(t,
				tc.wantProducerPrioritys[seqIdx],
				seq.ProducerPriority,
				"test case: %v, seqidator: %v",
				i,
				seqIdx)
		}
	}
}

func TestSafeAdd(t *testing.T) {
	f := func(a, b int64) bool {
		c, overflow := safeAdd(a, b)
		return overflow || (!overflow && c == a+b)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSafeAddClip(t *testing.T) {
	assert.EqualValues(t, math.MaxInt64, safeAddClip(math.MaxInt64, 10))
	assert.EqualValues(t, math.MaxInt64, safeAddClip(math.MaxInt64, math.MaxInt64))
	assert.EqualValues(t, math.MinInt64, safeAddClip(math.MinInt64, -10))
}

func TestSafeSubClip(t *testing.T) {
	assert.EqualValues(t, math.MinInt64, safeSubClip(math.MinInt64, 10))
	assert.EqualValues(t, 0, safeSubClip(math.MinInt64, math.MinInt64))
	assert.EqualValues(t, math.MinInt64, safeSubClip(math.MinInt64, math.MaxInt64))
	assert.EqualValues(t, math.MaxInt64, safeSubClip(math.MaxInt64, -10))
}

//-------------------------------------------------------------------

func TestEmptySet(t *testing.T) {
	var seqList []*Sequencer
	seqSet := NewSequencerSet(seqList)
	assert.Panics(t, func() { seqSet.IncrementProducerPriority(1) })
	assert.Panics(t, func() { seqSet.RescalePriorities(100) })
	assert.Panics(t, func() { seqSet.shiftByAvgProducerPriority() })
	assert.Panics(t, func() { assert.Zero(t, computeMaxMinPriorityDiff(seqSet)) })
	seqSet.GetProducer()

	// Add to empty set
	v1 := newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 100)
	v2 := newSequencer(byteToCommonAddr([]byte("v2")), randPubKey(), 100)
	seqList = []*Sequencer{v1, v2}
	assert.NoError(t, seqSet.UpdateWithChangeSet(seqList))
	verifySequencerSet(t, seqSet)

	// Delete all seqidators from set
	v1 = newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 0)
	v2 = newSequencer(byteToCommonAddr([]byte("v2")), randPubKey(), 0)
	delList := []*Sequencer{v1, v2}
	assert.Error(t, seqSet.UpdateWithChangeSet(delList))

	// Attempt delete from empty set
	assert.Error(t, seqSet.UpdateWithChangeSet(delList))
}

func TestUpdatesForNewSequencerSet(t *testing.T) {
	v1 := newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 100)
	v2 := newSequencer(byteToCommonAddr([]byte("v2")), randPubKey(), 100)
	seqList := []*Sequencer{v1, v2}
	seqSet := NewSequencerSet(seqList)
	verifySequencerSet(t, seqSet)

	// Verify duplicates are caught in NewSequencerSet() and it panics
	v111 := newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 100)
	v112 := newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 123)
	v113 := newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 234)
	seqList = []*Sequencer{v111, v112, v113}
	assert.Panics(t, func() { NewSequencerSet(seqList) })

	// Verify set including seqidator with voting power 0 cannot be created
	v1 = newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 0)
	v2 = newSequencer(byteToCommonAddr([]byte("v2")), randPubKey(), 22)
	v3 := newSequencer(byteToCommonAddr([]byte("v3")), randPubKey(), 33)
	seqList = []*Sequencer{v1, v2, v3}
	assert.Panics(t, func() { NewSequencerSet(seqList) })

	// Verify set including seqidator with negative voting power cannot be created
	v1 = newSequencer(byteToCommonAddr([]byte("v1")), randPubKey(), 10)
	v2 = newSequencer(byteToCommonAddr([]byte("v2")), randPubKey(), -20)
	v3 = newSequencer(byteToCommonAddr([]byte("v3")), randPubKey(), 30)
	seqList = []*Sequencer{v1, v2, v3}
	assert.Panics(t, func() { NewSequencerSet(seqList) })
}

type testSeq struct {
	name  string
	power int64
}

func permutation(seqList []testSeq) []testSeq {
	if len(seqList) == 0 {
		return nil
	}
	permList := make([]testSeq, len(seqList))
	perm := rand.Perm(len(seqList))
	for i, v := range perm {
		permList[v] = seqList[i]
	}
	return permList
}

func createNewSequencerList(testSeqList []testSeq) []*Sequencer {
	seqList := make([]*Sequencer, 0, len(testSeqList))
	for _, seq := range testSeqList {
		seqList = append(seqList, newSequencer(byteToCommonAddr([]byte(seq.name)), randPubKey(), seq.power))
	}
	return seqList
}

func createNewSequencerSet(testSeqList []testSeq) *SequencerSet {
	return NewSequencerSet(createNewSequencerList(testSeqList))
}

func seqSetTotalProducerPriority(seqSet *SequencerSet) int64 {
	sum := int64(0)
	for _, seq := range seqSet.Sequencers {
		// mind overflow
		sum = safeAddClip(sum, seq.ProducerPriority)
	}
	return sum
}

func verifySequencerSet(t *testing.T, seqSet *SequencerSet) {
	// verify that the capacity and length of seqidators is the same
	assert.Equal(t, len(seqSet.Sequencers), cap(seqSet.Sequencers))

	// verify that the set's total voting power has been updated
	tvp := seqSet.totalPower
	seqSet.updateTotalPower()
	expectedTvp := seqSet.TotalPower()
	assert.Equal(t, expectedTvp, tvp,
		"expected TVP %d. Got %d, seqSet=%s", expectedTvp, tvp, seqSet)

	// verify that seqidator priorities are centered
	seqsCount := int64(len(seqSet.Sequencers))
	tpp := seqSetTotalProducerPriority(seqSet)
	assert.True(t, tpp < seqsCount && tpp > -seqsCount,
		"expected total priority in (-%d, %d). Got %d", seqsCount, seqsCount, tpp)

	// verify that priorities are scaled
	dist := computeMaxMinPriorityDiff(seqSet)
	assert.True(t, dist <= PriorityWindowSizeFactor*tvp,
		"expected priority distance < %d. Got %d", PriorityWindowSizeFactor*tvp, dist)
}

func toTestSeqList(seqList []*Sequencer) []testSeq {
	testList := make([]testSeq, len(seqList))
	for i, seq := range seqList {
		testList[i].name = seq.Address.String()
		testList[i].power = seq.Power
	}
	return testList
}

func renameTestSeqList(otestseq []testSeq) []testSeq {
	ntestseq := make([]testSeq, len(otestseq))
	for i, ts := range otestseq {
		ntestseq[i].name = byteToCommonAddr([]byte(ts.name)).String()
		ntestseq[i].power = ts.power
	}
	return ntestseq
}

func testSeqSet(nSeqs int, power int64) []testSeq {
	seqs := make([]testSeq, nSeqs)
	for i := 0; i < nSeqs; i++ {
		seqs[i] = testSeq{fmt.Sprintf("v%d", i+1), power}
	}
	return seqs
}

type seqSetErrTestCase struct {
	startSeqs  []testSeq
	updateSeqs []testSeq
}

func executeSeqSetErrTestCase(t *testing.T, idx int, tt seqSetErrTestCase) {
	// create a new set and apply updates, keeping copies for the checks
	seqSet := createNewSequencerSet(tt.startSeqs)
	seqSetCopy := seqSet.Copy()
	seqList := createNewSequencerList(tt.updateSeqs)
	seqListCopy := sequencerListCopy(seqList)
	err := seqSet.UpdateWithChangeSet(seqList)

	// for errors check the seqidator set has not been changed
	assert.Error(t, err, "test %d", idx)
	assert.Equal(t, seqSet, seqSetCopy, "test %v", idx)

	// check the parameter list has not changed
	assert.Equal(t, seqList, seqListCopy, "test %v", idx)
}

func TestSeqSetUpdatesDuplicateEntries(t *testing.T) {
	testCases := []seqSetErrTestCase{
		// Duplicate entries in changes
		{ // first entry is duplicated change
			testSeqSet(2, 10),
			[]testSeq{{"v1", 11}, {"v1", 22}},
		},
		{ // second entry is duplicated change
			testSeqSet(2, 10),
			[]testSeq{{"v2", 11}, {"v2", 22}},
		},
		{ // change duplicates are separated by a seqid change
			testSeqSet(2, 10),
			[]testSeq{{"v1", 11}, {"v2", 22}, {"v1", 12}},
		},
		{ // change duplicates are separated by a seqid change
			testSeqSet(3, 10),
			[]testSeq{{"v1", 11}, {"v3", 22}, {"v1", 12}},
		},

		// Duplicate entries in remove
		{ // first entry is duplicated remove
			testSeqSet(2, 10),
			[]testSeq{{"v1", 0}, {"v1", 0}},
		},
		{ // second entry is duplicated remove
			testSeqSet(2, 10),
			[]testSeq{{"v2", 0}, {"v2", 0}},
		},
		{ // remove duplicates are separated by a seqid remove
			testSeqSet(2, 10),
			[]testSeq{{"v1", 0}, {"v2", 0}, {"v1", 0}},
		},
		{ // remove duplicates are separated by a seqid remove
			testSeqSet(3, 10),
			[]testSeq{{"v1", 0}, {"v3", 0}, {"v1", 0}},
		},

		{ // remove and update same seq
			testSeqSet(2, 10),
			[]testSeq{{"v1", 0}, {"v2", 20}, {"v1", 30}},
		},
		{ // duplicate entries in removes + changes
			testSeqSet(2, 10),
			[]testSeq{{"v1", 0}, {"v2", 20}, {"v2", 30}, {"v1", 0}},
		},
		{ // duplicate entries in removes + changes
			testSeqSet(3, 10),
			[]testSeq{{"v1", 0}, {"v3", 5}, {"v2", 20}, {"v2", 30}, {"v1", 0}},
		},
	}

	for i, tt := range testCases {
		executeSeqSetErrTestCase(t, i, tt)
	}
}

func TestSeqSetUpdatesOverflows(t *testing.T) {
	maxVP := MaxTotalPower
	testCases := []seqSetErrTestCase{
		{ // single update leading to overflow
			testSeqSet(2, 10),
			[]testSeq{{"v1", math.MaxInt64}},
		},
		{ // single update leading to overflow
			testSeqSet(2, 10),
			[]testSeq{{"v2", math.MaxInt64}},
		},
		{ // add seqidator leading to overflow
			testSeqSet(1, maxVP),
			[]testSeq{{"v2", math.MaxInt64}},
		},
		{ // add seqidator leading to exceed Max
			testSeqSet(1, maxVP-1),
			[]testSeq{{"v2", 5}},
		},
		{ // add seqidator leading to exceed Max
			testSeqSet(2, maxVP/3),
			[]testSeq{{"v3", maxVP / 2}},
		},
		{ // add seqidator leading to exceed Max
			testSeqSet(1, maxVP),
			[]testSeq{{"v2", maxVP}},
		},
	}

	for i, tt := range testCases {
		executeSeqSetErrTestCase(t, i, tt)
	}
}

func TestSeqSetUpdatesOtherErrors(t *testing.T) {
	testCases := []seqSetErrTestCase{
		{ // update with negative voting power
			testSeqSet(2, 10),
			[]testSeq{{"v1", -123}},
		},
		{ // update with negative voting power
			testSeqSet(2, 10),
			[]testSeq{{"v2", -123}},
		},
		{ // remove non-existing seqidator
			testSeqSet(2, 10),
			[]testSeq{{"v3", 0}},
		},
		{ // delete all seqidators
			[]testSeq{{"v1", 10}, {"v2", 20}, {"v3", 30}},
			[]testSeq{{"v1", 0}, {"v2", 0}, {"v3", 0}},
		},
	}

	for i, tt := range testCases {
		executeSeqSetErrTestCase(t, i, tt)
	}
}

func TestSeqSetUpdatesBasicTestsExecute(t *testing.T) {
	seqSetUpdatesBasicTests := []struct {
		startSeqs    []testSeq
		updateSeqs   []testSeq
		expectedSeqs []testSeq
	}{
		{ // no changes
			testSeqSet(2, 10),
			[]testSeq{},
			testSeqSet(2, 10),
		},
		{ // voting power changes
			testSeqSet(2, 10),
			[]testSeq{{"v2", 22}, {"v1", 11}},
			[]testSeq{{"v2", 22}, {"v1", 11}},
		},
		{ // add new seqidators
			[]testSeq{{"v2", 20}, {"v1", 10}},
			[]testSeq{{"v4", 40}, {"v3", 30}},
			[]testSeq{{"v4", 40}, {"v3", 30}, {"v2", 20}, {"v1", 10}},
		},
		{ // add new seqidator to middle
			[]testSeq{{"v3", 20}, {"v1", 10}},
			[]testSeq{{"v2", 30}},
			[]testSeq{{"v2", 30}, {"v3", 20}, {"v1", 10}},
		},
		{ // add new seqidator to beginning
			[]testSeq{{"v3", 20}, {"v2", 10}},
			[]testSeq{{"v1", 30}},
			[]testSeq{{"v1", 30}, {"v3", 20}, {"v2", 10}},
		},
		{ // delete seqidators
			[]testSeq{{"v3", 30}, {"v2", 20}, {"v1", 10}},
			[]testSeq{{"v2", 0}},
			[]testSeq{{"v3", 30}, {"v1", 10}},
		},
	}

	for i, tt := range seqSetUpdatesBasicTests {
		// create a new set and apply updates, keeping copies for the checks
		seqSet := createNewSequencerSet(tt.startSeqs)
		seqList := createNewSequencerList(tt.updateSeqs)
		err := seqSet.UpdateWithChangeSet(seqList)
		assert.NoError(t, err, "test %d", i)

		seqListCopy := sequencerListCopy(seqSet.Sequencers)
		// check that the voting power in the set's sequencers is not changing if the voting power
		// is changed in the list of sequencers previously passed as parameter to UpdateWithChangeSet.
		// this is to make sure copies of the sequencers are made by UpdateWithChangeSet.
		if len(seqList) > 0 {
			seqList[0].Power++
			assert.Equal(t, toTestSeqList(seqListCopy), toTestSeqList(seqSet.Sequencers), "test %v", i)

		}

		// check the final sequencer list is as expected and the set is properly scaled and centered.
		assert.Equal(t, renameTestSeqList(tt.expectedSeqs), toTestSeqList(seqSet.Sequencers), "test %v", i)
		verifySequencerSet(t, seqSet)
	}
}

// Test that different permutations of an update give the same result.
//func TestSeqSetUpdatesOrderIndependenceTestsExecute(t *testing.T) {
//	// startSeqs - initial sequencer to create the set with
//	// updateSeqs - a sequence of updates to be applied to the set.
//	// updateSeqs is shuffled a number of times during testing to check for same resulting sequencer set.
//	seqSetUpdatesOrderTests := []struct {
//		startSeqs  []testSeq
//		updateSeqs []testSeq
//	}{
//		0: { // order of changes should not matter, the final sequencer sets should be the same
//			[]testSeq{{"v4", 40}, {"v3", 30}, {"v2", 10}, {"v1", 10}},
//			[]testSeq{{"v4", 44}, {"v3", 33}, {"v2", 22}, {"v1", 11}}},
//
//		1: { // order of additions should not matter
//			[]testSeq{{"v2", 20}, {"v1", 10}},
//			[]testSeq{{"v3", 30}, {"v4", 40}, {"v5", 50}, {"v6", 60}}},
//
//		2: { // order of remoseqs should not matter
//			[]testSeq{{"v4", 40}, {"v3", 30}, {"v2", 20}, {"v1", 10}},
//			[]testSeq{{"v1", 0}, {"v3", 0}, {"v4", 0}}},
//
//		3: { // order of mixed operations should not matter
//			[]testSeq{{"v4", 40}, {"v3", 30}, {"v2", 20}, {"v1", 10}},
//			[]testSeq{{"v1", 0}, {"v3", 0}, {"v2", 22}, {"v5", 50}, {"v4", 44}}},
//	}
//
//	for i, tt := range seqSetUpdatesOrderTests {
//		// create a new set and apply updates
//		seqSet := createNewSequencerSet(tt.startSeqs)
//		seqSetCopy := seqSet.Copy()
//		seqList := createNewSequencerList(tt.updateSeqs)
//		assert.NoError(t, seqSetCopy.UpdateWithChangeSet(seqList))
//
//		// save the result as expected for next updates
//		seqSetExp := seqSetCopy.Copy()
//
//		// perform at most 20 permutations on the updates and call UpdateWithChangeSet()
//		n := len(tt.updateSeqs)
//		//maxNumPerms := int(math.Max(20, float64(n*n)))
//		for j := 0; j < n; j++ {
//			// create a copy of original set and apply a random permutation of updates
//			seqSetCopy := seqSet.Copy()
//			seqList := createNewSequencerList(permutation(tt.updateSeqs))
//
//			// check there was no error and the set is properly scaled and centered.
//			assert.NoError(t, seqSetCopy.UpdateWithChangeSet(seqList),
//				"test %v failed for permutation %v", i, seqList)
//			verifySequencerSet(t, seqSetCopy)
//
//			// verify the resulting test is same as the expected
//			assert.Equal(t, seqSetCopy, seqSetExp,
//				"test %v failed for permutation %v", i, seqList)
//		}
//	}
//}

// This tests the private function seqidator_set.go:applyUpdates() function, used only for additions and changes.
// Should perform a proper merge of updatedSeqs and startSeqs
func TestSeqSetApplyUpdatesTestsExecute(t *testing.T) {
	seqSetUpdatesBasicTests := []struct {
		startSeqs    []testSeq
		updateSeqs   []testSeq
		expectedSeqs []testSeq
	}{
		// additions
		0: { // prepend
			[]testSeq{{"v4", 44}, {"v5", 55}},
			[]testSeq{{"v1", 11}},
			[]testSeq{{"v1", 11}, {"v4", 44}, {"v5", 55}},
		},
		1: { // append
			[]testSeq{{"v4", 44}, {"v5", 55}},
			[]testSeq{{"v6", 66}},
			[]testSeq{{"v4", 44}, {"v5", 55}, {"v6", 66}},
		},
		2: { // insert
			[]testSeq{{"v4", 44}, {"v6", 66}},
			[]testSeq{{"v5", 55}},
			[]testSeq{{"v4", 44}, {"v5", 55}, {"v6", 66}},
		},
		3: { // insert multi
			[]testSeq{{"v4", 44}, {"v6", 66}, {"v9", 99}},
			[]testSeq{{"v5", 55}, {"v7", 77}, {"v8", 88}},
			[]testSeq{{"v4", 44}, {"v5", 55}, {"v6", 66}, {"v7", 77}, {"v8", 88}, {"v9", 99}},
		},
		// changes
		4: { // head
			[]testSeq{{"v1", 111}, {"v2", 22}},
			[]testSeq{{"v1", 11}},
			[]testSeq{{"v1", 11}, {"v2", 22}},
		},
		5: { // tail
			[]testSeq{{"v1", 11}, {"v2", 222}},
			[]testSeq{{"v2", 22}},
			[]testSeq{{"v1", 11}, {"v2", 22}},
		},
		6: { // middle
			[]testSeq{{"v1", 11}, {"v2", 222}, {"v3", 33}},
			[]testSeq{{"v2", 22}},
			[]testSeq{{"v1", 11}, {"v2", 22}, {"v3", 33}},
		},
		7: { // multi
			[]testSeq{{"v1", 111}, {"v2", 222}, {"v3", 333}},
			[]testSeq{{"v1", 11}, {"v2", 22}, {"v3", 33}},
			[]testSeq{{"v1", 11}, {"v2", 22}, {"v3", 33}},
		},
		// additions and changes
		8: {
			[]testSeq{{"v1", 111}, {"v2", 22}},
			[]testSeq{{"v1", 11}, {"v3", 33}, {"v4", 44}},
			[]testSeq{{"v1", 11}, {"v2", 22}, {"v3", 33}, {"v4", 44}},
		},
	}

	for i, tt := range seqSetUpdatesBasicTests {
		// create a new seqidator set with the start seques
		seqSet := createNewSequencerSet(tt.startSeqs)

		// applyUpdates() with the update seques
		seqList := createNewSequencerList(tt.updateSeqs)
		seqSet.applyUpdates(seqList)

		// check the new list of seqidators for proper merge
		assert.Equal(t, toTestSeqList(seqSet.Sequencers), renameTestSeqList(tt.expectedSeqs), "test %v", i)
	}
}

type testVSetCfg struct {
	name         string
	startSeqs    []testSeq
	deletedSeqs  []testSeq
	updatedSeqs  []testSeq
	addedSeqs    []testSeq
	expectedSeqs []testSeq
	expErr       error
}

//func randTestVSetCfg(t *testing.T, nBase, nAddMax int) testVSetCfg {
//	if nBase <= 0 || nAddMax < 0 {
//		panic(fmt.Sprintf("bad parameters %v %v", nBase, nAddMax))
//	}
//
//	const maxPower = 1000
//	var nOld, nDel, nChanged, nAdd int
//
//	nOld = int(tmrand.Uint()%uint(nBase)) + 1
//	if nBase-nOld > 0 {
//		nDel = int(tmrand.Uint() % uint(nBase-nOld))
//	}
//	nChanged = nBase - nOld - nDel
//
//	if nAddMax > 0 {
//		nAdd = tmrand.Int()%nAddMax + 1
//	}
//
//	cfg := testVSetCfg{}
//
//	cfg.startSeqs = make([]testSeq, nBase)
//	cfg.deletedSeqs = make([]testSeq, nDel)
//	cfg.addedSeqs = make([]testSeq, nAdd)
//	cfg.updatedSeqs = make([]testSeq, nChanged)
//	cfg.expectedSeqs = make([]testSeq, nBase-nDel+nAdd)
//
//	for i := 0; i < nBase; i++ {
//		cfg.startSeqs[i] = testSeq{fmt.Sprintf("v%d", i), int64(tmrand.Uint()%maxPower + 1)}
//		if i < nOld {
//			cfg.expectedSeqs[i] = cfg.startSeqs[i]
//		}
//		if i >= nOld && i < nOld+nChanged {
//			cfg.updatedSeqs[i-nOld] = testSeq{fmt.Sprintf("v%d", i), int64(tmrand.Uint()%maxPower + 1)}
//			cfg.expectedSeqs[i] = cfg.updatedSeqs[i-nOld]
//		}
//		if i >= nOld+nChanged {
//			cfg.deletedSeqs[i-nOld-nChanged] = testSeq{fmt.Sprintf("v%d", i), 0}
//		}
//	}
//
//	for i := nBase; i < nBase+nAdd; i++ {
//		cfg.addedSeqs[i-nBase] = testSeq{fmt.Sprintf("v%d", i), int64(tmrand.Uint()%maxPower + 1)}
//		cfg.expectedSeqs[i-nDel] = cfg.addedSeqs[i-nBase]
//	}
//
//	sort.Sort(testSeqsByVotingPower(cfg.startSeqs))
//	sort.Sort(testSeqsByVotingPower(cfg.deletedSeqs))
//	sort.Sort(testSeqsByVotingPower(cfg.updatedSeqs))
//	sort.Sort(testSeqsByVotingPower(cfg.addedSeqs))
//	sort.Sort(testSeqsByVotingPower(cfg.expectedSeqs))
//
//	return cfg
//
//}

func applyChangesToSeqSet(t *testing.T, expErr error, seqSet *SequencerSet, seqsLists ...[]testSeq) {
	changes := make([]testSeq, 0)
	for _, seqsList := range seqsLists {
		changes = append(changes, seqsList...)
	}
	seqList := createNewSequencerList(changes)
	err := seqSet.UpdateWithChangeSet(seqList)
	if expErr != nil {
		assert.Equal(t, expErr, err)
	} else {
		assert.NoError(t, err)
	}
}

//func TestValSetUpdatePriorityOrderTests(t *testing.T) {
//	const nMaxElections int32 = 5000
//
//	testCases := []testVSetCfg{
//		0: { // remove high power seqidator, keep old equal lower power seqidators
//			startSeqs:    []testSeq{{"v3", 1000}, {"v1", 1}, {"v2", 1}},
//			deletedSeqs:  []testSeq{{"v3", 0}},
//			updatedSeqs:  []testSeq{},
//			addedSeqs:    []testSeq{},
//			expectedSeqs: []testSeq{{"v1", 1}, {"v2", 1}},
//		},
//		1: { // remove high power seqidator, keep old different power seqidators
//			startSeqs:    []testSeq{{"v3", 1000}, {"v2", 10}, {"v1", 1}},
//			deletedSeqs:  []testSeq{{"v3", 0}},
//			updatedSeqs:  []testSeq{},
//			addedSeqs:    []testSeq{},
//			expectedSeqs: []testSeq{{"v2", 10}, {"v1", 1}},
//		},
//		2: { // remove high power seqidator, add new low power seqidators, keep old lower power
//			startSeqs:    []testSeq{{"v3", 1000}, {"v2", 2}, {"v1", 1}},
//			deletedSeqs:  []testSeq{{"v3", 0}},
//			updatedSeqs:  []testSeq{{"v2", 1}},
//			addedSeqs:    []testSeq{{"v5", 50}, {"v4", 40}},
//			expectedSeqs: []testSeq{{"v5", 50}, {"v4", 40}, {"v1", 1}, {"v2", 1}},
//		},
//
//		// generate a configuration with 100 seqidators,
//		// randomly select seqidators for updates and deletes, and
//		// generate 10 new seqidators to be added
//		3: randTestVSetCfg(t, 100, 10),
//
//		4: randTestVSetCfg(t, 1000, 100),
//
//		5: randTestVSetCfg(t, 10, 100),
//
//		6: randTestVSetCfg(t, 100, 1000),
//
//		7: randTestVSetCfg(t, 1000, 1000),
//	}
//
//	for _, cfg := range testCases {
//
//		// create a new seqidator set
//		seqSet := createNewSequencerSet(cfg.startSeqs)
//		verifySequencerSet(t, seqSet)
//
//		// run election up to nMaxElections times, apply changes and verify that the priority order is correct
//		verifyValSetUpdatePriorityOrder(t, seqSet, cfg, nMaxElections)
//	}
//}

func verifySeqSetUpdatePriorityOrder(t *testing.T, seqSet *SequencerSet, cfg testVSetCfg, nMaxElections int32) {
	// Run election up to nMaxElections times, sort seqidators by priorities
	seqSet.IncrementProducerPriority(rand.Int31()%nMaxElections + 1)

	// apply the changes, get the updated seqidators, sort by priorities
	applyChangesToSeqSet(t, nil, seqSet, cfg.addedSeqs, cfg.updatedSeqs, cfg.deletedSeqs)

	// basic checks
	assert.Equal(t, cfg.expectedSeqs, toTestSeqList(seqSet.Sequencers))
	verifySequencerSet(t, seqSet)

	// verify that the added seqidators have the smallest priority:
	//  - they should be at the beginning of updatedSeqsPriSorted since it is
	//  sorted by priority
	if len(cfg.addedSeqs) > 0 {
		updatedSeqsPriSorted := sequencerListCopy(seqSet.Sequencers)
		sort.Sort(sequencersByPriority(updatedSeqsPriSorted))

		addedSeqsPriSlice := updatedSeqsPriSorted[:len(cfg.addedSeqs)]
		sort.Sort(SequencersByPower(addedSeqsPriSlice))
		assert.Equal(t, renameTestSeqList(cfg.addedSeqs), toTestSeqList(addedSeqsPriSlice))

		//  - and should all have the same priority
		expectedPri := addedSeqsPriSlice[0].ProducerPriority
		for _, seq := range addedSeqsPriSlice[1:] {
			assert.Equal(t, expectedPri, seq.ProducerPriority)
		}
	}
}

//func TestNewSequencerSetFromExistingSequencers(t *testing.T) {
//	size := 5
//	seqs := make([]*Sequencer, size)
//	for i := 0; i < size; i++ {
//		pv := NewMockPV()
//		seqs[i] = pv.ExtractIntoSequencer(int64(i + 1))
//	}
//	seqSet := NewSequencerSet(seqs)
//	seqSet.IncrementProducerPriority(5)
//
//	newValSet := NewSequencerSet(seqSet.Sequencers)
//	assert.NotEqual(t, seqSet, newValSet)
//
//	existingValSet, err := SequencerSetFromExistingSequencers(seqSet.Sequencers)
//	assert.NoError(t, err)
//	assert.Equal(t, seqSet, existingValSet)
//	assert.Equal(t, seqSet.CopyIncrementProducerPriority(3), existingValSet.CopyIncrementProducerPriority(3))
//}

func TestSeqSetUpdateOverflowRelated(t *testing.T) {
	testCases := []testVSetCfg{
		{
			name:         "1 no false overflow error messages for updates",
			startSeqs:    []testSeq{{"v2", MaxTotalPower - 1}, {"v1", 1}},
			updatedSeqs:  []testSeq{{"v1", MaxTotalPower - 1}, {"v2", 1}},
			expectedSeqs: []testSeq{{"v1", MaxTotalPower - 1}, {"v2", 1}},
			expErr:       nil,
		},
		{
			// this test shows that it is important to apply the updates in the order of the change in power
			// i.e. apply first updates with decreases in power, v2 change in this case.
			name:         "2 no false overflow error messages for updates",
			startSeqs:    []testSeq{{"v2", MaxTotalPower - 1}, {"v1", 1}},
			updatedSeqs:  []testSeq{{"v1", MaxTotalPower/2 - 1}, {"v2", MaxTotalPower / 2}},
			expectedSeqs: []testSeq{{"v2", MaxTotalPower / 2}, {"v1", MaxTotalPower/2 - 1}},
			expErr:       nil,
		},
		{
			name:         "3 no false overflow error messages for deletes",
			startSeqs:    []testSeq{{"v1", MaxTotalPower - 2}, {"v2", 1}, {"v3", 1}},
			deletedSeqs:  []testSeq{{"v1", 0}},
			addedSeqs:    []testSeq{{"v4", MaxTotalPower - 2}},
			expectedSeqs: []testSeq{{"v4", MaxTotalPower - 2}, {"v2", 1}, {"v3", 1}},
			expErr:       nil,
		},
		{
			name: "4 no false overflow error messages for adds, updates and deletes",
			startSeqs: []testSeq{
				{"v1", MaxTotalPower / 4},
				{"v2", MaxTotalPower / 4},
				{"v3", MaxTotalPower / 4},
				{"v4", MaxTotalPower / 4},
			},
			deletedSeqs: []testSeq{{"v2", 0}},
			updatedSeqs: []testSeq{
				{"v1", MaxTotalPower/2 - 2}, {"v3", MaxTotalPower/2 - 3}, {"v4", 2},
			},
			addedSeqs: []testSeq{{"v5", 3}},
			expectedSeqs: []testSeq{
				{"v1", MaxTotalPower/2 - 2}, {"v3", MaxTotalPower/2 - 3}, {"v5", 3}, {"v4", 2},
			},
			expErr: nil,
		},
		{
			name: "5 check panic on overflow is prevented: update 8 seqidators with power int64(math.MaxInt64)/8",
			startSeqs: []testSeq{
				{"v1", 1},
				{"v2", 1},
				{"v3", 1},
				{"v4", 1},
				{"v5", 1},
				{"v6", 1},
				{"v7", 1},
				{"v8", 1},
				{"v9", 1},
			},
			updatedSeqs: []testSeq{
				{"v1", MaxTotalPower},
				{"v2", MaxTotalPower},
				{"v3", MaxTotalPower},
				{"v4", MaxTotalPower},
				{"v5", MaxTotalPower},
				{"v6", MaxTotalPower},
				{"v7", MaxTotalPower},
				{"v8", MaxTotalPower},
				{"v9", 8},
			},
			expectedSeqs: []testSeq{
				{"v1", 1},
				{"v2", 1},
				{"v3", 1},
				{"v4", 1},
				{"v5", 1},
				{"v6", 1},
				{"v7", 1},
				{"v8", 1},
				{"v9", 1},
			},
			expErr: ErrTotalPowerOverflow,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			seqSet := createNewSequencerSet(tt.startSeqs)
			verifySequencerSet(t, seqSet)

			// execute update and verify returned error is as expected
			applyChangesToSeqSet(t, tt.expErr, seqSet, tt.addedSeqs, tt.updatedSeqs, tt.deletedSeqs)

			// verify updated seqidator set is as expected
			assert.Equal(t, renameTestSeqList(tt.expectedSeqs), toTestSeqList(seqSet.Sequencers))
			verifySequencerSet(t, seqSet)
		})
	}
}

func TestSafeMul(t *testing.T) {
	testCases := []struct {
		a        int64
		b        int64
		c        int64
		overflow bool
	}{
		0: {0, 0, 0, false},
		1: {1, 0, 0, false},
		2: {2, 3, 6, false},
		3: {2, -3, -6, false},
		4: {-2, -3, 6, false},
		5: {-2, 3, -6, false},
		6: {math.MaxInt64, 1, math.MaxInt64, false},
		7: {math.MaxInt64 / 2, 2, math.MaxInt64 - 1, false},
		8: {math.MaxInt64 / 2, 3, 0, true},
		9: {math.MaxInt64, 2, 0, true},
	}

	for i, tc := range testCases {
		c, overflow := safeMul(tc.a, tc.b)
		assert.Equal(t, tc.c, c, "#%d", i)
		assert.Equal(t, tc.overflow, overflow, "#%d", i)
	}
}

// ---------------------
// Sort sequencers by priority and address
type sequencersByPriority []*Sequencer

func (seqz sequencersByPriority) Len() int {
	return len(seqz)
}

func (seqz sequencersByPriority) Less(i, j int) bool {
	if seqz[i].ProducerPriority < seqz[j].ProducerPriority {
		return true
	}
	if seqz[i].ProducerPriority > seqz[j].ProducerPriority {
		return false
	}
	return bytes.Compare(seqz[i].Address.Bytes(), seqz[j].Address.Bytes()) < 0
}

func (seqz sequencersByPriority) Swap(i, j int) {
	seqz[i], seqz[j] = seqz[j], seqz[i]
}

//-------------------------------------

type testSeqsByVotingPower []testSeq

func (tseqs testSeqsByVotingPower) Len() int {
	return len(tseqs)
}

func (tseqs testSeqsByVotingPower) Less(i, j int) bool {
	if tseqs[i].power == tseqs[j].power {
		return bytes.Compare([]byte(tseqs[i].name), []byte(tseqs[j].name)) == -1
	}
	return tseqs[i].power > tseqs[j].power
}

func (tseqs testSeqsByVotingPower) Swap(i, j int) {
	tseqs[i], tseqs[j] = tseqs[j], tseqs[i]
}

// -------------------------------------
// Benchmark tests
func BenchmarkUpdates(b *testing.B) {
	const (
		n = 100
		m = 2000
	)
	// Init with n sequencers
	vs := make([]*Sequencer, n)
	for j := 0; j < n; j++ {
		vs[j] = newSequencer(byteToCommonAddr([]byte(fmt.Sprintf("v%d", j))), randPubKey(), 100)
	}
	seqSet := NewSequencerSet(vs)
	l := len(seqSet.Sequencers)

	// Make m new sequencer
	newValList := make([]*Sequencer, m)
	for j := 0; j < m; j++ {
		newValList[j] = newSequencer(byteToCommonAddr([]byte(fmt.Sprintf("v%d", j+l))), randPubKey(), 1000)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Add m sequencer to seqSetCopy
		seqSetCopy := seqSet.Copy()
		assert.NoError(b, seqSetCopy.UpdateWithChangeSet(newValList))
	}
}
