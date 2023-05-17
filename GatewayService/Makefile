BIN := "./bin/app"
COMPOSE_PATH := "deployments/docker-compose.yml"

build:
	go build -v -o $(BIN) ./cmd/...

run: build
	$(BIN)

version: build
	$(BIN) version

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

up:
	docker compose -f $(COMPOSE_PATH) up -d --build --force-recreate

down:
	docker compose -f $(COMPOSE_PATH) down

.PHONY: build run build-img version test lint up down