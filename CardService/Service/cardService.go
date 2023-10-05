package Service

import (
	"card-service/Model/card"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	SayHello(ctx context.Context, name string) string
	//GetUserCard(ctx context.Context) (*card.UserFlashCard, error)
}

type CardService struct {
	context     context.Context
	databaseUrl string
}

func NewCardService(ctx context.Context, databaseUrl string) Service {
	return &CardService{
		context:     ctx,
		databaseUrl: databaseUrl,
	}
}

func (s *CardService) SayHello(ctx context.Context, name string) string {
	return "Hello, " + name
}

func (s *CardService) GetUserCard(ctx context.Context) (*card.UserFlashCard, error) {
	return &card.UserFlashCard{ID: primitive.ObjectID{}}, nil
}
