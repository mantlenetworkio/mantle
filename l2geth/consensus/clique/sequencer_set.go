package clique

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
)

const (
	// MaxTotalPower - the maximum allowed total power.
	// It needs to be sufficiently small to, in all cases:
	// 1. prevent clipping in incrementProposerPriority()
	// 2. let (diff+diffMax-1) not overflow in IncrementProposerPriority()
	// (Proof of 1 is tricky, left to the reader).
	// It could be higher, but this is sufficiently large for our purposes,
	// and leaves room for defensive purposes.
	MaxTotalPower = int64(math.MaxInt64) / 8

	// PriorityWindowSizeFactor - is a constant that when multiplied with the
	// total power gives the maximum allowed distance between sequencer
	// priorities.
	PriorityWindowSizeFactor = 2
)

// ErrTotalPowerOverflow is returned if the total power of the
// resulting sequencer set exceeds MaxTotalPower.
var ErrTotalPowerOverflow = fmt.Errorf("total power of resulting seqset exceeds max %d",
	MaxTotalPower)

// SequencerSet represent a set of *Sequencer at a given height.
//
// The sequencers can be fetched by address or index.
// The index is in order of .Power, so the indices are fixed for all
// rounds of a given blockchain height - ie. the sequencers are sorted by their
// power (descending). Secondary index - .Address (ascending).
//
// On the other hand, the .ProposerPriority of each sequencer and the
// designated .GetProposer() of a set changes every round, upon calling
// .IncrementProposerPriority().
//
// NOTE: Not goroutine-safe.
// NOTE: All get/set to sequencers should copy the value for safety.
type SequencerSet struct {
	// NOTE: persisted via reflect, must be exported.
	Sequencers []*Sequencer `json:"sequencers"`
	Producer   *Sequencer   `json:"producer"`

	// cached (unexported)
	totalPower int64
}

// NewSequencerSet initializes a SequencerSet by copying over the values from
// `seqz`, a list of Sequencers. If seqz is nil or empty, the new SequencerSet
// will have an empty list of Sequencers.
//
// The addresses of Sequencers in `seqz` must be unique otherwise the function
// panics.
//
// Note the sequencer set size has an implied limit equal to that of the
// MaxVotesCount - commits by a sequencer set larger than this will fail
// validation.
func NewSequencerSet(seqz []*Sequencer) *SequencerSet {
	seqs := &SequencerSet{}
	err := seqs.updateWithChangeSet(seqz, false)
	if err != nil {
		panic(fmt.Sprintf("Cannot create sequencer set: %v", err))
	}
	if len(seqz) > 0 {
		seqs.IncrementProducerPriority(1)
	}
	return seqs
}

func (seqs *SequencerSet) ValidateBasic() error {
	if seqs.IsNilOrEmpty() {
		return errors.New("sequencer set is nil or empty")
	}

	for idx, seq := range seqs.Sequencers {
		if err := seq.SequencerBasic(); err != nil {
			return fmt.Errorf("invalid sequencer #%d: %w", idx, err)
		}
	}

	if err := seqs.Producer.SequencerBasic(); err != nil {
		return fmt.Errorf("producer failed validate basic, error: %w", err)
	}

	return nil
}

// IsNilOrEmpty returns true if sequencer set is nil or empty.
func (seqs *SequencerSet) IsNilOrEmpty() bool {
	return seqs == nil || len(seqs.Sequencers) == 0
}

// CopyIncrementProducerPriority increments ProducerPriority and updates the
// producer on a copy, and returns it.
func (seqs *SequencerSet) CopyIncrementProducerPriority(times int32) *SequencerSet {
	copy := seqs.Copy()
	copy.IncrementProducerPriority(times)
	return copy
}

// IncrementProposerPriority increments ProducerPriority of each sequencer and
// updates the producer. Panics if sequencer set is empty.
// `times` must be positive.
func (seqs *SequencerSet) IncrementProducerPriority(times int32) {
	if seqs.IsNilOrEmpty() {
		panic("empty sequencer set")
	}
	if times <= 0 {
		panic("Cannot call IncrementProducerPriority with non-positive times")
	}

	// Cap the difference between priorities to be proportional to 2*totalPower by
	// re-normalizing priorities, i.e., rescale all priorities by multiplying with:
	//  2*totalPower/(maxPriority - minPriority)
	diffMax := PriorityWindowSizeFactor * seqs.TotalPower()
	seqs.RescalePriorities(diffMax)
	seqs.shiftByAvgProducerPriority()

	var producer *Sequencer
	// Call IncrementProposerPriority(1) times times.
	for i := int32(0); i < times; i++ {
		producer = seqs.incrementProducerPriority()
	}

	seqs.Producer = producer
}

