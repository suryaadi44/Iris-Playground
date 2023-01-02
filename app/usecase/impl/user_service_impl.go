package impl

import (
	"context"
	"log"

	"github.com/suryaadi44/iris-playground/app/dto"
	"github.com/suryaadi44/iris-playground/app/repository"
	"github.com/suryaadi44/iris-playground/app/usecase"
	"github.com/suryaadi44/iris-playground/utils/password"
	"github.com/suryaadi44/iris-playground/utils/response"
)

type UserServiceImpl struct {
	ur repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) usecase.UserService {
	return &UserServiceImpl{
		ur: userRepository,
	}
}

func (s *UserServiceImpl) SignUp(ctx context.Context, user *dto.UserSignUpRequest) error {
	hashedPassword, err := password.GeneratePassword(user.Password)
	if err != nil {
		log.Println("[User] Failed to hash password: ", err)
		return err
	}

	user.Password = hashedPassword
	userEntity := user.ToEntity()

	err = s.ur.AddUser(ctx, userEntity)
	if err != nil {
		log.Println("[User] Failed to add user: ", err)
		return err
	}

	return nil
}

func (s *UserServiceImpl) LogIn(ctx context.Context, user *dto.UserLoginRequest) (*dto.LoginResponse, error) {
	userEntity, err := s.ur.FindByEmail(ctx, user.Email)
	if err != nil {
		if err == response.ErrUserNotFound {
			return nil, response.ErrInvalidEmailOrPassword
		}

		log.Println("[User] Failed to find user: ", err)
		return nil, err
	}

	match, err := password.VerifyPassword(user.Password, userEntity.Password)
	if err != nil {
		log.Println("[User] Failed to compare password: ", err)
		return nil, err
	}

	if !match {
		return nil, response.ErrInvalidEmailOrPassword
	}

	// TODO: Generate JWT token
	refreshtoken := "refreshtoken"
	accessToken := "accesstoken"

	return dto.NewLoginResponse(userEntity, refreshtoken, accessToken), nil
}
