package orderController

import (
	"net/http"

	"lavka/internal/api/helper"
)

func (c controller) getOrders(w http.ResponseWriter, r *http.Request) {
	x := helper.New("getOrders", w, r)

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.GetOrders(limit, offset)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
