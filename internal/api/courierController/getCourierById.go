package courierController

import (
	"net/http"
)

func (c controller) getCourierById(w http.ResponseWriter, r *http.Request) {
	x := newHelper(w, r, "getCourierById")

	courierID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	courier, err := c.GetCourierById(x.Ctx(), courierID)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(courier)
}
