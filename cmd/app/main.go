package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"lavka/internal/api"
	"lavka/internal/config"
	"lavka/internal/middleware"
	"lavka/internal/swagger"
)

func main() {

	godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		logFatal("can't load config", err)
	}

	setupLogger(cfg.Logger)

	mux := http.NewServeMux()

	if err := swagger.Setup(cfg.Swagger, mux); err != nil {
		slog.Error("can't setup swagger", "error", err)
	}

	service := api.ServiceStub{}
	mux.Handle("/", middleware.Logging(api.New(service)))

	server := setupServer(cfg.Server, mux)

	log.Printf("startup http-server on %v", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logFatal("server fail", err)
	}
}

func logFatal(msg string, err error) {
	slog.Log(context.Background(), slog.LevelError+2, msg, "error", err)
	os.Exit(1)
}

func setupLogger(cfg config.Logger) {

	var h slog.Handler
	if cfg.PlainText {
		h = slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: cfg.Level,
			},
		)
	} else {
		h = slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: cfg.Level,
			},
		)
	}

	slog.SetDefault(slog.New(h))
}

func setupServer(cfg config.Server, router http.Handler) *http.Server {

	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	// setup graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		s := <-c
		log.Printf("got signal %v", s)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("can't shutdown server: %v", err)
		}
	}()

	return server
}
