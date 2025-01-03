package model

type CourierDto struct {
	CourierID    int64    `json:"courier_id,omitempty"`
	CourierType  string   `json:"courier_type,omitempty"`
	Regions      []int    `json:"regions,omitempty"`
	WorkingHours []string `json:"working_hours,omitempty"`
}
