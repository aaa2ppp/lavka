package model

import (
	"errors"
	"lavka/internal/model/courier"
	"lavka/internal/model/daytime"
)

type CreateCourierDto struct {
	CourierType  courier.Type     `json:"courier_type,omitempty"`
	Regions      []int            `json:"regions,omitempty"`
	WorkingHours []daytime.Period `json:"working_hours,omitempty"`
}

func (v CreateCourierDto) Validate() error {

	if len(v.Regions) == 0 {
		return errors.New("regions cannot be empty")
	}

	for _, region := range v.Regions {
		if region <= 0 {
			return errors.New("region must be > 0")
		}
	}

	if len(v.WorkingHours) == 0 {
		return errors.New("working_hours cannot be empty")
	}

	return nil
}
