package util

import (
	"context"
	"log"
)

type DummyEmailSender struct{}

func NewDummyEmailSender() *DummyEmailSender {
	return &DummyEmailSender{}
}

func (d *DummyEmailSender) SendResetOTP(ctx context.Context, email, otp string) error {
	log.Printf("[DUMMY EMAIL] Reset OTP for %s: %s\n", email, otp)
	return nil
}
