package Service

import (
	"card-service/Model/card"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetUserCard(ctx context.Context) (*card.UserFlashCard, error)
}

type CardService struct {
	databaseUrl string
}

func (s *CardService) GetUserCard(ctx context.Context) (*card.UserFlashCard, error) {
	return &card.UserFlashCard{ID: primitive.ObjectID{}}, nil
}
