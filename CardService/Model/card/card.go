package card

import (
	"card-service/Utilities"
	"context"
	"errors"
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
const PageSize = 1000

func CreateCardCollection(ctx context.Context, db *mongo.Database) error {
	err := db.CreateCollection(ctx, FlashCardCollectionName)
	if err != nil {
		log.Fatal("Something went wrong", err)
		return err
	}
	return nil
}

func CreateUserIndex(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(FlashCardCollectionName)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetName("username_index"),
	}
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Print("Username index created successfully.")
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

func GetCardsByUsername(ctx context.Context, db *mongo.Database, username string, page int) (
	[]UserFlashCard, error,
) {
	collection := db.Collection(FlashCardCollectionName)

	filter := bson.M{"username": strings.ToLower(username)}
	opts := options.Find().SetSkip(int64((page - 1) * PageSize)).SetLimit(int64(PageSize))

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
		err := cur.Decode(&card)
		if err != nil {
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

func FindByTitle(ctx context.Context, db *mongo.Database, title string, page int) ([]UserFlashCard, error) {
	coll := db.Collection(FlashCardCollectionName)

	filter := bson.D{{"title", title}}
	opts := options.Find().SetSkip(int64((page - 1) * PageSize)).SetLimit(int64(PageSize))
	queryRes, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var cards []UserFlashCard

	// Iterate through the result set and decode each document into a UserFlashCard struct
	for queryRes.Next(ctx) {
		var card UserFlashCard
		err := queryRes.Decode(&card)
		if err != nil {
			log.Fatal("Error decoding flashcard:", err)
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func UpdateById(ctx context.Context, db *mongo.Database, ID string, card UserFlashCard) error {

	coll := db.Collection(FlashCardCollectionName)
	// Convert the string ID to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	// Create a filter for the document to be updated
	filter := bson.M{"_id": objectID}

	// Create an update with the changes you want to apply
	update := bson.M{
		"$set": bson.M{
			"username":    card.Username,
			"title":       card.Title,
			"text":        card.Text,
			"answers":     card.Answers,
			"media":       card.Media,
			"languages":   card.Lang,
			"topics":      card.Topics,
			"dateCreated": card.DateCreated,
		},
	}

	// Specify additional options, if needed
	opts := options.Update().SetUpsert(false)

	// Perform the update operation
	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	// Check if the update affected any documents
	if result.ModifiedCount == 0 {
		return errors.New("no documents were updated")
	}

	return nil
}
