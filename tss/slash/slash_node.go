package slash

import (
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"

	tss "github.com/mantlenetworkio/mantle/tss/common"
	"github.com/mantlenetworkio/mantle/tss/index"
)

type SlashingNode struct {
	stateBatchStore index.StateBatchStore
	slashingStore   SlashingStore
}

func NewSlashingNode(sbs index.StateBatchStore, ss SlashingStore) SlashingNode {
	return SlashingNode{
		stateBatchStore: sbs,
		slashingStore:   ss,
	}
}

func (s SlashingNode) AfterStateBatchIndexed(root [32]byte) error {
	found, stateBatch := s.stateBatchStore.GetStateBatch(root)
	if !found {
		return errors.New("can not find the state batch with root: " + hexutil.Encode(root[:]))
	}

	// update signingInfo for absent nodes
	for _, absentNode := range stateBatch.AbsentNodes {
		address, err := tss.NodeToAddress(absentNode)
		if err != nil {
			return err
		}
		s.slashingStore.SetSlashingInfo(SlashingInfo{
			Address:    address,
			ElectionId: stateBatch.ElectionId,
			BatchIndex: stateBatch.BatchIndex,
			SlashType:  tss.SlashTypeLiveness,
		})
	}

	return nil
}
