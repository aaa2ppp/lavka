package courierController

import (
	"net/http"

	"lavka/internal/api/helper"
)

func (c controller) getCouriers(w http.ResponseWriter, r *http.Request) {
	x := helper.New("getCouriers", w, r)

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := c.GetCouriers(limit, offset)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := getCouriersResponse{
		Couriers: couriers,
		Limit:    limit,
		Offset:   offset,
	}

	x.WriteResponse(resp)
}
