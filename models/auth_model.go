package models


type RegisterRequest struct {
	Email    string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber int `json:"phone_number"`	
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}