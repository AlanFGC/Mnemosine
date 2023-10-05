package Server

import (
	"card-service/Service"
	"context"
)

type GrpcServer struct {
	pb      Service.CardService // Change "pb" to "pb.FlashCardServiceServer"
	service Service.Service
}

func NewGrpcServer(s Service.Service) GrpcServer {
	return GrpcServer{
		service: s,
	}
}

func (s *GrpcServer) mustEmbedUnimplementedCardServiceServer() {
	return
}

func (s *GrpcServer) SayHello(ctx context.Context, in *Name) (*Hello, error) {
	return &Hello{Hello: s.service.SayHello(ctx, in.Name)}, nil
}
