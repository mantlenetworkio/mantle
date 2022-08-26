package slash

import (
	"errors"

	"github.com/bitdao-io/bitnetwork/l2geth/common"
	"github.com/bitdao-io/bitnetwork/l2geth/common/hexutil"
	tss "github.com/bitdao-io/bitnetwork/tss/common"
	"github.com/bitdao-io/bitnetwork/tss/index"
)

const (
	signedBatchesWindow = 100
	minSignedInWindow   = 50
)

type Slashing struct {
	stateBatchStore index.StateBatchStore
	slashingStore   SlashingStore
}

func NewSlashing(sbs index.StateBatchStore, ss SlashingStore) Slashing {
	return Slashing{
		stateBatchStore: sbs,
		slashingStore:   ss,
	}
}

func (s Slashing) AfterStateBatchIndexed(root [32]byte) error {
	found, stateBatch, err := s.stateBatchStore.GetStateBatch(root)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("can not find the state batch with root: " + hexutil.Encode(root[:]))
	}

	// check whether it is a new round election
	var electionAdvanced bool
	if stateBatch.ElectionId > 1 {
		found, previousBatchRoot, err := s.stateBatchStore.GetIndexStateBatch(stateBatch.BatchIndex - 1)
		if err != nil {
			return err
		}
		if found {
			_, previousStateBatch, err := s.stateBatchStore.GetStateBatch(previousBatchRoot)
			if err != nil {
				return err
			}
			electionAdvanced = stateBatch.ElectionId != previousStateBatch.ElectionId
		}
	}

	maxMissed := signedBatchesWindow - minSignedInWindow
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
				SlashType:  1,
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
		idx := signingInfo.IndexOffset % signedBatchesWindow

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
	s.slashingStore.ClearNodeMissedBatchBitArray(address)
	return signingInfo
}
