package Model

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mtest "go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertOneCardIntegration(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Insert one card", func(mt *mtest.T) {

		cards := []UserFlashCard{{
			ID:       primitive.NewObjectID(),
			Username: "JohnDoe",
			Title:    "Sample Card",
			Text:     "This is a sample flash card.",
			Answers: []Answer{
				{
					Field:        1,
					Answers:      []string{"Correct Answer"},
					QuestionType: SingleAnswer,
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
}
