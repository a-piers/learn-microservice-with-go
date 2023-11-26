package models

import (
	"time"
)

type UserModel struct {
	Id          string    `json:"id" gorm:"primaryKey;type:varchar(25)"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	BirthDate   time.Time `json:"birth_date"`
}

// USER LOGIN ARGS & RESULT
type UserLoginArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResult struct {
	Id                  string            `json:"id"`
	UserInfos           map[string]string `json:"user_infos"`
	AuthenticationToken string            `json:"authentication_token"`
	Result              Result            `json:"result"`
}

// -------------------------
