package Server

import (
	"card-service/Service"
	"context"
)

type GrpcServer struct {
	pb      Service.CardService // Change "pb" to "pb.FlashCardServiceServer"
	service Service.Service
}

func (s *GrpcServer) CreateUserFlashCard(ctx context.Context, req *CreateFlashCardReq) (*CreateFlashCardRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) CreateDeck(ctx context.Context, req *CreateDeckReq) (*CreateDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) EditCard(ctx context.Context, req *EditCardReq) (*EditCardRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) EditDeck(ctx context.Context, req *EditDeckReq) (*EditDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) GetCardsByUsername(ctx context.Context, req *GetUserCardsByUsernameReq) (*GetUserCardsByUsernameRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GrpcServer) GetCardsByDeckId(ctx context.Context, req *GetCardsByDeckIdReq) (*GetCardsByDeckIdRes, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpcServer(s Service.Service) GrpcServer {
	return GrpcServer{
		service: s,
	}
}

func (s *GrpcServer) SayHello(ctx context.Context, in *Name) (*Hello, error) {
	return &Hello{Hello: s.service.SayHello(ctx, in.Name)}, nil
}

func (s *GrpcServer) mustEmbedUnimplementedCardServiceServer() {
	return
}
