package model

import (
	"lavka/internal/model/daytime"
)

type OrderDto struct {
	OrderID       int64            `json:"order_id,omitempty"`
	Weight        float64          `json:"weight,omitempty"`
	Regions       int              `json:"regions,omitempty"`
	DeliveryHours []daytime.Period `json:"delivery_hours,omitempty"`
	Cost          int              `json:"cost,omitempty"`
	CompletedTime NullTime         `json:"completed_time,omitempty"`
}
