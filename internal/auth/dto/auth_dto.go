package dto

import "time"

type GetUserDataRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUserDataResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
}