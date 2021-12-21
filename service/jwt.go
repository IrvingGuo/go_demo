package service

import (
	"resource-plan-improvement/config"
	"resource-plan-improvement/util"
	"time"
)

var (
	jwt = config.Conf.Jwt
)

func GenerateToken(userId uint) (signedToken string, err error) {
	return util.GenerateToken(userId, jwt.Secret, jwt.Duration)
}

func ValidateToken(token string) (uint, error) {
	return util.VerifyToken(token, jwt.Secret, time.Now().Unix())
}

func GetUserIdFromToken(token string) (uint, error) {
	return util.GetUserIdFromToken(token, jwt.Secret)
}
