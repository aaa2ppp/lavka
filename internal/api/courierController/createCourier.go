package courierController

import (
	"net/http"
)

func (c controller) createCourier(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "createCourier")

	var req createCourierRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := c.CreateCourier(x.Ctx(), req.Couriers)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(createCourierResponse{Couriers: couriers})
}
