build:
	@go mod tidy
	@go build -o bin/cardService

run: build
	./bin/cardService

test:
	go test -v ./...