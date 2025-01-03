package model

type OrderDto struct {
	OrderID       int64    `json:"order_id,omitempty"`
	Weight        float64  `json:"weight,omitempty"`
	Regions       int      `json:"regions,omitempty"`
	DeliveryHours []string `json:"delivery_hours,omitempty"`
	Cost          int      `json:"cost,omitempty"`
	CompletedTime string   `json:"completed_time,omitempty"`
}

