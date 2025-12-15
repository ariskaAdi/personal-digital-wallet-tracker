package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	Generate(userId int, email string) (string, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secret string) JWTService {
	return &jwtService{secret}
}

func (j *jwtService) Generate(userId int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id" : userId,
		"email" : email,
		"exp" : time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return  token.SignedString([]byte(j.secretKey))
}