package courierController

import (
	"errors"
	"fmt"
	
	"lavka/internal/model"
)

type createCourierRequest struct {
	Couriers []model.CreateCourierDto `json:"couriers,omitempty"`
}

func (v createCourierRequest) Validate() error {

	if len(v.Couriers) == 0 {
		return errors.New("couriers cannot be empty")
	}

	for i := range v.Couriers {
		if err := v.Couriers[i].Validate(); err != nil {
			return fmt.Errorf("couriers[%d]: %w", i, err)
		}
	}

	return nil
}
