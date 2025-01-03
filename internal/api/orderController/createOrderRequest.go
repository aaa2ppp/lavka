package orderController

import (
	. "lavka/internal/api/model"
)

type createOrderRequest struct {
	Orders []CreateOrderDto `json:"orders,omitempty"`
}

func (v createOrderRequest) Validate() error {
	if len(v.Orders) == 0 {
		return ErrCannotBeEmpty("orders")
	}

	for i := range v.Orders {
		if err := v.Orders[i].Validate(); err != nil {
			return ErrItemError("orders", i, err)
		}
	}

	return nil
}
