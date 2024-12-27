package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"lavka/internal/api"
	"lavka/internal/middleware"
)

func main() {

	// setup logger
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	// create router
	router := http.NewServeMux()

	// setup swagger
	// TODO: кто на ком стоял? а проще можно?
	swaggerDoc, err := os.ReadFile("docs/tz/openapi.json")
	if err != nil {
		log.Fatal(err)
	}
	router.Handle("GET /swagger/doc.json", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.Write(swaggerDoc) }))
	router.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	// setup api
	service := api.ServiceStub{}
	router.Handle("/", api.New(service))

	// setup server
	server := http.Server{
		Addr:         ":8080",
		Handler:      middleware.Logging(router),
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

	// startup server
	log.Printf("startup http-server on %v", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server fail: %v", err)
	}
}
