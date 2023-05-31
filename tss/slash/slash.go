package slash

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/l2geth/common/hexutil"
	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/index"
)

type Slashing struct {
	stateBatchStore index.StateBatchStore
	slashingStore   SlashingStore

	missSignedNumber int
}

func NewSlashing(sbs index.StateBatchStore, ss SlashingStore, missSignedNumber int) Slashing {
	return Slashing{
		stateBatchStore:     sbs,
		slashingStore:       ss,
		missSignedNumber: missSignedNumber,
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
		if updatedSigningInfo.MissedBlocksCounter > uint64(s.missSignedNumber) {
			s.slashingStore.SetSlashingInfo(SlashingInfo{
				Address:    address,
				ElectionId: stateBatch.ElectionId,
				BatchIndex: stateBatch.BatchIndex,
				SlashType:  tss.SlashTypeLiveness,
			})
			updatedSigningInfo.MissedBlocksCounter = 0
			s.slashingStore.SetSigningInfo(updatedSigningInfo)
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
		if missed {
			signingInfo.MissedBlocksCounter++;
		}
	}
	s.slashingStore.SetSigningInfo(signingInfo)
	return signingInfo
}

func (s Slashing) InitializeSigningInfo(batchIndex uint64, address common.Address, missed bool) SigningInfo {
	signingInfo := SigningInfo{
		Address:             address,
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
