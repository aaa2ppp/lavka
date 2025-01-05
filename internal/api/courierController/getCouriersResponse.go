package courierController

import (
	"lavka/internal/model"
)

type getCouriersResponse struct {
	Couriers []model.CourierDto
	Limit    int
	Offset   int
}
