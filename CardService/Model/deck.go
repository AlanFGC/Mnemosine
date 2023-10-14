package Model

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DeckCollectionName = "DeckCollection"
const DeckPageSize = 1000

func CreateDeckCollection(ctx context.Context, db *mongo.Database) error {

	filter := bson.D{{Key: "name", Value: DeckCollectionName}}
	names, err := db.ListCollectionNames(ctx, filter, nil)
	if err != nil {
		log.Fatal("List of names couldn't be retrieved for Deck Collection: ", err)
		return err
	}

	for _, name := range names {
		if name == DeckCollectionName {
			return nil
		}
	}

	err = db.CreateCollection(ctx, DeckCollectionName)
	if err != nil {
		log.Fatal("Failed to create collection Decks: ", err)
		return err
	}
	return nil
}

func InsertOneDeck(ctx context.Context, db *mongo.Database, deck Deck) (string, error) {
	collection := db.Collection(DeckCollectionName)
	res, err := collection.InsertOne(ctx, deck)
	if err != nil {
		return "", err
	}

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("Failed to convert InsertedID to ObjectID")
	}

	return insertedID.Hex(), nil
}

func GetDeckByUsername(ctx context.Context, db *mongo.Database, username string, page int) ([]Deck, error) {
	collection := db.Collection(DeckCollectionName)
	filter := bson.D{{Key: "username", Value: username}}
	opts := options.Find().SetSkip(int64((page - 1) * DeckPageSize)).SetLimit(int64(DeckPageSize))

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
	filter := bson.D{{Key: "_id", Value: deckID}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "cardIds", Value: bson.D{{Key: "$each", Value: cardIDs}}},
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
	filter := bson.D{{Key: "_id", Value: deckID}}

	opts := options.Update().SetUpsert(false)
	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "cardIds", Value: bson.D{{Key: "$each", Value: cardIDs}}},
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

func UpdateDeckById(ctx context.Context, db *mongo.Database, objectID primitive.ObjectID, deck Deck) error {
	coll := db.Collection(DeckCollectionName)

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"title":       deck.Title,
			"username":    deck.Username,
			"cardAuthors": deck.CardAuthors,
			"topics":      deck.Topics,
			"cardIds":     deck.Cards,
		},
	}
	opts := options.Update().SetUpsert(false)

	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
