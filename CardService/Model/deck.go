package Model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
