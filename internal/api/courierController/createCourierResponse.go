package courierController

import (
	m "lavka/internal/api/model"
)

type createCourierResponse struct {
	Couriers []m.CourierDto `json:"couriers,omitempty"`
}
