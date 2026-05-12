
.PHONY: run migrate

run:
	CGO_ENABLED=0 go run ./cmd/server

migrate:
	mkdir -p data
	sqlite3 ./data/grups.db < db/migrations/001_init.sql
