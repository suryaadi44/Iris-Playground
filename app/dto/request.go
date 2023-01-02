package dto

import (
	"github.com/suryaadi44/iris-playground/app/api/grpc/pb"
	"github.com/suryaadi44/iris-playground/app/entity"
)

type UserSignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *UserSignUpRequest) ToEntity() *entity.User {
	return &entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func NewUserLoginRequest(pb *pb.LogInRequest) *UserLoginRequest {
	return &UserLoginRequest{
		Email:    pb.Email,
		Password: pb.Password,
	}
}
