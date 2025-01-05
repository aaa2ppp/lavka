package orderController

import (
	"context"
	"net/http"

	. "lavka/internal/model"
)

type OrderService interface {
	CreateOrder(context.Context, []CreateOrderDto) ([]OrderDto, error)
	GetOrder(_ context.Context, orderID uint64) (OrderDto, error)
	GetOrders(_ context.Context, limit, offset int) ([]OrderDto, error)
	ComleteOrder(context.Context, []CompleteOrderDto) ([]OrderDto, error)
}

type controller struct {
	OrderService
}

func Setup(mux *http.ServeMux, service OrderService) {
	c := controller{service}
	mux.Handle("POST /orders", http.HandlerFunc(c.createOrder))
	mux.Handle("GET  /orders/{id}", http.HandlerFunc(c.getOrder))
	mux.Handle("GET  /orders", http.HandlerFunc(c.getOrders))
	mux.Handle("POST /orders/complete", http.HandlerFunc(c.completeOrder))
	mux.Handle("POST /orders/assign", http.HandlerFunc(c.ordersAssign))
}
