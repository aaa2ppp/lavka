package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	"lavka/internal/api"
	"lavka/internal/api/courierController"
	"lavka/internal/api/orderController"
	"lavka/internal/config"
	"lavka/internal/middleware"
	"lavka/internal/repo/courierRepo"
	"lavka/internal/repo/orderRepo"
	"lavka/internal/swagger"
)

type service struct {
	orderController.OrderService
	courierController.CourierService
}

func main() {

	godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		logFatal("can't load config", err)
	}

	setupLogger(cfg.Logger)

	db, err := openDB(cfg.DB)
	if err != nil {
		logFatal("can't open db", err)
	}
	defer db.Close()

	// up migrations
	if err := goose.Up(db, "migrations"); err != nil {
		logFatal("can't up migrations", err)
	}

	mux := http.NewServeMux()

	if err := swagger.Setup(cfg.Swagger, mux); err != nil {
		slog.Error("can't setup swagger", "error", err)
	}

	service := service{
		OrderService:   orderRepo.New(db),
		CourierService: courierRepo.New(db),
	}
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

func openDB(cfg config.DB) (*sql.DB, error) {
	const op = "openDB"

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
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
