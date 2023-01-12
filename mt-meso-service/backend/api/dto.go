package api

import (
	"encoding/json"
	"time"
)

type TransactionBatchIndex struct {
	Date time.Time `json:"date" copier:"Timestamp"`
}

func (t TransactionBatchIndex) MarshalJSON() ([]byte, error) {
	type Alias TransactionBatchIndex
	return json.Marshal(&struct {
		Alias
		Date string `json:"date"`
	}{
		Alias: (Alias)(t),
		Date:  t.Date.Format("2006-01-02 15:04:05"),
	})
}
