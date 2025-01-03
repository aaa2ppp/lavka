package orderController

import (
	"fmt"

	. "lavka/internal/api/model"
)

type completeOrderRequest struct {
	CompleteInfo []CompleteOrderDto `json:"complete_info,omitempty"`
}

func (v completeOrderRequest) Validate() error {
	if len(v.CompleteInfo) == 0 {
		return ErrCannotBeEmpty("complete_info")
	}

	for i := range v.CompleteInfo {
		if err := v.CompleteInfo[i].Validate(); err != nil {
			return fmt.Errorf("complete_info[%d]: %v", i, err)
		}
	}

	return nil
}
