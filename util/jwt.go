package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type customClaims struct {
	UserId uint `json:"userId"` // must be upper case, or parseToken() cannot get user id
	jwt.StandardClaims
}

func GenerateToken(userId uint, secret string, duration int) (signedToken string, err error) {
	customClaims := customClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(duration) * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(signedToken string, secret string, expectedExpiredAt int64) (userId uint, err error) {
	var customClaims *customClaims
	if customClaims, err = parseToken(signedToken, secret); err != nil {
		return
	}
	if customClaims.ExpiresAt < expectedExpiredAt {
		return 0, ErrExpiredToken
	}
	return customClaims.UserId, nil
}

func RefreshToken(signedToken string, secret string, duration int) (refreshedToken string, err error) {
	var claims *customClaims
	if claims, err = parseToken(signedToken, secret); err != nil {
		return
	}
	claims.ExpiresAt = time.Now().Add(time.Duration(duration) * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GetUserIdFromToken(signedToken string, secret string) (userId uint, err error) {
	var claims *customClaims
	if claims, err = parseToken(signedToken, secret); err != nil {
		return
	}
	return claims.UserId, nil
}

func parseToken(signedToken, secret string) (*customClaims, error) {
	var err error
	var token *jwt.Token
	claims := &customClaims{}
	if token, err = jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}); err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
