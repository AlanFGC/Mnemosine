package Service

import (
	"card-service/Model"
	"card-service/repository"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/openpgp/errors"

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
	Stop(ctx context.Context) error
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

	repository.CreateCardCollection(ctx, db)
	repository.CreateDeckCollection(ctx, db)
	repository.CreateUserIndex(ctx, db)

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
		log.Fatal("unable to get client: ", err)
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
	if card.Username == "" {
		return "", errors.InvalidArgumentError("username can't be null")
	}
	var cards []Model.UserFlashCard

	id, err := repository.InsertCards(ctx, s.db, cards)
	if err != nil {
		return "", err
	} else if len(id) != 1 {
		return "", errors.UnsupportedError("More than one id returned")
	}

	return id[0], nil
}

func (s *CardService) EditCard(ctx context.Context, card Model.UserFlashCard) error {
	if card.Username == "" {
		return errors.InvalidArgumentError("username can't be null")
	}
	return repository.UpdateById(ctx, s.db, card.ID.Hex(), card)
}

func (s *CardService) CreateDeck(ctx context.Context, deck Model.Deck) (string, error) {
	if deck.Username == "" {
		return "", errors.InvalidArgumentError("username can't be null")
	}
	deck.Cards = makeArrayObjectIdArrayUnique(deck.Cards)

	return repository.InsertOneDeck(ctx, s.db, deck)
}

func (s *CardService) EditDeck(ctx context.Context, deck Model.Deck) error {
	deck.Cards = makeArrayObjectIdArrayUnique(deck.Cards)
	return repository.UpdateDeckById(ctx, s.db, deck.ID, deck)
}

func (s *CardService) GetCardsByUser(ctx context.Context, username string) ([]Model.UserFlashCard, error) {
	return s.GetCardsByDeck(ctx, username)
}

func (s *CardService) GetCardsByDeck(ctx context.Context, deckID string) ([]Model.UserFlashCard, error) {
	return s.GetCardsByDeck(ctx, deckID)
}

func makeArrayObjectIdArrayUnique(ids []primitive.ObjectID) []primitive.ObjectID {
	set := map[primitive.ObjectID]bool{}
	var uniqueArray []primitive.ObjectID

	for _, val := range ids {
		if !set[val] {
			uniqueArray = append(uniqueArray, val)
			set[val] = true
		}
	}
	return uniqueArray
}
