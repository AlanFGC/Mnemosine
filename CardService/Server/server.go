package Server

import (
	"card-service/Model"
	"card-service/Service"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	service Service.Service
}

func NewGrpcServer(s Service.Service) *GrpcServer {
	return &GrpcServer{service: s}
}

func (g GrpcServer) SayHello(ctx context.Context, name *Name) (*Hello, error) {
	greet := g.service.SayHello(ctx, name.GetName())
	res := Hello{Hello: greet}
	return &res, nil
}

func (g GrpcServer) CreateUserFlashCard(ctx context.Context, in *CreateFlashCardReq) (*CreateFlashCardRes, error) {
	var answers []Model.Answer

	for _, val := range in.Card.Answers {
		answers = append(answers, Model.Answer{
			Field:            int(val.Field),
			Answers:          val.Answers,
			IncorrectAnswers: val.IncorrectAnswers,
			Explanation:      val.Explanation,
		})
	}

	insertCard := Model.UserFlashCard{
		Username:    in.Card.Username,
		Title:       in.Card.Title,
		Text:        in.Card.Text,
		Answers:     answers,
		Media:       in.Card.Media,
		Lang:        in.Card.Lang,
		Topics:      in.Card.Topics,
		DateCreated: 0,
	}

	id, err := g.service.CreateUserCard(ctx, insertCard)
	if err != nil {
		return nil, err
	}

	return &CreateFlashCardRes{
		Id: id,
	}, nil
}

func (g GrpcServer) CreateDeck(ctx context.Context, in *CreateDeckReq) (*CreateDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) EditCard(ctx context.Context, in *EditCardReq) (*EditCardRes, error) {
	id := in.CardId

	answers := protoAnswerToModelAnswer(in.Card.Answers)
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "No id provided")
	}

	primitiveId, err := primitive.ObjectIDFromHex(in.CardId)
	if err != nil {
		return nil, err
	}

	card := Model.UserFlashCard{
		ID:          primitiveId,
		Username:    in.Card.Username,
		Title:       in.Card.Title,
		Text:        in.Card.Text,
		Answers:     answers,
		Media:       in.Card.Media,
		Lang:        in.Card.Lang,
		Topics:      in.Card.Topics,
		DateCreated: 0,
	}

	err = g.service.EditCard(ctx, card)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (g GrpcServer) EditDeck(ctx context.Context, in *EditDeckReq) (*EditDeckRes, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) GetCardsByUsername(ctx context.Context, in *GetUserCardsByUsernameReq) (*GetUserCardsByUsernameRes, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) GetCardsByDeckId(ctx context.Context, in *GetCardsByDeckIdReq) (*GetCardsByDeckIdRes, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcServer) mustEmbedUnimplementedCardServiceServer() {
	//TODO implement me
	panic("implement me")
}

func protoAnswerToModelAnswer(in []*Answer) (out []Model.Answer) {
	var answers []Model.Answer

	for _, val := range in {
		var qType Model.QuestionType

		if val.QuestionType == 0 {
			qType = Model.Open
		} else if val.QuestionType == 1 {
			qType = Model.SingleAnswer
		} else if val.QuestionType == 2 {
			qType = Model.MultipleChoice
		} else {
			qType = Model.Undefined
		}

		answers = append(answers, Model.Answer{
			Field:            int(val.Field),
			Answers:          val.Answers,
			IncorrectAnswers: val.IncorrectAnswers,
			Explanation:      val.Explanation,
			QuestionType:     qType,
		})
	}

	return answers
}
