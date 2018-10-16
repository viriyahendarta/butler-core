clear:
	@clear

build-api: clear
	@go build -o ./bin/api ./cmd/api

run-api: clear
	@make build-api
	@./bin/api

build:
	@make build-api

run:
	@make run-api

.PHONY: clear run-api build-api