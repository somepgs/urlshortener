.PHONY: run lint test

run:
	go run ./cmd/urlshortener

lint:
	golangci-lint run ./...

test:
	go test ./...