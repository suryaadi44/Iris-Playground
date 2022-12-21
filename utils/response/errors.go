package response

import "errors"

var (
	ErrDuplicateEmail         = errors.New("email already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
