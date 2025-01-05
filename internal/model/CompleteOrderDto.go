package model

import (
	"errors"
	"time"
)

type CompleteOrderDto struct {
	CompleteTime time.Time `json:"complete_time,omitempty"` // date-time
	CourierID    int64     `json:"courier_id,omitempty"`
	OrderID      int64     `json:"order_id,omitempty"`
}

func (v CompleteOrderDto) Validate() error {

	if v.CompleteTime.IsZero() {
		return errors.New("complete_time cannot be empty")
	}

	if v.CourierID <= 0 {
		return errors.New("courier_id must be > 0")
	}

	if v.OrderID <= 0 {
		return errors.New("order_id must be > 0")
	}

	return nil
}
