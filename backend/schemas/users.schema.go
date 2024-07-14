package schemas

import (
	"time"
)

type CreateUser struct {
	UserName  string `json:"user_name" binding:"required,alphanum"`
	UserEmail string `json:"user_email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UpdateUser struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Password  string `json:"password"`
}

type CreateUserResponse struct {
	UserName          string    `json:"user_name"`
	UserEmail         string    `json:"user_email"`
	CreatedAt         time.Time `json:"updated_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

type UpdateUserResponse struct {
	UserName          string    `json:"user_name"`
	UserEmail         string    `json:"user_email"`
	UpdatedAt         time.Time `json:"updated_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string             `json:"access_token"`
	User        CreateUserResponse `json:"user"`
}
