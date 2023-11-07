package domain

import "context"

type UserRepository interface {
	SaveUser(ctx context.Context, user *User) (*User, error)
	FindUserById(ctx context.Context, userID string) *User
}
