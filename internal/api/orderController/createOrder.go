package orderController

import (
	"net/http"
)

func (c controller) createOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "createOrder")

	var req createOrderRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.CreateOrder(x.Ctx(), req.Orders)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