// RescalePriorities rescales the priorities such that the distance between the
// maximum and minimum is smaller than `diffMax`. Panics if sequencer set is
// empty.
func (seqs *SequencerSet) RescalePriorities(diffMax int64) {
	if seqs.IsNilOrEmpty() {
		panic("empty sequencer set")
	}
	// NOTE: This check is merely a sanity check which could be
	// removed if all tests would init. power appropriately;
	// i.e. diffMax should always be > 0
	if diffMax <= 0 {
		return
	}

	// Calculating ceil(diff/diffMax):
	// Re-normalization is performed by dividing by an integer for simplicity.
	// NOTE: This may make debugging priority issues easier as well.
	diff := computeMaxMinPriorityDiff(seqs)
	ratio := (diff + diffMax - 1) / diffMax
	if diff > diffMax {
		for _, val := range seqs.Sequencers {
			val.ProducerPriority /= ratio
		}
	}
}

func (seqs *SequencerSet) incrementProducerPriority() *Sequencer {
	for _, seq := range seqs.Sequencers {
		// Check for overflow for sum.
		newPriority := safeAddClip(seq.ProducerPriority, seq.Power)
		seq.ProducerPriority = newPriority
	}
	// Decrement the sequencer with most ProposerPriority.
	mostest := seqs.getSeqWithMostPriority()
	// Mind the underflow.
	mostest.ProducerPriority = safeSubClip(mostest.ProducerPriority, seqs.TotalPower())

	return mostest
}

// Should not be called on an empty sequencer set.
func (seqs *SequencerSet) computeAvgProducerPriority() int64 {
	n := int64(len(seqs.Sequencers))
	sum := big.NewInt(0)
	for _, seq := range seqs.Sequencers {
		sum.Add(sum, big.NewInt(seq.ProducerPriority))
	}
	avg := sum.Div(sum, big.NewInt(n))
	if avg.IsInt64() {
		return avg.Int64()
	}

	// This should never happen: each val.ProposerPriority is in bounds of int64.
	panic(fmt.Sprintf("Cannot represent avg ProposerPriority as an int64 %v", avg))
}

// Compute the difference between the max and min ProposerPriority of that set.
func computeMaxMinPriorityDiff(seqs *SequencerSet) int64 {
	if seqs.IsNilOrEmpty() {
		panic("empty sequencer set")
	}
	max := int64(math.MinInt64)
	min := int64(math.MaxInt64)
	for _, v := range seqs.Sequencers {
		if v.ProducerPriority < min {
			min = v.ProducerPriority
		}
		if v.ProducerPriority > max {
			max = v.ProducerPriority
		}
	}
	diff := max - min
	if diff < 0 {
		return -1 * diff
	}
	return diff
}

func (seqs *SequencerSet) getSeqWithMostPriority() *Sequencer {
	var res *Sequencer
	for _, seq := range seqs.Sequencers {
		res = res.CompareProducerPriority(seq)
	}
	return res
}

func (seqs *SequencerSet) shiftByAvgProducerPriority() {
	if seqs.IsNilOrEmpty() {
		panic("empty sequencer set")
	}
	avgProposerPriority := seqs.computeAvgProducerPriority()
	for _, val := range seqs.Sequencers {
		val.ProducerPriority = safeSubClip(val.ProducerPriority, avgProposerPriority)
	}
}

// Makes a copy of the sequencer list.
func sequencerListCopy(seqsList []*Sequencer) []*Sequencer {
	if seqsList == nil {
		return nil
	}
	seqsCopy := make([]*Sequencer, len(seqsList))
	for i, val := range seqsList {
		seqsCopy[i] = val.Copy()
	}
	return seqsCopy
}

// Copy each sequencer into a new SequencerSet.
func (seqs *SequencerSet) Copy() *SequencerSet {
	return &SequencerSet{
		Sequencers: sequencerListCopy(seqs.Sequencers),
		Producer:   seqs.Producer,
		totalPower: seqs.totalPower,
	}
}

// HasAddress returns true if address given is in the sequencer set, false -
// otherwise.
func (seqs *SequencerSet) HasAddress(address common.Address) bool {
	for _, seq := range seqs.Sequencers {
		if bytes.Equal(seq.Address.Bytes(), address.Bytes()) {
			return true
		}
	}
	return false
}

