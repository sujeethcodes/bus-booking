package entity

type User struct {
	UserID      string
	Name        string
	Email       string
	Password    string
	Token       string
	PhoneNumber int
	Address     interface{}
	Status      string
}
