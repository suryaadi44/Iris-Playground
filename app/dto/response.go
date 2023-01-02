package dto

import (
	"github.com/suryaadi44/iris-playground/app/api/grpc/pb"
	"github.com/suryaadi44/iris-playground/app/entity"
)

type LoginResponse struct {
	UID          string `json:"uid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Permission   int    `json:"permission"`
}

func NewLoginResponse(user *entity.User, accessToken, refreshToken string) *LoginResponse {
	return &LoginResponse{
		UID:          user.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Permission:   int(user.Permission),
	}
}

func (l *LoginResponse) ToProto() *pb.LogInResponse {
	return &pb.LogInResponse{
		Uid:          l.UID,
		AccessToken:  l.AccessToken,
		RefreshToken: l.RefreshToken,
		Permission:   int64(l.Permission),
	}
}
