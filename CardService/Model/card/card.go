package card

import (
	"card-service/Utilities"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

type Answer struct {
	Field            int      `bson:"field"`
	Answers          []string `bson:"answers"`
	IncorrectAnswers []string `bson:"incorrectAnswers, omitempty"`
	Explanation      string   `bson:"explanation,omitempty"`
}

type UserFlashCard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username, omitempty"`
	Title       string             `bson:"title"`
	Text        string             `bson:"text"`
	Answers     []Answer           `bson:"answers"`
	Media       []string           `bson:"media"`
	Lang        []string           `bson:"languages"`
	Topics      []string           `bson:"topics"`
	DateCreated primitive.DateTime `bson:"dateCreated"`
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

func InsertOne(ctx context.Context, db *mongo.Database, card UserFlashCard) error {
	collection := db.Collection(FlashCardCollectionName)
	_, err := collection.InsertOne(ctx, card)
	if err != nil {
		log.Fatal("Couldn't insert one flashcard")
	}
	return nil
}

func InsertMany(ctx context.Context, db *mongo.Database, cards []UserFlashCard) ([]string, error) {
	collection := db.Collection(FlashCardCollectionName)
	cardInterface := Utilities.ToInterfaceSlice(cards)

	res, err := collection.InsertMany(ctx, cardInterface)
	if err != nil {
		log.Fatal("Couldn't insert one flashcard")
	}
	result := []string{}

	for _, value := range res.InsertedIDs {
		if ID, ok := value.(primitive.ObjectID); ok {
			result = append(result, ID.Hex())
		}
	}

	return result, nil
}

func GetCardById(ctx context.Context, db *mongo.Database, ID string) (UserFlashCard, error) {
	var card UserFlashCard

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return card, err
	}

	filter := bson.M{"_id": objectID}

	err = db.Collection(FlashCardCollectionName).FindOne(ctx, filter).Decode(&card)
	if err != nil {
		return card, err
	}
	return card, nil
}

func GetCardsByUsername(ctx context.Context, db *mongo.Database, username string, page int, pageSize int) ([]UserFlashCard, error) {
	collection := db.Collection(FlashCardCollectionName)

	// Define a filter to search for flashcards with the given username
	filter := bson.M{"username": strings.ToLower(username)} // Convert username to lowercase for case-insensitive search

	// Define options for pagination
	opts := options.Find().SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))

	// Find the flashcards matching the filter and pagination options
	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Fatal("Error while finding flashcards:", err)
		return nil, err
	}

	var cards []UserFlashCard

	// Iterate through the result set and decode each document into a UserFlashCard struct
	for cur.Next(ctx) {
		var card UserFlashCard
		if err := cur.Decode(&card); err != nil {
			log.Fatal("Error decoding flashcard:", err)
			return nil, err
		}
		cards = append(cards, card)
	}

	// Check for errors from iterating over the cursor
	if err := cur.Err(); err != nil {
		log.Fatal("Error during cursor iteration:", err)
		return nil, err
	}

	return cards, nil
}
