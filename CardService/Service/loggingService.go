package Service

import (
	"card-service/Server"
	"context"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) (Service, error) {
	return &LoggingService{
		next: next,
	}, nil
}

func (s LoggingService) SayHello(ctx context.Context, name string) string {
	return s.next.SayHello(ctx, name)
}

func (s LoggingService) CreateUserCard(ctx context.Context, in *Server.CreateFlashCardReq) (*Server.CreateFlashCardRes, error) {
	return s.next.CreateUserCard(ctx, in)
}

func (s LoggingService) EditCard(ctx context.Context, in *Server.EditCardReq) (*Server.EditCardRes, error) {
	return s.EditCard(ctx, in)
}

func (s LoggingService) CreateDeck(ctx context.Context, in *Server.CreateDeckReq) (*Server.CreateDeckRes, error) {
	return s.CreateDeck(ctx, in)
}

func (s LoggingService) EditDeck(ctx context.Context, in *Server.EditDeckReq) (*Server.EditDeckReq, error) {
	return s.EditDeck(ctx, in)
}

func (s LoggingService) GetCardsByUser(ctx context.Context, in *Server.GetUserCardsByUsernameReq) (*Server.GetUserCardsByUsernameRes, error) {
	return s.next.GetCardsByUser(ctx, in)
}

func (s LoggingService) GetCardsByDeck(ctx context.Context, in *Server.GetCardsByDeckIdReq) (*Server.GetCardsByDeckIdRes, error) {
	return s.next.GetCardsByDeck(ctx, in)
}
