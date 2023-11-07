package command

import (
	"context"
	"encoding/json"
	"log"
	"user-acquisition/domain"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type CreateUser struct {
	ID   string
	Name string
}

type CreateUserHandler struct {
	repo domain.UserRepository
	pub  message.Publisher
}

func NewCreateUserHandler(repo domain.UserRepository, pub message.Publisher) CreateUserHandler {
	return CreateUserHandler{repo, pub}
}

func (h *CreateUserHandler) Handle(ctx context.Context, c *CreateUser) (*domain.User, error) {
	user := domain.NewUser(c.ID, c.Name)
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte(userJSON))
	h.pub.Publish("users", msg)

	return h.repo.SaveUser(ctx, user)
}
