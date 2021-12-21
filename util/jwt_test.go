package util_test

import (
	"fmt"
	"resource-plan-improvement/util"
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	secret := "1b2n3n4"
	userId := 2
	duration := 24
	token, _ := util.GenerateToken(uint(userId), secret, duration)
	fmt.Println(token)
	var testUserId1 uint
	var err error
	if testUserId1, err = util.VerifyToken(token, secret, time.Now().Unix()); err != nil {
		t.Fatalf("token operation error, %s", err.Error())
	} else if testUserId1 != uint(userId) {
		t.Fatalf("user id parsing error")
	}
	var testUserId2 uint
	if testUserId2, err = util.GetUserIdFromToken(token, secret); err != nil {
		t.Fatalf("token operation error, %s", err.Error())
	}
	if testUserId2 != uint(userId) {
		t.Fatalf("user id parsing error")
	}
}
