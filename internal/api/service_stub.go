package api

import (
	. "lavka/internal/api/model"
)

type ServiceStub struct{}

// ComleteOrder implements Service.
func (f ServiceStub) ComleteOrder([]CompleteOrderDto) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

// CreateCourier implements Service.
func (f ServiceStub) CreateCourier(couriers []CreateCourierDto) ([]CourierDto, error) {
	zero := []CourierDto{}
	return zero, nil
}

// CreateOrder implements Service.
func (f ServiceStub) CreateOrder(orders []CreateOrderDto) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

// GetCourierById implements Service.
func (f ServiceStub) GetCourierById(courierID uint64) (CourierDto, error) {
	zero := CourierDto{}
	return zero, nil
}

// GetCouriers implements Service.
func (f ServiceStub) GetCouriers(limit int, offset int) ([]CourierDto, error) {
	zero := []CourierDto{}
	return zero, nil
}

// GetOrder implements Service.
func (f ServiceStub) GetOrder(orderID uint64) (OrderDto, error) {
	zero := OrderDto{}
	return zero, nil
}

// GetOrders implements Service.
func (f ServiceStub) GetOrders(limit int, offset int) ([]OrderDto, error) {
	zero := []OrderDto{}
	return zero, nil
}

var _ Service = ServiceStub{}
