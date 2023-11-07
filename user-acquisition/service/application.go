package service

import (
	"context"
	"user-acquisition/adapters"
	"user-acquisition/app"
	"user-acquisition/app/command"
	"user-acquisition/app/query"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

func NewApp(ctx context.Context) app.Application {
	userRepo := adapters.NewRepo()
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)

	if err != nil {
		panic(err)
	}

	return app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(userRepo, publisher),
		},
		Queries: app.Queries{
			FindUser: query.NewFindUserHandler(userRepo),
		},
	}
}
