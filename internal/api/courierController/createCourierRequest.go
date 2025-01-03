package courierController

import (
	m "lavka/internal/api/model"
)

type createCourierRequest struct {
	Couriers []m.CreateCourierDto `json:"couriers,omitempty"`
}

func (v createCourierRequest) Validate() error {

	if len(v.Couriers) == 0 {
		return m.ErrCannotBeEmpty("couriers")
	}

	for i := range v.Couriers {
		if err := v.Couriers[i].Validate(); err != nil {
			return m.ErrItemError("couriers", i, err)
		}
	}

	return nil
}
