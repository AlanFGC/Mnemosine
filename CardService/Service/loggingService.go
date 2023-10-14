package Service

import (
	"card-service/Model"
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

func (s LoggingService) CreateUserCard(ctx context.Context, card Model.UserFlashCard) (string, error) {
	return s.next.CreateUserCard(ctx, card)
}

func (s LoggingService) EditCard(ctx context.Context, card Model.UserFlashCard) error {
	return s.next.EditCard(ctx, card)
}

func (s LoggingService) CreateDeck(ctx context.Context, deck Model.Deck) (string, error) {
	return s.next.CreateDeck(ctx, deck)
}

func (s LoggingService) EditDeck(ctx context.Context, deck Model.Deck) error {
	return s.next.EditDeck(ctx, deck)
}

func (s LoggingService) GetCardsByUser(ctx context.Context, username string) ([]Model.UserFlashCard, error) {
	return s.next.GetCardsByUser(ctx, username)
}

func (s LoggingService) GetCardsByDeck(ctx context.Context, deckID string) ([]Model.UserFlashCard, error) {
	return s.next.GetCardsByUser(ctx, deckID)
}

func (s LoggingService) Stop(ctx context.Context) error {
	return s.next.Stop(ctx)
}
