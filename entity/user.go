package entity

import "encoding/json"

type User struct {
	UserID      string          `json:"user_id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	Token       string          `json:"token"`
	PhoneNumber string          `json:"phone_number"`
	Address     json.RawMessage `json:"address"`
	Status      string          `json:"status"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}
