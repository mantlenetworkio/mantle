package store

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/mantle/tss/slash"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (s *Storage) SetSigningInfo(signingInfo slash.SigningInfo) {
	bz, err := json.Marshal(signingInfo)
	if err != nil {
		panic(err)
	}
	if err := s.db.Put(getSigningInfoKey(signingInfo.Address), bz, nil); err != nil {
		panic(err)
	}
}

func (s *Storage) GetSigningInfo(address common.Address) (bool, slash.SigningInfo) {
	bz, err := s.db.Get(getSigningInfoKey(address), nil)
	if err != nil {

		return handleError2(slash.SigningInfo{}, err)
	}
	var signingInfo slash.SigningInfo
	if err = json.Unmarshal(bz, &signingInfo); err != nil {
		panic(err)
	}
	return true, signingInfo
}

func (s *Storage) GetNodeMissedBatchBitArray(address common.Address, index uint64) bool {
	bz, err := s.db.Get(getNodeMissedBatchBitArrayKey(address, index), nil)
	if err != nil {

		if err == leveldb.ErrNotFound {
			return false // lazy: treat empty key as not missed
		}
		panic(err)
	}
	return bz[0] == 1
}

func (s *Storage) SetNodeMissedBatchBitArray(address common.Address, index uint64, missed bool) {
	var missedBz byte
	if missed {
		missedBz = 1
	}
	if err := s.db.Put(getNodeMissedBatchBitArrayKey(address, index), []byte{missedBz}, nil); err != nil {
		panic(err)
	}
}

func (s *Storage) ClearNodeMissedBatchBitArray(address common.Address) {
	iterator := s.db.NewIterator(util.BytesPrefix(getNodeMissedBatchBitArrayAddressPrefixKey(address)), nil)
	defer iterator.Release()

	for iterator.Next() {
		if err := s.db.Delete(iterator.Key(), nil); err != nil {
			panic(err)
		}
	}
}

func (s *Storage) SetSlashingInfo(slashingInfo slash.SlashingInfo) {
	bz, err := json.Marshal(slashingInfo)
	if err != nil {
		panic(err)
	}
	if err = s.db.Put(getSlashingInfoKey(slashingInfo.Address, slashingInfo.BatchIndex), bz, nil); err != nil {
		panic(err)
	}
}

func (s *Storage) GetSlashingInfo(address common.Address, batchIndex uint64) (bool, slash.SlashingInfo) {
	bz, err := s.db.Get(getSlashingInfoKey(address, batchIndex), nil)
	if err != nil {
		return handleError2(slash.SlashingInfo{}, err)
	}
	var slashingInfo slash.SlashingInfo
	if err = json.Unmarshal(bz, &slashingInfo); err != nil {
		panic(err)
	}
	return true, slashingInfo
}

func (s *Storage) IsInSlashing(address common.Address) bool {
	iterator := s.db.NewIterator(util.BytesPrefix(getSlashingInfoAddressKey(address)), nil)
	defer iterator.Release()
	for iterator.Next() {
		buf := iterator.Value()
		if len(buf) > 0 {
			return true
		}
	}
	return false
}

func (s *Storage) ListSlashingInfo() (slashingInfos []slash.SlashingInfo) {
	iterator := s.db.NewIterator(util.BytesPrefix(SlashingInfoKeyPrefix), nil)
	defer iterator.Release()
	for iterator.Next() {
		buf := iterator.Value()
		if len(buf) == 0 {
			continue
		}
		var slashingInfo slash.SlashingInfo
		if err := json.Unmarshal(iterator.Value(), &slashingInfo); err != nil {
			panic(err)
		}
		slashingInfos = append(slashingInfos, slashingInfo)
	}
	return
}

func (s *Storage) RemoveSlashingInfo(address common.Address, batchIndex uint64) {
	if err := s.db.Delete(getSlashingInfoKey(address, batchIndex), nil); err != nil {
		panic(err)
	}
}
