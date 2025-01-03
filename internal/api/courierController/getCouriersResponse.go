package courierController

import (
	m "lavka/internal/api/model"
)

type getCouriersResponse struct {
	Couriers []m.CourierDto
	Limit    int
	Offset   int
}
