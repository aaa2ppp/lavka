.PHONY: build run

build:
	go build -o bin/app ./cmd/app/main.go

run:
	bin/app
