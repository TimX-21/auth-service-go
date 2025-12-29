package config

import "time"

type ResetConfig struct {
	OTPTTL             time.Duration
	ResetTokenTTL      time.Duration
	MaxAttempts        int
	ResendCooldown     time.Duration
	ResendLimitPerHour int
	HMACSecret         string
}


