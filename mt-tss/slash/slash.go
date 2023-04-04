package slash

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tss "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/mantlenetworkio/mantle/mt-tss/index"
)

type Slashing struct {
	outputStore   index.OutputStore
	slashingStore SlashingStore

	signedBatchesWindow int
	minSignedInWindow   int
}

func NewSlashing(ops index.OutputStore, ss SlashingStore, signedBatchesWindow, minSignedInWindow int) Slashing {
	return Slashing{
		outputStore:         ops,
		slashingStore:       ss,
		signedBatchesWindow: signedBatchesWindow,
		minSignedInWindow:   minSignedInWindow,
	}
}

func (s Slashing) AfterStateBatchIndexed(root [32]byte) error {
	found, output := s.outputStore.GetOutput(root)
	if !found {
		return errors.New("can not find the state batch with root: " + hexutil.Encode(root[:]))
	}

	// check whether it is a new round election
	var electionAdvanced bool
	if output.ElectionId > 1 {
		found, previousBatchRoot := s.outputStore.GetIndexOutput(output.OutputIndex - 1)
		if found {
			_, previousStateBatch := s.outputStore.GetOutput(previousBatchRoot)
			electionAdvanced = output.ElectionId != previousStateBatch.ElectionId
		}
	}

	maxMissed := s.signedBatchesWindow - s.minSignedInWindow
	// update signingInfo for working nodes
	for _, workingNode := range output.WorkingNodes {
		address, err := tss.NodeToAddress(workingNode)
		if err != nil {
			return err
		}
		s.UpdateSigningInfo(output.OutputIndex, address, electionAdvanced, false)
		if err != nil {
			return err
		}
	}

	// update signingInfo for absent nodes
	for _, absentNode := range output.AbsentNodes {
		address, err := tss.NodeToAddress(absentNode)
		if err != nil {
			return err
		}
		updatedSigningInfo := s.UpdateSigningInfo(output.OutputIndex, address, electionAdvanced, true)
		if err != nil {
			return err
		}
		if updatedSigningInfo.MissedBlocksCounter > uint64(maxMissed) {
			s.slashingStore.SetSlashingInfo(SlashingInfo{
				Address:    address,
				ElectionId: output.ElectionId,
				BatchIndex: output.OutputIndex,
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
