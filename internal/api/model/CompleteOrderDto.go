package model

import (
	"time"
)

type CompleteOrderDto struct {
	CompleteTime string `json:"complete_time,omitempty"` // date-time
	CourierID    int64  `json:"courier_id,omitempty"`
	OrderID      int64  `json:"order_id,omitempty"`
}

func (v CompleteOrderDto) Validate() error {

	if _, err := time.Parse(time.RFC3339, v.CompleteTime); err != nil {
		return ErrMustBeDateTime("complete_time")
	}

	if v.CourierID <= 0 {
		return ErrMustBeGreaterThanZero("courier_id")
	}

	if v.OrderID <= 0 {
		return ErrMustBeGreaterThanZero("order_id")
	}

	return nil
}
