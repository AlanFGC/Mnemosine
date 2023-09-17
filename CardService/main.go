package main

import (
	"card-service/Model"
	"card-service/Utilities"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	uri := "mongodb://localhost:27017"
	ctx := context.TODO()
	client, err := getClient(ctx, uri)
	if err != nil {
		fmt.Print("FAILED TO CONNECT")
		return
	}
	db := client.Database("mnemosine")
	//Model.CreateCardCollection(ctx, db)
	card := Model.UserFlashCard{
		Username: "Alan",
		Title:    "Sample Card",
		Text:     "This is a sample flash card.",
		Answers:  []string{"Answer 1", "Answer 2"},
		Media:    []string{"image.jpg", "audio.mp3"},
		Lang:     []string{"English", "Spanish"},
	}
	err = Model.InsertOneFlashCard(ctx, db, card)
	if err != nil {
		fmt.Print("FAILED TO INSERT")
		return
	}
	var size int
	size = 10
	cards := make([]Model.UserFlashCard, size)
	for i := 0; i < size; i++ {
		cards[i] = Model.UserFlashCard{
			Username: Utilities.GenerateRandString(5),
			Title:    Utilities.GenerateRandString(5),
			Text:     Utilities.GenerateRandString(5),
			Answers:  []string{Utilities.GenerateRandString(5), Utilities.GenerateRandString(5)},
			Media:    []string{Utilities.GenerateRandString(5), Utilities.GenerateRandString(5)},
			Lang:     []string{Utilities.GenerateRandString(5)},
		}
	}

	err = Model.InsertManyFlashCards(ctx, db, cards)
	if err != nil {
		fmt.Print("FAILED TO INSERT MANY")
	}

	//db.Drop(ctx)
	disconnectDB(context.TODO(), client)
}

func getClient(ctx context.Context, URI string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("unable to get client", err)
		return nil, err
	}
	return client, err
}

func disconnectDB(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		return
	}
}