// GetByAddress returns an index of the sequencer with address and sequencer
// itself (copy) if found. Otherwise, -1 and nil are returned.
func (seqs *SequencerSet) GetByAddress(address common.Address) (index int32, seq *Sequencer) {
	for idx, seq := range seqs.Sequencers {
		if bytes.Equal(seq.Address.Bytes(), address.Bytes()) {
			return int32(idx), seq.Copy()
		}
	}
	return -1, nil
}

// GetByIndex returns the sequencer's address and sequencer itself (copy) by
// index.
// It returns nil values if index is less than 0 or greater or equal to
// len(SequencerSet.Sequencers).
func (seqs *SequencerSet) GetByIndex(index int32) (address common.Address, seq *Sequencer) {
	if index < 0 || int(index) >= len(seqs.Sequencers) {
		return [20]byte{}, nil
	}
	seq = seqs.Sequencers[index]
	return seq.Address, seq.Copy()
}

// Size returns the length of the sequencer set.
func (seqs *SequencerSet) Size() int {
	return len(seqs.Sequencers)
}

// Forces recalculation of the set's total power.
// Panics if total power is bigger than MaxTotalPower.
func (seqs *SequencerSet) updateTotalPower() {
	sum := int64(0)
	for _, val := range seqs.Sequencers {
		// mind overflow
		sum = safeAddClip(sum, val.Power)
		if sum > MaxTotalPower {
			panic(fmt.Sprintf(
				"Total power should be guarded to not exceed %v; got: %v",
				MaxTotalPower,
				sum))
		}
	}

	seqs.totalPower = sum
}

// TotalPower returns the sum of the powers of all sequencers.
// It recomputes the total power if required.
func (seqs *SequencerSet) TotalPower() int64 {
	if seqs.totalPower == 0 {
		seqs.updateTotalPower()
	}
	return seqs.totalPower
}

// GetProposer returns the current proposer. If the sequencer set is empty, nil
// is returned.
func (seqs *SequencerSet) GetProducer() (producer *Sequencer) {
	if len(seqs.Sequencers) == 0 {
		return nil
	}
	if seqs.Producer == nil {
		seqs.Producer = seqs.findProducer()
	}
	return seqs.Producer.Copy()
}

func (seqs *SequencerSet) findProducer() *Sequencer {
	var producer *Sequencer
	for _, seq := range seqs.Sequencers {
		if producer == nil || !bytes.Equal(seq.Address.Bytes(), producer.Address.Bytes()) {
			producer = producer.CompareProducerPriority(seq)
		}
	}
	return producer
}

// Hash returns the Merkle root hash build using sequencers (as leaves) in the
// set.
func (seqs *SequencerSet) Hash() common.Hash {
	return types.DeriveSha(seqs)
}

// Iterate will run the given function over the set.
func (seqs *SequencerSet) Iterate(fn func(index int, val *Sequencer) bool) {
	for i, val := range seqs.Sequencers {
		stop := fn(i, val.Copy())
		if stop {
			break
		}
	}
}

// Checks changes against duplicates, splits the changes in updates and
// removals, sorts them by address.
//
// Returns:
// updates, removals - the sorted lists of updates and removals
// err - non-nil if duplicate entries or entries with negative power are seen
//
// No changes are made to 'origChanges'.
func processChanges(origChanges []*Sequencer) (updates, removals []*Sequencer, err error) {
	// Make a deep copy of the changes and sort by address.
	changes := sequencerListCopy(origChanges)
	sort.Sort(SequencersByAddress(changes))

	removals = make([]*Sequencer, 0, len(changes))
	updates = make([]*Sequencer, 0, len(changes))
	var prevAddr common.Address
	// Scan changes by address and append valid sequencers to updates or removals lists.
	for _, seqUpdate := range changes {
		if bytes.Equal(seqUpdate.Address.Bytes(), prevAddr.Bytes()) {
			err = fmt.Errorf("duplicate entry %v in %v", seqUpdate, changes)
			return nil, nil, err
		}

		switch {
		case seqUpdate.Power < 0:
			err = fmt.Errorf("power can't be negative: %d", seqUpdate.Power)
			return nil, nil, err
		case seqUpdate.Power > MaxTotalPower:
			err = fmt.Errorf("to prevent clipping/overflow, power can't be higher than %d, got %d",
				MaxTotalPower, seqUpdate.Power)
			return nil, nil, err
		case seqUpdate.Power == 0:
			removals = append(removals, seqUpdate)
		default:
			updates = append(updates, seqUpdate)
		}

		prevAddr = seqUpdate.Address
	}

	return updates, removals, err
}

