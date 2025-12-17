package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	Generate(userId int, email string) (string, error)
	Verify(tokenString string) (*jwt.Token, error)
	Decode(tokenSring string) (jwt.MapClaims, error)
}

type jwtService struct {
	secretKey string
}

func NewJWTService(secret string) JWTService {
	if secret == "" {
		panic("secret is required")
	}
	return &jwtService{secret}
}

func (j *jwtService) Generate(userId int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id" : int(userId),
		"email" : email,
		"exp" : time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return  token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) Verify(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}


func (j *jwtService) Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := j.Verify(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}