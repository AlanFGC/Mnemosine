package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type Deck struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"name"`
	Username    string               `bson:"username,omitempty"`
	CardAuthors []string             `bson:"cardAuthors"`
	Topics      []string             `bson:"topics"`
	Cards       []primitive.ObjectID `bson:"cardIds"`
}
