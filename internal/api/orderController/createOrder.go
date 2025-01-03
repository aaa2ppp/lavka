package orderController

import (
	"net/http"

	"lavka/internal/api/helper"
)

func (c controller) createOrder(w http.ResponseWriter, r *http.Request) {
	x := helper.New("createOrder", w, r)

	var req createOrderRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.CreateOrder(req.Orders)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
