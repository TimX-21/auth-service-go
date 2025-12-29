package config

import (
	"os"
	"strconv"
	"time"
)

func LoadResetConfig() ResetConfig {
	otpTTLMinutes, _ := strconv.Atoi(os.Getenv("RESET_OTP_TTL_MINUTES"))
	resetTokenTTLMinutes, _ := strconv.Atoi(os.Getenv("RESET_TOKEN_TTL_MINUTES"))
	maxAttempts, _ := strconv.Atoi(os.Getenv("RESET_OTP_MAX_ATTEMPTS"))
	resendCooldownSeconds, _ := strconv.Atoi(os.Getenv("RESET_OTP_RESEND_COOLDOWN_SECONDS"))
	resendLimitPerHour, _ := strconv.Atoi(os.Getenv("RESET_OTP_RESEND_LIMIT_PER_HOUR"))

	return ResetConfig{
		OTPTTL:             time.Minute * time.Duration(otpTTLMinutes),
		ResetTokenTTL:      time.Minute * time.Duration(resetTokenTTLMinutes),
		MaxAttempts:        maxAttempts,
		ResendCooldown:     time.Second * time.Duration(resendCooldownSeconds),
		ResendLimitPerHour: resendLimitPerHour,
		HMACSecret:         os.Getenv("RESET_OTP_HMAC_SECRET"),
	}
}

