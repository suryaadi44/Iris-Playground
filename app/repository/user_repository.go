package repository

import (
	"context"

	"suryaadi44/iris-playground/app/entity"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
