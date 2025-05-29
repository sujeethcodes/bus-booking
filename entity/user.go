package entity

type User struct {
	UserID      string
	Name        string
	Email       string
	Password    string
	PhoneNumber int
	Address     Address
}

type Address struct {
	Street   string
	City     string
	PostCode int
	State    string
	Country  string
}
