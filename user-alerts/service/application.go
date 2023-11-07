package service

import (
	"context"
	"user-alerts/adapters"
	"user-alerts/app"
	"user-alerts/app/event"
	"user-alerts/app/query"
)

func NewApp(ctx context.Context) app.Application {
	userRepo := adapters.NewRepo()

	return app.Application{
		Queries: app.Queries{
			FindUser: query.NewFindUserHandler(userRepo),
		},
		Events: app.Events{
			UserCreated: event.NewUserCreatedHandler(ctx, userRepo),
		},
	}
}
