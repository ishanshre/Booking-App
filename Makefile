run:
	go run ./cmd/web/ 8000

coverage:
	go test -coverprofile=coverage.out && go tool cover -html=coverage.out   