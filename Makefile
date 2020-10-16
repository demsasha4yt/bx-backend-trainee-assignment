SHELL := /bin/bash

.PHONY: setup
setup: migrate

.PHONY: build
build:
	go build -v ./cmd/bx

.PHONY: setup run
run: build
	source .env && ./bx

.PHONY: test
test:
	go test -v -race ./...

.PHONY: migrate
migrate:
	migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations up

.PHONY: migrate_down
migrate_down:
	migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations down


.DEFAULT_GOAL := build

# migrate create -ext sql -dir migrations -seq create_customers_table
# migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations up
# migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations down
