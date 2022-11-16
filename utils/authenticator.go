package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTSecret define jwt secret key
const JWTSecret string = "secret-key"

// JWTClaims jwt claims struct
type JWTClaims struct {
	jwt.StandardClaims
	Authorized bool   `json:"authorized"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

// GenerateJWT generate jwt token
func GenerateJWT(email string, name string) (string, error) {
	c := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 300).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "auth-app",
		},
		Authorized: true,
		Name:       name,
		Email:      email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(JWTSecret))
}

// ParseToken parse jwt token
func ParseToken(bearerHeader string) (interface{}, error) {
	fmt.Println("Haloo", bearerHeader)

	parts := strings.SplitN(bearerHeader, " ", 2)

	if len(parts) < 2 {
		return nil, errors.New("invalid header")
	}

	prefix := parts[0]
	tokenString := parts[1]

	if prefix != "Bearer" {
		return nil, errors.New("invalid header")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ParseAppID parse app id
func ParseAppID(appIDHeader string) (interface{}, error) {
	fmt.Println("Haloo 2", appIDHeader)
	switch appIDHeader {
	case "e-core":
		fmt.Println("Masuuk")
		return true, nil
	default:
		return nil, errors.New("Invalid App ID")
	}
}
