package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var SecretKey = "SECRET_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenSring string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenSring)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims);
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")

}

func GetUserInfoFromToken(tokenString string) (name string, email string, userId uint,  err error) {
	claims, err := DecodeToken(tokenString)

	if err != nil {
		return "", "", 0, err
	}

	name, _ = claims["name"].(string)
	email, _ = claims["email"].(string)
	uidFloat, _ := claims["userId"].(float64)

	if name == "" || email == "" {
		return "", "", 0, fmt.Errorf("invalid token")
	}

	userId = uint(uidFloat)

	return name, email, userId, nil
}