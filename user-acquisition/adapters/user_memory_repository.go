package adapters

import (
	"context"
	"user-acquisition/domain"
)

type UserMemoryReoo struct {
	users map[string]*domain.User
}

func NewRepo() *UserMemoryReoo {
	return &UserMemoryReoo{users: make(map[string]*domain.User)}
}

func (r *UserMemoryReoo) FindUserById(ctx context.Context, userID string) *domain.User {
	return r.users[userID]
}

func (r *UserMemoryReoo) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	r.users[user.ID] = user

	return user, nil
}
