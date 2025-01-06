package orderController

import (
	"context"
	"net/http"

	"lavka/internal/middleware"
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

var limitRPS = middleware.LimitRPS

func Setup(mux *http.ServeMux, service OrderService, rps int) {
	c := controller{service}
	mux.Handle("POST /orders", limitRPS(rps, http.HandlerFunc(c.createOrder)))
	mux.Handle("GET  /orders/{id}", limitRPS(rps, http.HandlerFunc(c.getOrder)))
	mux.Handle("GET  /orders", limitRPS(rps, http.HandlerFunc(c.getOrders)))
	mux.Handle("POST /orders/complete", limitRPS(rps, http.HandlerFunc(c.completeOrder)))
	mux.Handle("POST /orders/assign", limitRPS(rps, http.HandlerFunc(c.ordersAssign)))
}
