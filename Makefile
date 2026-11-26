.RECIPEPREFIX := >
-include .env
export

.PHONY: up down run migrate-up migrate-down test tidy

up:
> docker compose up -d

down:
> docker compose down

run:
> go run ./cmd/api

migrate-up:
> goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
> goose -dir migrations postgres "$(DATABASE_URL)" down

test:
> go test ./...

tidy:
> go mod tidy
