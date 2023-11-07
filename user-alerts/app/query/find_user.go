package query

import (
	"context"
	"errors"
	"user-alerts/domain"
)

type FindUser struct {
	ID string
}

type FindUserHandler struct {
	repo domain.UserRepository
}

func NewFindUserHandler(repo domain.UserRepository) FindUserHandler {
	return FindUserHandler{repo}
}

func (h *FindUserHandler) Handle(ctx context.Context, q *FindUser) (*domain.User, error) {
	user := h.repo.FindUserById(ctx, q.ID)

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
