package utils

import (
	"errors"
	"math/rand"

	"github.com/labstack/echo/v4"
)

func GenerateUserID() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	for i := 0; i < 5; i++ {
		id += string(chars[rand.Intn(len(chars))])
	}
	return id
}

func VerifyUserId(user_id string, c echo.Context) error {
	verifyUserid := c.Get("userID")
	if verifyUserid != user_id {
		return errors.New("invaild userid ")
	}
	return nil
}
