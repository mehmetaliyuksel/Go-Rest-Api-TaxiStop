package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var jwtKey = []byte("taxistopsecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string, id uint) (tokenString string, err error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
func ValidateToken(signedToken string) (err error) {
	signedToken = strings.Split(signedToken, "Bearer ")[1]

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
