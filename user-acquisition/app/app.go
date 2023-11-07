package app

import (
	"user-acquisition/app/command"
	"user-acquisition/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser command.CreateUserHandler
}

type Queries struct {
	FindUser query.FindUserHandler
}
