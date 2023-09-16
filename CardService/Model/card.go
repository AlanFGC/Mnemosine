package Model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
}

func CreateCardCollection(ctx context.Context, db *mongo.Database) error {
	collectionOptions := options.CreateCollection()
	err := db.CreateCollection(ctx, "UserFlashCard", collectionOptions)
	if err != nil {
		log.Fatal("Something went wrong", err)
		return err
	}

	return nil
}
