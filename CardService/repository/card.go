package repository

import (
	"card-service/Model"
	"card-service/Utilities"
	"context"
	"errors"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const FlashCardCollectionName = "UserFlashCard"
const PageSize = 1000

// SETUP

func CreateCardCollection(ctx context.Context, db *mongo.Database) error {
	filter := bson.D{{Key: "name", Value: FlashCardCollectionName}}
	names, err := db.ListCollectionNames(ctx, filter, nil)
	if err != nil {
		log.Fatal("List of names couldn't be retrieved for Card Collection: ", err)
		return err
	}

	for _, name := range names {
		if name == FlashCardCollectionName {
			return nil
		}
	}

	err = db.CreateCollection(ctx, FlashCardCollectionName)
	if err != nil {
		log.Fatal("Failed to create collection flashcards: ", err)
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

	return nil
}

// OPERATIONS

func InsertCards(ctx context.Context, db *mongo.Database, cards []Model.UserFlashCard) ([]string, error) {
	collection := db.Collection(FlashCardCollectionName)
	cardInterface := Utilities.ToInterfaceSlice(cards)

	res, err := collection.InsertMany(ctx, cardInterface)
	if err != nil {
		log.Fatal("Couldn't insert one flashcard")
	}
	var result []string

	for _, value := range res.InsertedIDs {
		if ID, ok := value.(primitive.ObjectID); ok {
			result = append(result, ID.Hex())
		}
	}

	return result, nil
}

func GetCardByIds(ctx context.Context, db *mongo.Database, IDs []string) ([]Model.UserFlashCard, error) {
	var cards []Model.UserFlashCard
	var newId primitive.ObjectID
	var objectIDs []primitive.ObjectID
	var err error

	for _, id := range IDs {
		newId, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, newId)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	cursor, err := db.Collection(FlashCardCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var card Model.UserFlashCard
		err := cursor.Decode(&card)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCardsByUsername(ctx context.Context, db *mongo.Database, username string, page int) (
	[]Model.UserFlashCard, error,
) {
	collection := db.Collection(FlashCardCollectionName)

	filter := bson.M{"username": strings.ToLower(username)}
	opts := options.Find().SetSkip(int64((page - 1) * PageSize)).SetLimit(int64(PageSize))

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Fatal("Error while finding flashcards:", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var cards []Model.UserFlashCard

	for cur.Next(ctx) {
		var card Model.UserFlashCard
		err := cur.Decode(&card)
		if err != nil {
			log.Fatal("Error decoding flashcard:", err)
			return nil, err
		}
		cards = append(cards, card)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("Error during cursor iteration:", err)
		return nil, err
	}

	return cards, nil
}

func FindByTitle(ctx context.Context, db *mongo.Database, title string, page int) ([]Model.UserFlashCard, error) {
	coll := db.Collection(FlashCardCollectionName)

	filter := bson.D{{Key: "title", Value: title}}
	opts := options.Find().SetSkip(int64((page - 1) * PageSize)).SetLimit(int64(PageSize))
	queryRes, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var cards []Model.UserFlashCard

	// Iterate through the result set and decode each document into a UserFlashCard struct
	for queryRes.Next(ctx) {
		var card Model.UserFlashCard
		err := queryRes.Decode(&card)
		if err != nil {
			log.Fatal("Error decoding flashcard:", err)
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func UpdateById(ctx context.Context, db *mongo.Database, ID string, card Model.UserFlashCard) error {

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

	opts := options.Update().SetUpsert(false)

	// Perform the update
	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no documents were updated")
	}

	return nil
}
