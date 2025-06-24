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
	UserType    string          `json:"user_type"`
}

type EditUserReq struct {
	UserID      string          `json:"user_id"`
	Name        string          `json:"name,omitempty"`
	Email       string          `json:"email,omitempty"`
	Password    string          `json:"password,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Address     json.RawMessage `json:"address,omitempty"`
	Status      string          `json:"status,omitempty"`
}

type DeleteUserReq struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
}
