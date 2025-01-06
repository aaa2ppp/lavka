package courierController

import (
	"context"
	"net/http"

	"lavka/internal/middleware"
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

var limitRPS = middleware.LimitRPS

func Setup(mux *http.ServeMux, service CourierService, rps int) {
	c := controller{service}
	mux.Handle("POST /couriers", limitRPS(rps, http.HandlerFunc(c.createCourier)))
	mux.Handle("GET  /couriers/{id}", limitRPS(rps, http.HandlerFunc(c.getCourierById)))
	mux.Handle("GET  /couriers", limitRPS(rps, http.HandlerFunc(c.getCouriers)))
	mux.Handle("GET  /couriers/meta-info/{id}", limitRPS(rps, http.HandlerFunc(c.getCourierMetaInfo)))
	mux.Handle("GET  /couriers/assignments", limitRPS(rps, http.HandlerFunc(c.couriersAssignments)))
}
