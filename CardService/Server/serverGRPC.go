package Server

import (
	"card-service/Proto/generated/card-service/pb"
	"card-service/Service"
	"context"
)

type GrpcServer struct {
	service Service.Service
}

func NewGrpcServer(s Service.Service) GrpcServer {
	return GrpcServer{
		service: s,
	}
}

func (s *GrpcServer) SayHello(ctx context.Context, in *pb.Name) (*pb.Hello, error) {
	return &pb.Hello{Hello: s.service.SayHello(ctx, in.Name)}, nil
}
