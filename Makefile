DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
MIGRATIONS_DIR=backend/migrations

up:
	docker compose up -d

build:
	docker compose down -v --rmi all && docker compose build --no-cache && docker compose up -d

build-up: build migrate-up

down:
	docker compose down

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down