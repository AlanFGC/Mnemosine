package main

import (
	"card-service/Model/card"
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
	flashCard := card.UserFlashCard{
		Username: "Alan",
		Title:    "Sample Card",
		Text:     "This is a sample flash card.",
		Answers: []card.Answer{
			{
				Field:            0,
				Answers:          []string{"Answer1, Answer2"},
				IncorrectAnswers: []string{"bad1", "bad2"},
				Explanation:      "The answer one and answer two are correct because they are sample answers.",
			},
		},
		Media: []string{"image.jpg", "audio.mp3"},
		Lang:  []string{"English", "Spanish"},
	}
	err = card.InsertOne(ctx, db, flashCard)
	if err != nil {
		fmt.Print("FAILED TO INSERT")
		return
	}
	var size int
	size = 10
	cards := make([]card.UserFlashCard, size)
	for i := 0; i < size; i++ {
		cards[i] = card.UserFlashCard{
			Username: Utilities.GenerateRandString(5),
			Title:    Utilities.GenerateRandString(5),
			Text:     Utilities.GenerateRandString(5),
			Answers: []card.Answer{
				{
					Field:            1,
					Answers:          []string{"example"},
					IncorrectAnswers: []string{"incorrect"},
					Explanation:      "The answer one and answer two are correct because they are sample answers.",
				},
			},
			Media: []string{Utilities.GenerateRandString(5), Utilities.GenerateRandString(5)},
			Lang:  []string{Utilities.GenerateRandString(5)},
		}
	}

	var IDs []string
	IDs, err = card.InsertMany(ctx, db, cards)
	if err != nil || len(IDs) < len(cards) {
		fmt.Print("FAILED TO INSERT MANY")
	}

	queryResult, err := card.GetCardById(ctx, db, IDs[0])
	fmt.Print(queryResult)
	if err != nil {
		fmt.Print("FAILED TO query by one by id")
	}

	multipleCards, err := card.GetCardsByUsername(ctx, db, "Alan", 0, 100000000)
	if err != nil {
		fmt.Print("FAILED TO query by username")
	}

	fmt.Print(multipleCards)
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
