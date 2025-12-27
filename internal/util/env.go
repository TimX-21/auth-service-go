package util

import (
	"fmt"
	"os"

	"github.com/TimX-21/auth-service-go/internal/common"
)

func GetJWTSecret() (string, error) {
	s := os.Getenv(common.JWTSecret)
	if s == "" {
		return "", fmt.Errorf("JWT_SECRET not set")
	}

	return s, nil
}
