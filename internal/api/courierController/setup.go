package courierController

import (
	"net/http"

	m "lavka/internal/api/model"
)

type Service interface {
	CreateCourier(couriers []m.CreateCourierDto) ([]m.CourierDto, error)
	GetCourierById(courierID uint64) (m.CourierDto, error)
	GetCouriers(limit, offset int) ([]m.CourierDto, error)
}

type controller struct {
	Service
}

func Setup(mux *http.ServeMux, service Service) {
	c := controller{service}
	mux.Handle("POST /couriers", http.HandlerFunc(c.createCourier))
	mux.Handle("GET  /couriers/{id}", http.HandlerFunc(c.getCourierById))
	mux.Handle("GET  /couriers", http.HandlerFunc(c.getCouriers))
	mux.Handle("GET  /couriers/meta-info/{id}", http.HandlerFunc(c.getCourierMetaInfo))
	mux.Handle("GET  /couriers/assignments", http.HandlerFunc(c.couriersAssignments))
}
