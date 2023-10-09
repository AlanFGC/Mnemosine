package Service

import (
	"card-service/Model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	SayHello(ctx context.Context, name string) string
	CreateUserCard(ctx context.Context, card Model.UserFlashCard) (string, error)
	EditCard(ctx context.Context, card Model.UserFlashCard) error
	CreateDeck(ctx context.Context, deck Model.Deck) (string, error)
	EditDeck(ctx context.Context, deck Model.Deck) error
	GetCardsByUser(ctx context.Context, username string) ([]Model.UserFlashCard, error)
	GetCardsByDeck(ctx context.Context, deckID string) ([]Model.UserFlashCard, error)
}

type CardService struct {
	context      context.Context
	databaseUrl  string
	databaseName string
	db           *mongo.Database
	client       *mongo.Client
}

func NewCardService(ctx context.Context, databaseUrl string, databaseName string) (Service, error) {
	client, err := getClient(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}
	db := client.Database(databaseName)


	Model.CreateCardCollection(ctx, db)
	Model.CreateDeckCollection(ctx, db)
	Model.CreateUserIndex(ctx, db)


	return &CardService{
		context:      ctx,
		databaseUrl:  databaseUrl,
		databaseName: databaseName,
		db:           db,
		client:       client,
	}, nil
}

func getClient(ctx context.Context, URI string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("unable to get client", err)
		return nil, err
	}
	return client, err
}

func (s *CardService) Stop(ctx context.Context) error {
	err := s.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardService) StopClean(ctx context.Context) error {
	err := s.db.Drop(ctx)
	if err != nil {
		return err
	}
	err = s.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// METHODS FOR DATA

func (s *CardService) SayHello(ctx context.Context, name string) string {
	return "Hello " + name
}

func (s *CardService) CreateUserCard(ctx context.Context, card Model.UserFlashCard) (string, error) {
	return Model.InsertOneCard(ctx, s.db, card)
}

func (s *CardService) EditCard(ctx context.Context, card Model.UserFlashCard) error {
	return Model.UpdateById(ctx, s.db, card.ID.Hex(), card)
}

func (s *CardService) CreateDeck(ctx context.Context, deck Model.Deck) (string, error) {
	return Model.
}

func (s *CardService) EditDeck(ctx context.Context, deck Model.Deck) error {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) GetCardsByUser(ctx context.Context, username string) ([]Model.UserFlashCard, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) GetCardsByDeck(ctx context.Context, deckID string) ([]Model.UserFlashCard, error) {
	//TODO implement me
	panic("implement me")
}
