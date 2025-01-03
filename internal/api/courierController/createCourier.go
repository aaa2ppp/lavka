package courierController

import (
	"lavka/internal/api/helper"
	"net/http"
)

func (c controller) createCourier(w http.ResponseWriter, r *http.Request) {
	x := helper.New("createCourier", w, r)

	var req createCourierRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := c.CreateCourier(req.Couriers)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(createCourierResponse{Couriers: couriers})
}
