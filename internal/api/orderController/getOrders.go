package orderController

import (
	"net/http"
)

func (c controller) getOrders(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "getOrders")

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.GetOrders(x.Ctx(), limit, offset)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
