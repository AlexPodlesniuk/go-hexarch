package event

import (
	"context"
	"user-alerts/domain"
)

type UserCreated struct {
	ID   string
	Name string
}

type UserCreatedHandler struct {
	repo domain.UserRepository
}

func NewUserCreatedHandler(ctx context.Context, repo domain.UserRepository) UserCreatedHandler {
	return UserCreatedHandler{repo: repo}
}

func (h *UserCreatedHandler) Handle(ctx context.Context, e *UserCreated) error {

	user := domain.NewUser(e.ID, e.Name)
	_, err := h.repo.SaveUser(ctx, user)
	return err
}
