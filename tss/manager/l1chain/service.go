package l1chain

import "github.com/bitdao-io/bitnetwork/tss/manager/types"

type QueryService struct {
}

func (q QueryService) QueryInfo() types.TssInfos {
	return types.TssInfos{}
}
