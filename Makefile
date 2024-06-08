.DEFAULT_GOAL := start

.PHONY: fmt vet build-api build-cli start

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build-api: vet
	@go build cmd/api/main.go

build-cli: vet
	@go build -o nytaxis cmd/cli/main.go

start:
	@docker compose up
