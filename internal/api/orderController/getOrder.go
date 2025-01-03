package orderController

import (
	"net/http"

	"lavka/internal/api/helper"
)

func (c controller) getOrder(w http.ResponseWriter, r *http.Request) {
	x := helper.New("getOrder", w, r)

	orderID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	order, err := c.GetOrder(orderID)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := order
	x.WriteResponse(resp)
}
