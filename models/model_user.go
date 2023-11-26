package models

import (
	"time"
)

type UserBase struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	BirthDate   time.Time `json:"birth_date"`
}

type UserModel struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	UserBase
}

// USER LOGIN ARGS & RESULT
type UserLoginArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResult struct {
	AuthenticationToken string `json:"authentication_token"`
	Result              Result `json:"result"`
}

// -------------------------
