build:
	go build -o bin/http ./cmd/http

run:
	go run ./cmd/http

clean:
	rm -rf bin
	go clean