#!make
include .env

DB_URL=postgresql://${m_db_username}:${m_db_password}@localhost:5432/${m_db_name}?sslmode=disable

runOLD:
	go run ./cmd/web/main.go ./cmd/web/middleware.go ./cmd/web/router.go ./cmd/web/run.go 8000

run:
	go run ./cmd/web/main.go ./cmd/web/middleware.go ./cmd/web/router.go ./cmd/web/run.go ./cmd/web/send-email.go 8000

coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out   

db:
	docker run --name bookingdb -e POSTGRES_USER=${m_db_username} -e POSTGRES_PASSWORD=${m_db_password} -p 5432:5432 -d postgres

migrateUp:
	migrate -path migrations -database "${DB_URL}" -verbose up

migrateDown:
	migrate -path migrations -database "${DB_URL}" -verbose down

migrateCreate:
	migrate create -ext sql -dir migrations -seq create-user-table