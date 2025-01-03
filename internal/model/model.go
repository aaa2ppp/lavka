package model

import (
	"lavka/internal/model/courier"
	"lavka/internal/model/daytime"
)

type CourierDto struct {
	CourierID    int64            `json:"courier_id,omitempty"`
	CourierType  courier.Type     `json:"courier_type,omitempty"`
	Regions      []int            `json:"regions,omitempty"`
	WorkingHours []daytime.Period `json:"working_hours,omitempty"`
}
