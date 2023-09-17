package Model

import (
	"card-service/Utilities"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserFlashCard struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username, omitempty"`
	Title    string             `bson:"title"`
	Text     string             `bson:"text"`
	Answers  []string           `bson:"answers"`
	Media    []string           `bson:"media"`
	Lang     []string           `bson:"languages"`
	Topics   []string           `bson:"topics"`
}

const FlashCardCollectionName = "UserFlashCard"

func CreateCardCollection(ctx context.Context, db *mongo.Database) error {
	err := db.CreateCollection(ctx, FlashCardCollectionName)
	if err != nil {
		log.Fatal("Something went wrong", err)
		return err
	}
	return nil
}

func InsertOneFlashCard(ctx context.Context, db *mongo.Database, card UserFlashCard) error {
	collection := db.Collection(FlashCardCollectionName)
	_, err := collection.InsertOne(ctx, card)
	if err != nil {
		log.Fatal("Couldn't insert one flashcard")
	}
	return nil
}

func InsertManyFlashCards(ctx context.Context, db *mongo.Database, cards []UserFlashCard) error {
	collection := db.Collection(FlashCardCollectionName)
	cardInterface := Utilities.ToInterfaceSlice(cards)
	_, err := collection.InsertMany(ctx, cardInterface)
	if err != nil {
		log.Fatal("Couldn't insert one flashcard")
	}
	return nil
}
