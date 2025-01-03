package orderController

import (
	"net/http"

	. "lavka/internal/api/model"
)

type Service interface {
	CreateOrder(orders []CreateOrderDto) ([]OrderDto, error)
	GetOrder(orderID uint64) (OrderDto, error)
	GetOrders(limit, offset int) ([]OrderDto, error)
	ComleteOrder([]CompleteOrderDto) ([]OrderDto, error)
}

type controller struct {
	Service
}

func Setup(mux *http.ServeMux, service Service) {
	c := controller{service}
	mux.Handle("POST /orders", http.HandlerFunc(c.createOrder))
	mux.Handle("GET  /orders/{id}", http.HandlerFunc(c.getOrder))
	mux.Handle("GET  /orders", http.HandlerFunc(c.getOrders))
	mux.Handle("POST /orders/complete", http.HandlerFunc(c.completeOrder))
	mux.Handle("POST /orders/assign", http.HandlerFunc(c.ordersAssign))
}
