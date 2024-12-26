package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lavka/internal/handler"
)

func main() {

	// TODO: setup the logger

	service := struct { // fake service
		handler.Service
	}{}

	server := http.Server{
		Addr:         ":8080",
		Handler:      handler.New(service),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// setup graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		s := <-c
		log.Printf("got signal %v", s)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("can't shutdown server: %v", err)
		}
	}()


	log.Printf("startup http-server on %v", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server fail: %v", err)
	}
}
