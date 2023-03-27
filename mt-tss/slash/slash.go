package slash

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/index"
)

type Slashing struct {
	stateBatchStore index.StateBatchStore
	slashingStore   SlashingStore

	signedBatchesWindow int
	minSignedInWindow   int
}

func NewSlashing(sbs index.StateBatchStore, ss SlashingStore, signedBatchesWindow, minSignedInWindow int) Slashing {
	return Slashing{
		stateBatchStore:     sbs,
		slashingStore:       ss,
		signedBatchesWindow: signedBatchesWindow,
		minSignedInWindow:   minSignedInWindow,
	}
}

func (s Slashing) AfterStateBatchIndexed(root [32]byte) error {
	found, stateBatch := s.stateBatchStore.GetStateBatch(root)
	if !found {
		return errors.New("can not find the state batch with root: " + hexutil.Encode(root[:]))
	}

	// check whether it is a new round election
	var electionAdvanced bool
	if stateBatch.ElectionId > 1 {
		found, previousBatchRoot := s.stateBatchStore.GetIndexStateBatch(stateBatch.BatchIndex - 1)
		if found {
			_, previousStateBatch := s.stateBatchStore.GetStateBatch(previousBatchRoot)
			electionAdvanced = stateBatch.ElectionId != previousStateBatch.ElectionId
		}
	}

	maxMissed := s.signedBatchesWindow - s.minSignedInWindow
	// update signingInfo for working nodes
	for _, workingNode := range stateBatch.WorkingNodes {
		address, err := tss.NodeToAddress(workingNode)
		if err != nil {
			return err
		}
		s.UpdateSigningInfo(stateBatch.BatchIndex, address, electionAdvanced, false)
		if err != nil {
			return err
		}
	}

	// update signingInfo for absent nodes
	for _, absentNode := range stateBatch.AbsentNodes {
		address, err := tss.NodeToAddress(absentNode)
		if err != nil {
			return err
		}
		updatedSigningInfo := s.UpdateSigningInfo(stateBatch.BatchIndex, address, electionAdvanced, true)
		if err != nil {
			return err
		}
		if updatedSigningInfo.MissedBlocksCounter > uint64(maxMissed) {
			s.slashingStore.SetSlashingInfo(SlashingInfo{
				Address:    address,
				ElectionId: stateBatch.ElectionId,
				BatchIndex: stateBatch.BatchIndex,
				SlashType:  tss.SlashTypeLiveness,
			})
		}
	}

	return nil
}

func (s Slashing) UpdateSigningInfo(batchIndex uint64, address common.Address, electionAdvanced, missed bool) SigningInfo {
	if electionAdvanced {
		return s.InitializeSigningInfo(batchIndex, address, missed)
	}

	found, signingInfo := s.slashingStore.GetSigningInfo(address)
	if !found {
		signingInfo = s.InitializeSigningInfo(batchIndex, address, missed)
	} else {
		signingInfo.IndexOffset++

		idx := signingInfo.IndexOffset % uint64(s.signedBatchesWindow)

		previous := s.slashingStore.GetNodeMissedBatchBitArray(address, idx)
		switch {
		case !previous && missed:
			s.slashingStore.SetNodeMissedBatchBitArray(address, idx, true)
			signingInfo.MissedBlocksCounter++
		case previous && !missed:
			s.slashingStore.SetNodeMissedBatchBitArray(address, idx, false)
			signingInfo.MissedBlocksCounter--
		default:
			// array value at this index has not changed, no need to update counter
		}
	}
	s.slashingStore.SetSigningInfo(signingInfo)
	return signingInfo
}

func (s Slashing) InitializeSigningInfo(batchIndex uint64, address common.Address, missed bool) SigningInfo {
	signingInfo := SigningInfo{
		Address:             address,
		StartBatchIndex:     batchIndex,
		IndexOffset:         0,
		MissedBlocksCounter: 0,
	}
	if missed {
		signingInfo.MissedBlocksCounter++
	}
	s.slashingStore.SetSigningInfo(signingInfo)

	// clear historic data
	s.slashingStore.ClearNodeMissedBatchBitArray(address)
	// init the first one
	s.slashingStore.SetNodeMissedBatchBitArray(address, 0, missed)
	return signingInfo
}
