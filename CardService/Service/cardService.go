package Service

import (
	"card-service/Model/card"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	SayHello(ctx context.Context) string
	//GetUserCard(ctx context.Context) (*card.UserFlashCard, error)
}

type CardService struct {
	databaseUrl string
}

func NewCardService(databaseUrl string) Service {
	return &CardService{
		databaseUrl: databaseUrl,
	}
}

func (s *CardService) SayHello(ctx context.Context) string {
	return "Hello World"
}

func (s *CardService) GetUserCard(ctx context.Context) (*card.UserFlashCard, error) {
	return &card.UserFlashCard{ID: primitive.ObjectID{}}, nil
}
