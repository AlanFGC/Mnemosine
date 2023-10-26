package repository

import (
	"card-service/Model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mtest "go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// DB NAME
const TestDBName = "testDB"

// DB USERNAMES
const TESTUSERNAME1 = "John Doe"
const TESTUSERNAME2 = "Juan Doe"
const TESTUSERNAME3 = "Joan Doe"
const TESTANSWER1 = "Correct Answer"
const TESTANSWER2 = "Incorrect Answer"
const TESTANSWER3 = "Dinosaurs are extinct"

func TestInsertOneCardIntegration(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Insert one card", func(mt *mtest.T) {

		cards := []Model.UserFlashCard{{
			Username: "JohnDoe",
			Title:    "Sample Card",
			Text:     "This is a sample flash card.",
			Answers: []Model.Answer{
				{
					Field:        1,
					Answer:       TESTANSWER1,
					QuestionType: Model.SingleAnswer,
				},
			},
			Media:       []string{"sample.jpg"},
			Lang:        []string{"en", "es"},
			Topics:      []string{"sample"},
			DateCreated: primitive.NewDateTimeFromTime(time.Now()),
		}}

		mockResponse := mtest.CreateSuccessResponse()
		mt.AddMockResponses(mockResponse)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id, err := InsertCards(ctx, mt.DB, cards)
		if err != nil {
			mt.Log("Error found when inserting a card:", err)
			mt.Fail()
		}

		assert.NotEqual(t, "", id)
	})

	mt.Run("Insert many card", func(mt *mtest.T) {

		cards := []Model.UserFlashCard{{
			Username: "JohnDoe",
			Title:    "Sample Card",
			Text:     "This is a sample flash card.",
			Answers: []Model.Answer{
				{
					Field:        1,
					Answer:       TESTANSWER1,
					QuestionType: Model.SingleAnswer,
				},
			},
			Media:       []string{"sample.jpg"},
			Lang:        []string{"en", "es"},
			Topics:      []string{"sample"},
			DateCreated: primitive.NewDateTimeFromTime(time.Now()),
		},
			{
				Username: "JohnDoe",
				Title:    "Sample Card",
				Text:     "This is a sample flash card.",
				Answers: []Model.Answer{
					{
						Field:        1,
						Answer:       TESTANSWER1,
						QuestionType: Model.SingleAnswer,
					},
				},
				Media:       []string{"sample.jpg"},
				Lang:        []string{"en", "es"},
				Topics:      []string{"sample"},
				DateCreated: primitive.NewDateTimeFromTime(time.Now()),
			},
			{
				Username: "JohnDoe",
				Title:    "Sample Card",
				Text:     "This is a sample flash card.",
				Answers: []Model.Answer{
					{
						Field:        1,
						Answer:       TESTANSWER1,
						QuestionType: Model.SingleAnswer,
					},
				},
				Media:       []string{"sample.jpg"},
				Lang:        []string{"en", "es"},
				Topics:      []string{"sample"},
				DateCreated: primitive.NewDateTimeFromTime(time.Now()),
			}}

		for _ = range cards {
			mt.AddMockResponses(mtest.CreateSuccessResponse())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id, err := InsertCards(ctx, mt.DB, cards)
		if err != nil {
			mt.Fatalf("Error found when inserting a card: %s", err)
		}

		assert.NotEqual(t, "", id)
	})
}

func TestGetCardsByUsernameIntegration(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("get cards by username", func(mt *mtest.T) {

		responses := []bson.D{
			{
				{"username", TESTUSERNAME1},
				{"title", "Sample Card 1"},
				{"text", "This is a sample flash card."},
				{"answers", bson.A{
					bson.D{
						{"field", 1},
						{"answers", bson.A{TESTANSWER1}},
						{"questionType", Model.SingleAnswer},
					},
				}},
				{"media", bson.A{"sample.jpg"}},
				{"lang", bson.A{"en", "es"}},
				{"topics", bson.A{"sample"}},
				{"dateCreated", primitive.NewDateTimeFromTime(time.Now())},
			},
			{
				{"username", TESTUSERNAME1},
				{"title", "Sample Card 2"},
				{"text", "This is a sample flash card."},
				{"answers", bson.A{
					bson.D{
						{"field", 1},
						{"answers", bson.A{TESTANSWER2}},
						{"questionType", Model.SingleAnswer},
					},
				}},
				{"media", bson.A{"sample.jpg"}},
				{"lang", bson.A{"en", "es"}},
				{"topics", bson.A{"sample"}},
				{"dateCreated", primitive.NewDateTimeFromTime(time.Now())},
			},
			{
				{"username", TESTUSERNAME1},
				{"title", "Sample Card 3"},
				{"text", "This is a sample flash card."},
				{"answers", bson.A{
					bson.D{
						{"field", 1},
						{"answers", bson.A{TESTANSWER3}},
						{"questionType", Model.SingleAnswer},
					},
				}},
				{"media", bson.A{"sample.jpg"}},
				{"lang", bson.A{"en", "es"}},
				{"topics", bson.A{"sample"}},
				{"dateCreated", primitive.NewDateTimeFromTime(time.Now())},
			},
		}

		find := mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", mt.DB.Name(), TestDBName),
			mtest.FirstBatch,
			responses...)
		fmt.Print(find, "\n")
		mt.AddMockResponses(find)

		// Call GetCardsByUsername.
		returnedCards, err := GetCardsByUsername(context.TODO(), mt.DB, TESTUSERNAME1, 0)
		if err != nil {
			mt.Fatalf("Error found when getting cards by username: %s", err)
		}

		assert.Equal(t, len(responses), len(returnedCards))
		for i, val := range returnedCards {
			assert.Equal(t, val.Username, TESTUSERNAME1)
			assert.Equal(t, val.Title, fmt.Sprintf("Sample Card %d", i+1))
			assert.Equal(t, val.Text, "This is a sample flash card.")
			assert.Equal(t, len(val.Answers), 1)
			if i == 0 {
				assert.Equal(t, val.Answers[0].Answer, TESTANSWER1)
			}
			if i == 1 {
				assert.Equal(t, val.Answers[0].Answer, TESTANSWER2)
			}
			if i == 2 {
				assert.Equal(t, val.Answers[0].Answer, TESTANSWER3)
			}
		}
	})
}
