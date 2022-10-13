package store

import "encoding/binary"

func (s *Storage) UpdateHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getScannedHeightKey(), heightBz, nil)
}

func (s *Storage) GetScannedHeight() (uint64, error) {
	bz, err := s.db.Get(getScannedHeightKey(), nil)
	if err != nil {
		return handleError(uint64(0), err)
	}
	return binary.BigEndian.Uint64(bz), nil
}
