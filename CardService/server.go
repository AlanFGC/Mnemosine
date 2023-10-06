package main

import (
	"card-service/Server"
	"card-service/Service"
	"context"
)

type GrpcServer struct {
	pb      Service.CardService
	service Service.Service
}

func (s *GrpcServer) CreateUserFlashCard(ctx context.Context, req *Server.CreateFlashCardReq) (*Server.CreateFlashCardRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) CreateDeck(ctx context.Context, req *Server.CreateDeckReq) (*Server.CreateDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) EditCard(ctx context.Context, req *Server.EditCardReq) (*Server.EditCardRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) EditDeck(ctx context.Context, req *Server.EditDeckReq) (*Server.EditDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) GetCardsByUsername(ctx context.Context, req *Server.GetUserCardsByUsernameReq) (*Server.GetUserCardsByUsernameRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) GetCardsByDeckId(ctx context.Context, req *Server.GetCardsByDeckIdReq) (*Server.GetCardsByDeckIdRes, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpcServer(s Service.Service) GrpcServer {
	return GrpcServer{
		service: s,
	}
}

func (s *GrpcServer) SayHello(ctx context.Context, in *Server.Name) (*Server.Hello, error) {
	return &Server.Hello{Hello: s.service.SayHello(ctx, in.Name)}, nil
}

func (s *GrpcServer) mustEmbedUnimplementedCardServiceServer() {
	return
}
