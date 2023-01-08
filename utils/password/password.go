package password

import (
	"github.com/alexedwards/argon2id"
)

func GeneratePassword(password string) (hashedPassword string, err error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func VerifyPassword(password string, hashedPassword string) (match bool, err error) {
	return argon2id.ComparePasswordAndHash(password, hashedPassword)
}
