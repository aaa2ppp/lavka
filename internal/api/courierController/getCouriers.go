package courierController

import (
	"net/http"
)

func (c controller) getCouriers(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "getCouriers")

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := c.GetCouriers(x.Ctx(), limit, offset)
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
