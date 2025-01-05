package orderController

import (
	"net/http"
)

func (c controller) completeOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "completeOrder")

	var req completeOrderRequest
	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.ComleteOrder(x.Ctx(), req.CompleteInfo)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