// verifyUpdates verifies a list of updates against a sequencer set, making sure the allowed
// total power would not be exceeded if these updates would be applied to the set.
//
// Inputs:
// updates - a list of proper sequencer changes, i.e. they have been verified by processChanges for duplicates
//
//	and invalid values.
//
// seqs - the original sequencer set. Note that seqs is NOT modified by this function.
// removedPower - the total power that will be removed after the updates are verified and applied.
//
// Returns:
// tvpAfterUpdatesBeforeRemovals -  the new total power if these updates would be applied without the removals.
//
//	Note that this will be < 2 * MaxTotalPower in case high power sequencers are removed and
//	sequencers are added/ updated with high power values.
//
// err - non-nil if the maximum allowed total power would be exceeded
func verifyUpdates(
	updates []*Sequencer,
	seqs *SequencerSet,
	removedPower int64,
) (tvpAfterUpdatesBeforeRemovals int64, err error) {
	delta := func(update *Sequencer, seqs *SequencerSet) int64 {
		_, seq := seqs.GetByAddress(update.Address)
		if seq != nil {
			return update.Power - seq.Power
		}
		return update.Power
	}

	updatesCopy := sequencerListCopy(updates)
	sort.Slice(updatesCopy, func(i, j int) bool {
		return delta(updatesCopy[i], seqs) < delta(updatesCopy[j], seqs)
	})

	tvpAfterRemovals := seqs.TotalPower() - removedPower
	for _, upd := range updatesCopy {
		tvpAfterRemovals += delta(upd, seqs)
		if tvpAfterRemovals > MaxTotalPower {
			return 0, ErrTotalPowerOverflow
		}
	}
	return tvpAfterRemovals + removedPower, nil
}

func numNewSequencers(updates []*Sequencer, seqs *SequencerSet) int {
	numNewSequencers := 0
	for _, valUpdate := range updates {
		if !seqs.HasAddress(valUpdate.Address) {
			numNewSequencers++
		}
	}
	return numNewSequencers
}

// computeNewPriorities computes the proposer priority for the sequencers not present in the set based on
// 'updatedTotalPower'.
// Leaves unchanged the priorities of sequencers that are changed.
//
// 'updates' parameter must be a list of unique sequencers to be added or updated.
//
// 'updatedTotalPower' is the total  power of a set where all updates would be applied but
//
//	not the removals. It must be < 2*MaxTotalPower and may be close to this limit if close to
//	MaxTotalPower will be removed. This is still safe from overflow since MaxTotalPower is maxInt64/8.
//
// No changes are made to the sequencer set 'seqs'.
func computeNewPriorities(updates []*Sequencer, seqs *SequencerSet, updatedTotalPower int64) {
	for _, valUpdate := range updates {
		address := valUpdate.Address
		_, val := seqs.GetByAddress(address)
		if val == nil {
			// add val
			// Set ProposerPriority to -C*totalPower (with C ~= 1.125) to make sure sequencers can't
			// un-bond and then re-bond to reset their (potentially previously negative) ProposerPriority to zero.
			//
			// Contract: updatedPower < 2 * MaxTotalPower to ensure ProposerPriority does
			// not exceed the bounds of int64.
			//
			// Compute ProposerPriority = -1.125*totalPower == -(updatedPower + (updatedPower >> 3)).
			valUpdate.ProducerPriority = -(updatedTotalPower + (updatedTotalPower >> 3))
		} else {
			valUpdate.ProducerPriority = val.ProducerPriority
		}
	}
}

// Merges the seqs' sequencer list with the updates list.
// When two elements with same address are seen, the one from updates is selected.
// Expects updates to be a list of updates sorted by address with no duplicates or errors,
// must have been validated with verifyUpdates() and priorities computed with computeNewPriorities().
func (seqs *SequencerSet) applyUpdates(updates []*Sequencer) {
	existing := seqs.Sequencers
	sort.Sort(SequencersByAddress(existing))

	merged := make([]*Sequencer, len(existing)+len(updates))
	i := 0

	for len(existing) > 0 && len(updates) > 0 {
		if bytes.Compare(existing[0].Address.Bytes(), updates[0].Address.Bytes()) < 0 { // unchanged sequencer
			merged[i] = existing[0]
			existing = existing[1:]
		} else {
			// Apply to add or update.
			merged[i] = updates[0]
			if bytes.Equal(existing[0].Address.Bytes(), updates[0].Address.Bytes()) {
				// Sequencer is present in both, advance existing.
				existing = existing[1:]
			}
			updates = updates[1:]
		}
		i++
	}

	// Add the elements which are left.
	for j := 0; j < len(existing); j++ {
		merged[i] = existing[j]
		i++
	}
	// OR add updates which are left.
	for j := 0; j < len(updates); j++ {
		merged[i] = updates[j]
		i++
	}

	seqs.Sequencers = merged[:i]
}

