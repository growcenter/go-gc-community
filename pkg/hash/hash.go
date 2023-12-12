package hash

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

/*func Salt() ([]byte, error) {
	salt := make([]byte, bcrypt.DefaultCost)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}*/

func Generate(password string) (string, error) {
	salt := make([]byte, bcrypt.DefaultCost)
	
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	salted := append([]byte(password), salt...)
	
	hash, err := bcrypt.GenerateFromPassword(salted, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Validate(hashed string, input []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), input)
	return err
}