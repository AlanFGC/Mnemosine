build:
	@go mod tidy
	@go build -o bin/fact

run: build
	./bin/fact

test:
	go test -v ./...