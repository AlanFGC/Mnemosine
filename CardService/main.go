package main

import (
	"card-service/Model"
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
