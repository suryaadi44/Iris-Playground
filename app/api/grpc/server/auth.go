package server

import (
	"context"

	"github.com/suryaadi44/iris-playground/app/api/grpc/pb"
	"github.com/suryaadi44/iris-playground/app/dto"
	"github.com/suryaadi44/iris-playground/app/usecase"
	"github.com/suryaadi44/iris-playground/utils/response"
	"github.com/suryaadi44/iris-playground/utils/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthenticateServer
	us        usecase.UserService
	validator validator.Validator
}

func NewAuthServer(us usecase.UserService, validator validator.Validator) *AuthServer {
	return &AuthServer{
		us:        us,
		validator: validator,
	}
}

func (a *AuthServer) LogIn(ctx context.Context, reqProto *pb.LogInRequest) (*pb.LogInResponse, error) {
	req := dto.NewUserLoginRequest(reqProto)
	if errs := a.validator.ValidateJSON(req); errs != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "invalid request: %v", errs)
	}

	token, err := a.us.LogIn(ctx, req)
	if err != nil {
		switch err {
		case response.ErrInvalidEmailOrPassword:
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return token.ToProto(), nil
}
