package config

import (
	"log/slog"
	"time"

	"lavka/internal/getenv"
)

type Logger struct {
	Level     slog.Level
	PlainText bool
}

type Server struct {
	Addr            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type DB struct {
}

type Swagger struct {
}

type Config struct {
	Logger  Logger
	Server  Server
	DB      DB
	Swagger Swagger
}

func Load() (Config, error) {
	ge := getenv.New()
	const required = true

	return Config{
		Logger: Logger{
			Level:     ge.LogLevel("LOG_LEVEL", !required, slog.LevelInfo),
			PlainText: ge.Bool("LOG_PLAIN_TEXT", !required, false),
		},
		Server: Server{
			Addr:            ge.String("SERVER_ADDR", !required, ":8080"),
			WriteTimeout:    ge.Duration("SERVER_WRITE_TIMEOUT", !required, 10*time.Second),
			ReadTimeout:     ge.Duration("SERVER_READ_TIMEOUT", !required, 10*time.Second),
			ShutdownTimeout: ge.Duration("SERVER_SHUTDOWN_TIMEOUT", !required, 10*time.Second),
		},
		DB: DB{},
	}, ge.Err()
}