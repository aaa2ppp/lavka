package handler

import (
	"log"
	"net/http"
)

type Service interface {
}

func New(service Service) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /ping", http.HandlerFunc(ping))
	return mux
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("ping from %v", r.RemoteAddr)
	w.Write([]byte("pong\n"))
}
