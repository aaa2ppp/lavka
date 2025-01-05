package model

import (
	"errors"

	"lavka/internal/model/daytime"
)

type CreateOrderDto struct {
	Weight        float64          `json:"weight,omitempty"`
	Regions       int              `json:"regions,omitempty"`
	DeliveryHours []daytime.Period `json:"delivery_hours,omitempty"`
	Cost          int              `json:"cost,omitempty"`
}

func (v CreateOrderDto) Validate() error {

	if v.Weight <= 0 {
		return errors.New("weight must be > 0")
	}

	if v.Regions == 0 {
		return errors.New("regions must be > 0")
	}

	if len(v.DeliveryHours) == 0 {
		return errors.New("delivery_hours cannot be empty")
	}

	if v.Cost <= 0 {
		return errors.New("cost must be > 0")
	}

	return nil
}
