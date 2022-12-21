package impl

import (
	"context"
	"errors"
	"suryaadi44/iris-playground/app/entity"
	"suryaadi44/iris-playground/app/repository"
	"suryaadi44/iris-playground/utils/response"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) AddUser(ctx context.Context, user *entity.User) error {
	err := r.db.
		WithContext(ctx).
		Create(user).
		Error
	if err != nil {
		if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
			if pgError.Code == "23505" && pgError.ConstraintName == "idx_users_email" {
				return response.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}
