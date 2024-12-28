package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("ASDWASDQWERQASDACAQWE")

func NewService() Service {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// signing token
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return SECRET_KEY, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}