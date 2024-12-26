package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("ASDWASDQWERQASDACAQWE")

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
