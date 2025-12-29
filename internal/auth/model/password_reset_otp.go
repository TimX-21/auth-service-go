package model

import "time"

type PasswordResetOTP struct {
	Id           int64
	UserId       int64
	OTP          string
	ExpiresAt    time.Time
	UsedAt       *time.Time
	AttemptCount int
	CreatedAt    time.Time
}
