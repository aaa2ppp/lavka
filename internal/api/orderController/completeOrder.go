package orderController

import (
	"net/http"

	"lavka/internal/api/helper"
)

func (c controller) completeOrder(w http.ResponseWriter, r *http.Request) {
	x := helper.New("completeOrder", w, r)

	var req completeOrderRequest
	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := c.ComleteOrder(req.CompleteInfo)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}
