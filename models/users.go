package models

type UserModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber int    `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsVerified  bool 	`json:"is_verified"`
	ID          int    `json:"id"`
}
