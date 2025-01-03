package api

import (
	"net/http"

	"lavka/internal/api/courierController"
	"lavka/internal/api/orderController"
)

type Service interface {
	courierController.Service
	orderController.Service
}

func New(service Service) http.Handler {
	mux := http.NewServeMux()
	Setup(mux, service)
	return mux
}

func Setup(mux *http.ServeMux, service Service) {
	courierController.Setup(mux, service)
	orderController.Setup(mux, service)
}
