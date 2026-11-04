.PHONY: run build test lint

run:
	go run ./cmd/event-service

build:
	go build -o bin/event-service ./cmd/event-service

test:
	go test ./... -race -cover

lint:
	golangci-lint run ./...