// Checks that the sequencers to be removed are part of the sequencer set.
// No changes are made to the sequencer set 'seqs'.
func verifyRemovals(deletes []*Sequencer, seqs *SequencerSet) (power int64, err error) {
	removedPower := int64(0)
	for _, valUpdate := range deletes {
		address := valUpdate.Address
		_, val := seqs.GetByAddress(address)
		if val == nil {
			return removedPower, fmt.Errorf("failed to find sequencer %X to remove", address)
		}
		removedPower += val.Power
	}
	if len(deletes) > len(seqs.Sequencers) {
		panic("more deletes than sequencers")
	}
	return removedPower, nil
}

// Removes the sequencers specified in 'deletes' from sequencer set 'seqs'.
// Should not fail as verification has been done before.
// Expects seqs to be sorted by address (done by applyUpdates).
func (seqs *SequencerSet) applyRemovals(deletes []*Sequencer) {
	existing := seqs.Sequencers

	merged := make([]*Sequencer, len(existing)-len(deletes))
	i := 0

	// Loop over deletes until we removed all of them.
	for len(deletes) > 0 {
		if bytes.Equal(existing[0].Address.Bytes(), deletes[0].Address.Bytes()) {
			deletes = deletes[1:]
		} else { // Leave it in the resulting slice.
			merged[i] = existing[0]
			i++
		}
		existing = existing[1:]
	}

	// Add the elements which are left.
	for j := 0; j < len(existing); j++ {
		merged[i] = existing[j]
		i++
	}

	seqs.Sequencers = merged[:i]
}

// Main function used by UpdateWithChangeSet() and NewSequencerSet().
// If 'allowDeletes' is false then delete operations (identified by sequencers with power 0)
// are not allowed and will trigger an error if present in 'changes'.
// The 'allowDeletes' flag is set to false by NewSequencerSet() and to true by UpdateWithChangeSet().
func (seqs *SequencerSet) updateWithChangeSet(changes []*Sequencer, allowDeletes bool) error {
	if len(changes) == 0 {
		return nil
	}

	// Check for duplicates within changes, split in 'updates' and 'deletes' lists (sorted).
	updates, deletes, err := processChanges(changes)
	if err != nil {
		return err
	}

	if !allowDeletes && len(deletes) != 0 {
		return fmt.Errorf("cannot process sequencers with power 0: %v", deletes)
	}

	// Check that the resulting set will not be empty.
	if numNewSequencers(updates, seqs) == 0 && len(seqs.Sequencers) == len(deletes) {
		return errors.New("applying the sequencer changes would result in empty set")
	}

	// Verify that applying the 'deletes' against 'seqs' will not result in error.
	// Get the power that is going to be removed.
	removedPower, err := verifyRemovals(deletes, seqs)
	if err != nil {
		return err
	}

	// Verify that applying the 'updates' against 'seqs' will not result in error.
	// Get the updated total power before removal. Note that this is < 2 * MaxTotalPower
	tvpAfterUpdatesBeforeRemovals, err := verifyUpdates(updates, seqs, removedPower)
	if err != nil {
		return err
	}

	// Compute the priorities for updates.
	computeNewPriorities(updates, seqs, tvpAfterUpdatesBeforeRemovals)

	// Apply updates and removals.
	seqs.applyUpdates(updates)
	seqs.applyRemovals(deletes)

	seqs.updateTotalPower() // will panic if total power > MaxTotalPower

	// Scale and center.
	seqs.RescalePriorities(PriorityWindowSizeFactor * seqs.TotalPower())
	seqs.shiftByAvgProducerPriority()

	sort.Sort(SequencersByPower(seqs.Sequencers))

	return nil
}

