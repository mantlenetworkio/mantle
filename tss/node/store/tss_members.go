package store

import (
	"encoding/json"

	"github.com/mantlenetworkio/mantle/tss/node/types"
)

func (s *Storage) SetInactiveMembers(members types.TssMembers) error {
	bz, err := json.Marshal(members)
	if err != nil {
		return err
	}
	if err := s.db.Put(getInactiveMemberKey(), bz, nil); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetInactiveMembers() (bool, types.TssMembers) {
	bz, err := s.db.Get(getInactiveMemberKey(), nil)
	if err != nil {
		return handleError2(types.TssMembers{}, err)
	}
	var tssMembers types.TssMembers
	if err = json.Unmarshal(bz, &tssMembers); err != nil {
		panic(err)
	}
	return true, tssMembers
}

func (s *Storage) SetActiveMembers(members types.TssMembers) error {
	bz, err := json.Marshal(members)
	if err != nil {
		return err
	}
	if err := s.db.Put(getActiveMemberKey(), bz, nil); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetActiveMembers() (bool, types.TssMembers) {
	bz, err := s.db.Get(getActiveMemberKey(), nil)
	if err != nil {
		return handleError2(types.TssMembers{}, err)
	}
	var tssMembers types.TssMembers
	if err = json.Unmarshal(bz, &tssMembers); err != nil {
		panic(err)
	}
	return true, tssMembers
}
