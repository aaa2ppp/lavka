package api

import (
	"fmt"
	"slices"
	"time"

	"lavka/internal/lib"
)

func errCannotBeEmpty(name string) error {
	return fmt.Errorf("%s cannot be empty", name)
}

func errMustBeGreaterThanZero(name string) error {
	return fmt.Errorf("%s must be greater than 0", name)
}

func errItemMustBeGreaterThanZero(name string, i int) error {
	return fmt.Errorf("%s[%d] must be greater than 0", name, i)
}

func errMustBeOnOf(name string, enum []string) error {
	return fmt.Errorf("%s must be one of %v", name, enum)
}

func errMustBeDateTime(name string) error {
	return fmt.Errorf("%s must be date-time in RFC3339 format", name)
}

func errMustBeDayTimePeriod(name string) error {
	return fmt.Errorf("%s must be time period in HH:MM-HH:MM format", name)
}

func errItemMustBeDayTimePeriod(name string, i int) error {
	return fmt.Errorf("%s[%d] must be time period in HH:MM-HH:MM format", name, i)
}

type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers,omitempty"`
}

func (v CreateCourierRequest) Validate() error {

	if len(v.Couriers) == 0 {
		return errCannotBeEmpty("couriers")
	}

	for i := range v.Couriers {
		if err := v.Couriers[i].Validate(); err != nil {
			return fmt.Errorf("couriers[%d]: %v", i, err)
		}
	}

	return nil
}

type CreateCourierResponse struct {
	Couriers []CourierDto `json:"couriers,omitempty"`
}

type CreateCourierDto struct {
	CourierType  string   `json:"courier_type,omitempty"`
	Regions      []int    `json:"regions,omitempty"`
	WorkingHours []string `json:"working_hours,omitempty"`
}

var allowCourierTypes = []string{"FOOT", "BIKE", "AUTO"}

func (v CreateCourierDto) Validate() error {

	if !slices.Contains(allowCourierTypes, v.CourierType) {
		return errMustBeOnOf("courier_type", allowCourierTypes)
	}

	if len(v.Regions) == 0 {
		return errCannotBeEmpty("regions")
	}
	for i, region := range v.Regions {
		if region <= 0 {
			return errItemMustBeGreaterThanZero("regions", i)
		}
	}

	if len(v.WorkingHours) <= 0 {
		return errCannotBeEmpty("working_hours")
	}
	for i, period := range v.WorkingHours {
		if _, _, err := lib.ParseDayTimePeriod(period); err != nil {
			return errItemMustBeDayTimePeriod("working_hours", i)
		}
	}

	return nil
}

type CourierDto struct {
	CourierID    int64
	CourierType  string   `json:"courier_type,omitempty"`
	Regions      []int    `json:"regions,omitempty"`
	WorkingHours []string `json:"working_hours,omitempty"`
}

type GetCouriersResponse struct {
	Couriers []CourierDto
	Limit    int
	Offset   int
}

type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders,omitempty"`
}

func (v CreateOrderRequest) Validate() error {
	if len(v.Orders) == 0 {
		return errCannotBeEmpty("orders")
	}

	for i := range v.Orders {
		if err := v.Orders[i].Validate(); err != nil {
			return fmt.Errorf("orders[%d]: %v", i, err)
		}
	}

	return nil
}

type CreateOrderDto struct {
	Weight        float64  `json:"weight,omitempty"`
	Regions       int      `json:"regions,omitempty"`
	DeliveryHours []string `json:"delivery_hours,omitempty"`
	Cost          int      `json:"cost,omitempty"`
}

func (v CreateOrderDto) Validate() error {

	if v.Weight <= 0 {
		return errMustBeGreaterThanZero("weight")
	}

	if v.Regions == 0 {
		return errMustBeGreaterThanZero("regions")
	}

	if len(v.DeliveryHours) == 0 {
		return errCannotBeEmpty("delivery_hours")
	}
	for i, period := range v.DeliveryHours {
		if _, _, err := lib.ParseDayTimePeriod(period); err != nil {
			return errItemMustBeDayTimePeriod("delivery_hours", i)
		}
	}

	if v.Cost <= 0 {
		return errMustBeGreaterThanZero("cost")
	}

	return nil
}

type OrderDto struct {
	OrderID       int64    `json:"order_id,omitempty"`
	Weight        float64  `json:"weight,omitempty"`
	Regions       int      `json:"regions,omitempty"`
	DeliveryHours []string `json:"delivery_hours,omitempty"`
	Cost          int      `json:"cost,omitempty"`
	CompletedTime string   `json:"completed_time,omitempty"`
}

type CompleteOrderRequestDto struct {
	CompleteInfo []CompleteOrder `json:"complete_info,omitempty"`
}

func (v CompleteOrderRequestDto) Validate() error {
	if len(v.CompleteInfo) == 0 {
		return errCannotBeEmpty("complete_info")
	}

	for i := range v.CompleteInfo {
		if err := v.CompleteInfo[i].Validate(); err != nil {
			return fmt.Errorf("complete_info[%d]: %v", i, err)
		}
	}

	return nil
}

type CompleteOrder struct {
	CompleteTime string `json:"complete_time,omitempty"` // date-time
	CourierID    int64  `json:"courier_id,omitempty"`
	OrderID      int64  `json:"order_id,omitempty"`
}

func (v CompleteOrder) Validate() error {

	if _, err := time.Parse(time.RFC3339, v.CompleteTime); err != nil {
		return errMustBeDateTime("complete_time")
	}

	if v.CourierID <= 0 {
		return errMustBeGreaterThanZero("courier_id")
	}

	if v.OrderID <= 0 {
		return errMustBeGreaterThanZero("order_id")
	}

	return nil
}
