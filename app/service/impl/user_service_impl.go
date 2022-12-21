package impl

import (
	"context"
	"log"
	"suryaadi44/iris-playground/app/dto"
	"suryaadi44/iris-playground/app/repository"
	"suryaadi44/iris-playground/app/service"
	"suryaadi44/iris-playground/utils/password"
	"suryaadi44/iris-playground/utils/response"
)

type UserServiceImpl struct {
	ur repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) service.UserService {
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

func (s *UserServiceImpl) LogIn(ctx context.Context, user *dto.UserLoginRequest) error {
	userEntity, err := s.ur.FindByEmail(ctx, user.Email)
	if err != nil {
		if err == response.ErrUserNotFound {
			return response.ErrInvalidEmailOrPassword
		}

		log.Println("[User] Failed to find user: ", err)
		return err
	}

	match, err := password.VerifyPassword(user.Password, userEntity.Password)
	if err != nil {
		log.Println("[User] Failed to compare password: ", err)
		return err
	}

	if !match {
		return response.ErrInvalidEmailOrPassword
	}

	return nil
}
