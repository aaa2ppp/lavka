package courierController

import (
	"lavka/internal/model"
)

type createCourierResponse struct {
	Couriers []model.CourierDto `json:"couriers,omitempty"`
}
