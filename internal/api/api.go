package api

import (
	"net/http"
)

type Service interface {
	CreateCourier(couriers []CreateCourierDto) ([]CourierDto, error)
	GetCourierById(courierID uint64) (CourierDto, error)
	GetCouriers(limit, offset int) ([]CourierDto, error)

	CreateOrder(orders []CreateOrderDto) ([]OrderDto, error)
	GetOrder(orderID uint64) (OrderDto, error)
	GetOrders(limit, offset int) ([]OrderDto, error)
	ComleteOrder([]CompleteOrder) ([]OrderDto, error)
}

func New(service Service) *http.ServeMux {
	router := http.NewServeMux()
	h := handler{s: service}

	// courier-controller
	router.Handle("POST /couriers", http.HandlerFunc(h.createCourier))
	router.Handle("GET  /couriers/{id}", http.HandlerFunc(h.getCourierById))
	router.Handle("GET  /couriers", http.HandlerFunc(h.getCouriers))
	router.Handle("GET  /couriers/meta-info/{id}", http.HandlerFunc(h.getCourierMetaInfo))
	router.Handle("GET  /couriers/assignments", http.HandlerFunc(h.couriersAssignments))

	// order-controller
	router.Handle("POST /orders", http.HandlerFunc(h.createOrder))
	router.Handle("GET  /orders/{id}", http.HandlerFunc(h.getOrder))
	router.Handle("GET  /orders", http.HandlerFunc(h.getOrders))
	router.Handle("POST /orders/complete", http.HandlerFunc(h.completeOrder))
	router.Handle("POST /orders/assign", http.HandlerFunc(h.ordersAssign))

	return router
}

type handler struct {
	s Service
}

func (h handler) createCourier(w http.ResponseWriter, r *http.Request) {
	x := newHelper("createCourier", w, r)

	var req CreateCourierRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := h.s.CreateCourier(req.Couriers)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(CreateCourierResponse{Couriers: couriers})
}

func (h handler) getCourierById(w http.ResponseWriter, r *http.Request) {
	x := newHelper("getCourierById", w, r)

	courierID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	courier, err := h.s.GetCourierById(courierID)
	if err != nil {
		x.WriteError(err)
		return
	}

	x.WriteResponse(courier)
}

func (h handler) getCouriers(w http.ResponseWriter, r *http.Request) {
	x := newHelper("getCouriers", w, r)

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	couriers, err := h.s.GetCouriers(limit, offset)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := GetCouriersResponse{
		Couriers: couriers,
		Limit:    limit,
		Offset:   offset,
	}

	x.WriteResponse(resp)
}

func (h handler) getCourierMetaInfo(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) couriersAssignments(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) createOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper("createOrder", w, r)

	var req CreateOrderRequest

	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := h.s.CreateOrder(req.Orders)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}

func (h handler) getOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper("getOrder", w, r)

	orderID, err := x.GetID()
	if err != nil {
		x.WriteError(err)
		return
	}

	order, err := h.s.GetOrder(orderID)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := order
	x.WriteResponse(resp)
}

func (h handler) getOrders(w http.ResponseWriter, r *http.Request) {
	x := newHelper("getOrders", w, r)

	limit, offset, err := x.GetLimitOffset(1, 0)
	if err != nil {
		x.WriteError(err)
		return
	}

	orders, err := h.s.GetOrders(limit, offset)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}

func (h handler) completeOrder(w http.ResponseWriter, r *http.Request) {
	x := newHelper("completeOrder", w, r)

	var req CompleteOrderRequestDto
	if err := x.ParseBody(&req); err != nil {
		x.WriteError(err)
		return
	}

	orders, err := h.s.ComleteOrder(req.CompleteInfo)
	if err != nil {
		x.WriteError(err)
		return
	}

	resp := orders
	x.WriteResponse(resp)
}

func (h handler) ordersAssign(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}
