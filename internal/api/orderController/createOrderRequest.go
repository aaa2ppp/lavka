package orderController

import (
	"errors"
	"fmt"

	"lavka/internal/model"
)

type createOrderRequest struct {
	Orders []model.CreateOrderDto `json:"orders,omitempty"`
}

func (v createOrderRequest) Validate() error {
	if len(v.Orders) == 0 {
		return errors.New("orders cannot be empty")
	}

	for i := range v.Orders {
		if err := v.Orders[i].Validate(); err != nil {
			return fmt.Errorf("orders[%d]: %w", i, err)
		}
	}

	return nil
}
