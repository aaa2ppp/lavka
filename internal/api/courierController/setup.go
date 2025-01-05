package courierController

import (
	"context"
	"net/http"

	"lavka/internal/model"
)

type CourierService interface {
	CreateCourier(context.Context, []model.CreateCourierDto) ([]model.CourierDto, error)
	GetCourierById(_ context.Context, courierID uint64) (model.CourierDto, error)
	GetCouriers(_ context.Context, limit, offset int) ([]model.CourierDto, error)
}

type controller struct {
	CourierService
}

func Setup(mux *http.ServeMux, service CourierService) {
	c := controller{service}
	mux.Handle("POST /couriers", http.HandlerFunc(c.createCourier))
	mux.Handle("GET  /couriers/{id}", http.HandlerFunc(c.getCourierById))
	mux.Handle("GET  /couriers", http.HandlerFunc(c.getCouriers))
	mux.Handle("GET  /couriers/meta-info/{id}", http.HandlerFunc(c.getCourierMetaInfo))
	mux.Handle("GET  /couriers/assignments", http.HandlerFunc(c.couriersAssignments))
}
