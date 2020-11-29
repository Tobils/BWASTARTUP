package auth

import "github.com/dgrijalva/jwt-go"

/**
1. generate token
2. validasi token
*/

type Service interface {
	GenerateToken(userID int) (string, error)
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
