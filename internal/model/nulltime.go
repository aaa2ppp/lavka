package model

import (
	"encoding/json"
	"time"
)

type NullTime struct {
	time.Time
}

// MarshalJSON implements json.Marshaler.
func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

var _ json.Marshaler = NullTime{}
