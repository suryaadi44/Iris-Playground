package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Hasher interface {
	GenerateFromPassword(password string, hashParams Params) (encodeHash string, err error)
	CompareHashAndPassword(hashedPassword, password string) (bool, error)
}

type Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	// ErrInvalidHash is returned when the encoded hash is not in the correct format.
	ErrInvalidHash = errors.New("the encoded hash is not in the correct format")

	// ErrIncompatibleVersion is returned when the version of argon2 is not supported.
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

var DefaultParams = Params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

type Argon2 struct{}

func NewArgon2Hasher() Hasher {
	return &Argon2{}
}

func (Argon2) GenerateFromPassword(password string, hashParams Params) (encodeHash string, err error) {
	salt, err := generateRandomBytes(hashParams.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, hashParams.iterations, hashParams.memory, hashParams.parallelism, hashParams.keyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodeHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, hashParams.memory, hashParams.iterations, hashParams.parallelism, b64Salt, b64Hash)

	return encodeHash, nil
}

func (Argon2) CompareHashAndPassword(hashedPassword, password string) (bool, error) {
	param, savedSalt, savedHash, err := decodeHash(hashedPassword)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(password), savedSalt, param.iterations, param.memory, param.parallelism, param.keyLength)
	if subtle.ConstantTimeCompare(hash, savedHash) == 1 {
		return true, nil
	}

	return false, nil
}

func generateRandomBytes(n uint32) (b []byte, err error) {
	b = make([]byte, n)
	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func decodeHash(encodedHash string) (p *Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
