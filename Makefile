# Makefile for Task Manager API

.PHONY: sqlc-generate build test run migrate-up migrate-down

sqlc-generate:
	@export PATH=$$PATH:$(shell go env GOPATH)/bin && sqlc generate

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v -cover ./...

run:
	go run cmd/server/main.go
