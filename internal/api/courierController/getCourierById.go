package courierController

import (
	"lavka/internal/api/helper"
	"net/http"
)

func (c controller) getCourierById(w http.ResponseWriter, r *http.Request) {
	x := helper.New("getCourierById", w, r)

	courierID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	courier, err := c.GetCourierById(courierID)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(courier)
}
