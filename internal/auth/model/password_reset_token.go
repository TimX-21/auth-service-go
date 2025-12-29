package model

import "time"

type PasswordResetToken struct {
	Id        int64
	UserId    int64
	Token string
	ExpiresAt time.Time
	UsedAt    *time.Time
	CreatedAt time.Time
}
