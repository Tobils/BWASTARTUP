package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

/**
1. generate token
2. validasi token
*/

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3t_k3y")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// tanda tangani token
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token !")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

}
