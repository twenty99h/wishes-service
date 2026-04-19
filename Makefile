DSN:=postgresql://postgres:postgres@localhost:5432/wishes?sslmode=disable

run:
	go run ./cmd/app

up:
	docker compose up -d

down:
	docker compose down

migrate-up:
	migrate -path migrations -database "$(DSN)" up

migrate-down:
	migrate -path migrations -database "$(DSN)" down 1