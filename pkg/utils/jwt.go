package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secertpassword")

// type StandardClaims struct {
// 	Audience  string `json:"aud"` // The audience for which the token is intended
// 	ExpiresAt int64  `json:"exp"` // Expiration time of the token
// 	ID        string `json:"jti"` // Unique identifier for the token
// 	IssuedAt  int64  `json:"iat"` // Time at which the token was issued
// 	Issuer    string `json:"iss"` // The issuer of the token
// 	NotBefore int64  `json:"nbf"` // Time before which the token is not valid
// 	Subject   string `json:"sub"` // The subject of the token
// }

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func GenerateTokenUsingClaims(userID string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    userID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		return "", errors.New("could not login")
	}

	return token, nil
}

func VerifyTokenUsingClaims(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", errors.New("unauthenticated")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
