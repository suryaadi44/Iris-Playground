package usecase

import (
	"context"

	"suryaadi44/iris-playground/app/dto"
)

type UserService interface {
	SignUp(ctx context.Context, user *dto.UserSignUpRequest) error
	LogIn(ctx context.Context, user *dto.UserLoginRequest) error
}
