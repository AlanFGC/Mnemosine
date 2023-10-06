package Service

import (
	"card-service/Model/card"
	pb "card-service/Server"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Service interface {
	SayHello(ctx context.Context, name string) string
	CreateUserCard(ctx context.Context, in *pb.CreateFlashCardReq) (*pb.CreateFlashCardRes, error)
	EditCard(ctx context.Context, in *pb.EditCardReq) (*pb.EditCardRes, error)
	CreateDeck(ctx context.Context, in *pb.CreateDeckReq) (*pb.CreateDeckRes, error)
	EditDeck(ctx context.Context, in *pb.EditDeckReq) (*pb.EditDeckReq, error)
	GetCardsByUser(ctx context.Context, in *pb.GetUserCardsByUsernameReq) (*pb.GetUserCardsByUsernameRes, error)
	GetCardsByDeck(ctx context.Context, in *pb.GetCardsByDeckIdReq) (*pb.GetCardsByDeckIdRes, error)
}

type CardService struct {
	context      context.Context
	databaseUrl  string
	databaseName string
	db           *mongo.Database
	client       *mongo.Client
}

func NewCardService(ctx context.Context, databaseUrl string, databaseName string) (Service, error) {
	client, err := getClient(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}
	db := client.Database(databaseName)
	return &CardService{
		context:      ctx,
		databaseUrl:  databaseUrl,
		databaseName: databaseName,
		db:           db,
		client:       client,
	}, nil
}

func getClient(ctx context.Context, URI string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal("unable to get client", err)
		return nil, err
	}
	return client, err
}

func (s *CardService) Stop(ctx context.Context) error {
	err := s.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardService) StopAndDrop(ctx context.Context) error {
	err := s.db.Drop(ctx)
	if err != nil {
		return err
	}
	err = s.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardService) CreateDeck(ctx context.Context, in *pb.CreateDeckReq) (*pb.CreateDeckRes, error) {
	return nil, nil
}

func (s *CardService) EditDeck(ctx context.Context, in *pb.EditDeckReq) (*pb.EditDeckReq, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) GetCardsByDeck(ctx context.Context, in *pb.GetCardsByDeckIdReq) (*pb.GetCardsByDeckIdRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) EditCard(ctx context.Context, in *pb.EditCardReq) (*pb.EditCardRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) GetCardsByUser(ctx context.Context, in *pb.GetUserCardsByUsernameReq) (*pb.GetUserCardsByUsernameRes, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CardService) SayHello(ctx context.Context, name string) string {
	return "Hello, " + name
}

func (s *CardService) CreateUserCard(ctx context.Context, in *pb.CreateFlashCardReq) (*pb.CreateFlashCardRes, error) {
	var answers []card.Answer

	for _, val := range in.Card.Answers {
		answers = append(answers, card.Answer{
			Field:            int(val.Field),
			Answers:          val.Answers,
			IncorrectAnswers: val.IncorrectAnswers,
			Explanation:      val.Explanation,
		})
	}

	insertCard := card.UserFlashCard{
		Username:    in.Card.Username,
		Title:       in.Card.Title,
		Text:        in.Card.Text,
		Answers:     answers,
		Media:       in.Card.Media,
		Lang:        in.Card.Lang,
		Topics:      in.Card.Topics,
		DateCreated: 0,
	}

	id, err := card.InsertOne(ctx, s.db, insertCard)
	if err != nil {
		return nil, err
	}
	return &pb.CreateFlashCardRes{
		Id: id,
	}, nil
}
