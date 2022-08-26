package l1chain

import "github.com/bitdao-io/bitnetwork/tss/manager/types"

type QueryService struct {
}

func (q QueryService) QueryActiveInfo() types.TssCommitteeInfo {
	return types.TssCommitteeInfo{}
}

func (q QueryService) QueryInactiveInfo() types.TssCommitteeInfo {
	return types.TssCommitteeInfo{}
}
