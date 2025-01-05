package orderController

import (
	"errors"
	"fmt"

	"lavka/internal/model"
)

type completeOrderRequest struct {
	CompleteInfo []model.CompleteOrderDto `json:"complete_info,omitempty"`
}

func (v completeOrderRequest) Validate() error {
	if len(v.CompleteInfo) == 0 {
		return errors.New("complete_info cannot be empty")
	}

	for i := range v.CompleteInfo {
		if err := v.CompleteInfo[i].Validate(); err != nil {
			return fmt.Errorf("complete_info[%d]: %w", i, err)
		}
	}

	return nil
}
