package model

import (
	"lavka/internal/lib"
	"slices"
)

type CreateCourierDto struct {
	CourierType  string   `json:"courier_type,omitempty"`
	Regions      []int    `json:"regions,omitempty"`
	WorkingHours []string `json:"working_hours,omitempty"`
}

var allowCourierTypes = []string{"FOOT", "BIKE", "AUTO"}

func (v CreateCourierDto) Validate() error {

	if !slices.Contains(allowCourierTypes, v.CourierType) {
		return ErrMustBeOnOf("courier_type", allowCourierTypes)
	}

	if len(v.Regions) == 0 {
		return ErrCannotBeEmpty("regions")
	}
	for i, region := range v.Regions {
		if region <= 0 {
			return ErrItemMustBeGreaterThanZero("regions", i)
		}
	}

	if len(v.WorkingHours) <= 0 {
		return ErrCannotBeEmpty("working_hours")
	}
	for i, period := range v.WorkingHours {
		if _, _, err := lib.ParseDayTimePeriod(period); err != nil {
			return ErrItemMustBeDayTimePeriod("working_hours", i)
		}
	}

	return nil
}
