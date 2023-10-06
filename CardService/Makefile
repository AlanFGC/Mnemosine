build:
	go mod tidy
	go build -o bin/cardService

run: build
	./bin/cardService

test:
	grpcurl -plaintext -proto ./Server/cardService.proto -d '{"name":"World"}' localhost:50051 CardService/SayHello
	go test -v ./...
