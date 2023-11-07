package app

import (
	"user-alerts/app/event"
	"user-alerts/app/query"
)

type Application struct {
	Events  Events
	Queries Queries
}

type Events struct {
	UserCreated event.UserCreatedHandler
}

type Queries struct {
	FindUser query.FindUserHandler
}
