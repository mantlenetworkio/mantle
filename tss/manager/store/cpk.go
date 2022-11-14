package store

import (
	"encoding/json"
	"github.com/mantlenetworkio/mantle/tss/manager/types"
)

func (s *Storage) Insert(cpkData types.CpkData) error {
	bz, err := json.Marshal(cpkData)
	if err != nil {
		return err
	}
	return s.db.Put(getCPKDataKey(cpkData.ElectionId), bz, nil)
}

func (s *Storage) GetByElectionId(electionId uint64) (types.CpkData, error) {
	bz, err := s.db.Get(getCPKDataKey(electionId), nil)
	if err != nil {
		return handleError(types.CpkData{}, err)
	}
	var cpkData types.CpkData
	if err = json.Unmarshal(bz, &cpkData); err != nil {
		return types.CpkData{}, err
	}
	return cpkData, nil
}
