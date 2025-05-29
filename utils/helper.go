package utils

import (
	"math/rand"
)

func GenerateUserID() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := ""
	for i := 0; i < 5; i++ {
		id += string(chars[rand.Intn(len(chars))])
	}
	return id
}
