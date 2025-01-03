package model

import "lavka/internal/lib"

type CreateOrderDto struct {
	Weight        float64  `json:"weight,omitempty"`
	Regions       int      `json:"regions,omitempty"`
	DeliveryHours []string `json:"delivery_hours,omitempty"`
	Cost          int      `json:"cost,omitempty"`
}

func (v CreateOrderDto) Validate() error {

	if v.Weight <= 0 {
		return ErrMustBeGreaterThanZero("weight")
	}

	if v.Regions == 0 {
		return ErrMustBeGreaterThanZero("regions")
	}

	if len(v.DeliveryHours) == 0 {
		return ErrCannotBeEmpty("delivery_hours")
	}
	for i, period := range v.DeliveryHours {
		if _, _, err := lib.ParseDayTimePeriod(period); err != nil {
			return ErrItemError("delivery_hours", i, err)
		}
	}

	if v.Cost <= 0 {
		return ErrMustBeGreaterThanZero("cost")
	}

	return nil
}
