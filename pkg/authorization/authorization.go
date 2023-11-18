package authorization

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Authorization interface {
	Generate(accountNumber string, userId int) (string, error)
	Validate(token string) (*jwt.Token, error)
}

type Auth struct {
	secret string
	td	int
}

func NewAuthorization(secret string, td int) (*Auth, error) {
	if secret == "" {
		return nil, errors.New("empty signing key")
	}

	if td == 0 {
		return nil, errors.New("empty duration")
	}

	return &Auth{secret: secret, td: td}, nil
}

/*func (a *Auth) Generate(id string, td time.Duration) (string, error) {
	expiry := time.Now().Add(td).Unix()
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiry,
		Subject: id,
	})	

	return token.SignedString([]byte(a.secret))
}*/

func (a *Auth) Generate(accountNumber string, userId int) (string, error) {
	expiry := time.Now().Add(time.Hour * time.Duration(a.td)).Unix()
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
		"accountNumber": accountNumber,
		"createdAt": time.Now().Unix(),
		"expiredAt": expiry,
	})

	return token.SignedString([]byte(a.secret))
}

func (a *Auth) Validate(encoded string) (*jwt.Token, error) {
	token, err := jwt.Parse(encoded, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(a.secret), nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	/*claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claim")
	}

	return claims["sub"].(string), nil*/

	return token, nil
}