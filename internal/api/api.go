package api

import (
	"net/http"

	"lavka/internal/api/courierController"
	"lavka/internal/api/orderController"
)

type Service interface {
	courierController.CourierService
	orderController.OrderService
}

func New(service Service, rps int) http.Handler {
	mux := http.NewServeMux()
	Setup(mux, service, rps)
	return mux
}

func Setup(mux *http.ServeMux, service Service, rps int) {
	courierController.Setup(mux, service, rps)
	orderController.Setup(mux, service, rps)
}
