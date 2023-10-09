package main

import (
	"card-service/Server"
	"card-service/Service"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const database = "mongodb://localhost:27017"
const port = ":50051"
const dbName = "mnemosine"

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	cardService, err := Service.NewCardService(context.Background(), database, dbName)
	if err != nil {
		log.Fatal("Failed to create card service.")
	}

	defer cardService.StopClean(context.Background())

	loggingService, err := Service.NewLoggingService(cardService)
	if err != nil {
		log.Fatal("Failed to create logging service.")
	}
	cardServer := Server.NewGrpcServer(loggingService)
	s := grpc.NewServer()

	// TODO
	// Change the parsing of data as the responsibility of grpc server, service should only handle native
	// data structures, such that we can decouple server and service.
	Server.RegisterCardServiceServer(s, cardServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	defer s.Stop()
}
