package orderController

import (
	"net/http"
)

func (c controller) getOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "getOrder")

	orderID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	order, err := c.GetOrder(x.Ctx(), orderID)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := order
	x.WriteResponse(resp)
}
