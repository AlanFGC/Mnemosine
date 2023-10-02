package deck

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Deck struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"name"`
	Username    string               `bson:"username,omitempty"`
	CardAuthors []string             `bson:"cardAuthors"`
	Topics      []string             `bson:"topics"`
	Cards       []primitive.ObjectID `bson:"cardIds"`
}

const DeckCollectionName = "DeckCollection"
const PageSize = 1000

func CreateDeckCollection(ctx context.Context, db *mongo.Database) error {
	err := db.CreateCollection(ctx, DeckCollectionName)
	if err != nil {
		log.Fatal("Something went wrong", err)
		return err
	}
	return nil
}

func InsertOneDeck(ctx context.Context, db *mongo.Database, deck Deck) {
	collection := db.Collection(DeckCollectionName)
	_, err := collection.InsertOne(ctx, deck)
	if err != nil {
		log.Print("Failed to insert one deck")
	}
}

func GetDeckByUsername(ctx context.Context, db *mongo.Database, username string, page int) ([]Deck, error) {
	collection := db.Collection(DeckCollectionName)
	filter := bson.D{{"username", username}}
	opts := options.Find().SetSkip(int64((page - 1) * PageSize)).SetLimit(int64(PageSize))

	queryRes, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var decks []Deck

	for queryRes.Next(ctx) {
		var deck Deck
		err := queryRes.Decode(&deck)
		if err != nil {
			log.Fatal("Error decoding deck:", err)
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, nil
}

func InsertCardsToDeck(ctx context.Context, db *mongo.Database, deckID string, cardIDs []string) error {
	collection := db.Collection(DeckCollectionName)
	filter := bson.D{{"_id", deckID}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{"$addToSet", bson.D{
			{"cardIds", bson.D{{"$each", cardIDs}}},
		}},
	}

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no documents were updated")
	}

	return nil
}

func RemoveCardsFromDeck(ctx context.Context, db *mongo.Database, deckID string, cardIDs []string) error {
	collection := db.Collection(DeckCollectionName)
	filter := bson.D{{"_id", deckID}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{"$pull", bson.D{
			{"cardIds", bson.D{{"$each", cardIDs}}},
		}},
	}

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no documents were updated")
	}

	return nil
}
