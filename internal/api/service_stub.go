package api

import (
	"context"
	. "lavka/internal/model"
)

type ServiceStub struct{}

// ComleteOrder implements Service.
func (f ServiceStub) ComleteOrder(ctx context.Context, orders []CompleteOrderDto) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

// CreateCourier implements Service.
func (f ServiceStub) CreateCourier(ctx context.Context, couriers []CreateCourierDto) ([]CourierDto, error) {
	zero := []CourierDto{}
	return zero, nil
}

// CreateOrder implements Service.
func (f ServiceStub) CreateOrder(ctx context.Context, orders []CreateOrderDto) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

// GetCourierById implements Service.
func (f ServiceStub) GetCourierById(ctx context.Context, courierID uint64) (CourierDto, error) {
	zero := CourierDto{}
	return zero, nil
}

// GetCouriers implements Service.
func (f ServiceStub) GetCouriers(ctx context.Context, limit int, offset int) ([]CourierDto, error) {
	zero := []CourierDto{}
	return zero, nil
}

// GetOrder implements Service.
func (f ServiceStub) GetOrder(ctx context.Context, orderID uint64) (OrderDto, error) {
	zero := OrderDto{}
	return zero, nil
}

// GetOrders implements Service.
func (f ServiceStub) GetOrders(ctx context.Context, limit int, offset int) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

var _ Service = ServiceStub{}
