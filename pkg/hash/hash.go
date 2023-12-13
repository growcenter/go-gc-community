package hash

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func Salt() ([]byte, error) {
	salt := make([]byte, bcrypt.DefaultCost)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func Generate(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Validate(hashed string, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))	
}