run:
	go run ./cmd/web/main.go ./cmd/web/middleware.go ./cmd/web/router.go ./cmd/web/run.go 8000
coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out   