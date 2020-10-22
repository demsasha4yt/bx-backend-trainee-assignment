SHELL := /bin/bash

.PHONY: migrate
migrate:
	migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations up

.PHONY: migrate_down
migrate_down:
	migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations down

.PHONY: setup
setup: migrate

.PHONY: build_mock
build_mock:
	go build -v ./cmd/avitomock

.PHONY: run_mock
run_mock: setup build_mock
	./avitomock

.PHONY: build
build:
	go build -v ./cmd/bx
	
.PHONY: run
run: setup build
	source .env && ./bx

.PHONY: test
test:
	go test -v -race ./...

.DEFAULT_GOAL := build