// UpdateWithChangeSet attempts to update the sequencer set with 'changes'.
// It performs the following steps:
//   - validates the changes making sure there are no duplicates and splits them in updates and deletes
//   - verifies that applying the changes will not result in errors
//   - computes the total power BEFORE removals to ensure that in the next steps the priorities
//     across old and newly added sequencers are fair
//   - computes the priorities of new sequencers against the final set
//   - applies the updates against the sequencer set
//   - applies the removals against the sequencer set
//   - performs scaling and centering of priority values
//
// If an error is detected during verification steps, it is returned and the sequencer set
// is not changed.
func (seqs *SequencerSet) UpdateWithChangeSet(changes []*Sequencer) error {
	return seqs.updateWithChangeSet(changes, true)
}

// findPreviousProposer reverses the compare proposer priority function to find the sequencer
// with the lowest proposer priority which would have been the previous proposer.
//
// Is used when recreating a sequencer set from an existing array of sequencers.
func (seqs *SequencerSet) findPreviousProducer() *Sequencer {
	var previousProposer *Sequencer
	for _, val := range seqs.Sequencers {
		if previousProposer == nil {
			previousProposer = val
			continue
		}
		if previousProposer == previousProposer.CompareProducerPriority(val) {
			previousProposer = val
		}
	}
	return previousProposer
}

//----------------

// String returns a string representation of SequencerSet.
//
// See StringIndented.
func (seqs *SequencerSet) String() string {
	return seqs.StringIndented("")
}

// StringIndented returns an intended String.
//
// See Sequencer#String.
func (seqs *SequencerSet) StringIndented(indent string) string {
	if seqs == nil {
		return "nil-SequencerSet"
	}
	var valStrings []string
	seqs.Iterate(func(index int, val *Sequencer) bool {
		valStrings = append(valStrings, val.String())
		return false
	})
	return fmt.Sprintf(`SequencerSet{
%s  Proposer: %v
%s  Sequencers:
%s    %v
%s}`,
		indent, seqs.GetProducer().String(),
		indent,
		indent, strings.Join(valStrings, "\n"+indent+"    "),
		indent)
}

// Len returns the length of s.
func (seqs *SequencerSet) Len() int { return len(seqs.Sequencers) }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (seqs *SequencerSet) GetRlp(i int) []byte {
	_, seq := seqs.GetByIndex(int32(i))
	enc, _ := rlp.EncodeToBytes(seq)
	return enc
}

//-------------------------------------

// SequencersByPower implements sort.Interface for []*Sequencer based on
// the Power and Address fields.
type SequencersByPower []*Sequencer

func (seqz SequencersByPower) Len() int { return len(seqz) }

func (seqz SequencersByPower) Less(i, j int) bool {
	if seqz[i].Power == seqz[j].Power {
		return bytes.Compare(seqz[i].Address.Bytes(), seqz[j].Address.Bytes()) == -1
	}
	return seqz[i].Power > seqz[j].Power
}

func (seqz SequencersByPower) Swap(i, j int) {
	seqz[i], seqz[j] = seqz[j], seqz[i]
}

// SequencersByAddress implements sort.Interface for []*Sequencer based on
// the Address field.
type SequencersByAddress []*Sequencer

func (seqz SequencersByAddress) Len() int { return len(seqz) }

func (seqz SequencersByAddress) Less(i, j int) bool {
	return bytes.Compare(seqz[i].Address.Bytes(), seqz[j].Address.Bytes()) == -1
}

func (seqz SequencersByAddress) Swap(i, j int) {
	seqz[i], seqz[j] = seqz[j], seqz[i]
}

// safe addition/subtraction/multiplication

func safeAdd(a, b int64) (int64, bool) {
	if b > 0 && a > math.MaxInt64-b {
		return -1, true
	} else if b < 0 && a < math.MinInt64-b {
		return -1, true
	}
	return a + b, false
}

func safeSub(a, b int64) (int64, bool) {
	if b > 0 && a < math.MinInt64+b {
		return -1, true
	} else if b < 0 && a > math.MaxInt64+b {
		return -1, true
	}
	return a - b, false
}

func safeAddClip(a, b int64) int64 {
	c, overflow := safeAdd(a, b)
	if overflow {
		if b < 0 {
			return math.MinInt64
		}
		return math.MaxInt64
	}
	return c
}

func safeSubClip(a, b int64) int64 {
	c, overflow := safeSub(a, b)
	if overflow {
		if b > 0 {
			return math.MinInt64
		}
		return math.MaxInt64
	}
	return c
}

func safeMul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	absOfB := b
	if b < 0 {
		absOfB = -b
	}

	absOfA := a
	if a < 0 {
		absOfA = -a
	}

	if absOfA > math.MaxInt64/absOfB {
		return 0, true
	}

	return a * b, false
}
