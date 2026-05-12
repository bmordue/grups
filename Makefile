.PHONY: up down run migrate

up:
	docker compose up -d

down:
	docker compose down

run:
	go run ./cmd/server

migrate:
	psql "$$DATABASE_URL" -f db/migrations/001_init.sql
