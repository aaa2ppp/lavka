package orderController

import (
	"net/http"
	"lavka/internal/api/helper"
)

func newHelper(w http.ResponseWriter, r *http.Request, op string) *helper.Helper {
	return helper.New(w, r, "orderController", op)
}
