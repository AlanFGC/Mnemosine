package Service

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) SayHello(ctx context.Context, name string) (msg string) {

	defer func(start time.Time) {
		fmt.Printf("MSG: %v took=%v", msg, time.Since(start))
	}(time.Now())

	return s.next.SayHello(ctx, name)
}
