package dto

import "time"

type GetUserDataRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUserDataResponse struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

type VerifyResetOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6,numeric"`
}

type VerifyResetOTPResponse struct {
	ResetToken string `json:"reset_token"`
}

type ResetPasswordRequest struct {
	ResetToken      string `json:"reset_token" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}
