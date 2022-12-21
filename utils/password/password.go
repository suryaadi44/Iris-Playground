package password

import (
	"github.com/alexedwards/argon2id"
)

type Hasher interface {
	GeneratePassword(password string) (hashedPassword string, err error)
	VerifyPassword(password string, hashedPassword string) (match bool, err error)
}

type Argon2id struct{}

func NewArgon2idHasher() Hasher {
	return &Argon2id{}
}

func (Argon2id) GeneratePassword(password string) (hashedPassword string, err error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func (Argon2id) VerifyPassword(password string, hashedPassword string) (match bool, err error) {
	return argon2id.ComparePasswordAndHash(password, hashedPassword)
}
