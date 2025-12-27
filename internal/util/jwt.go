package util

import (
	"strconv"
	"time"

	"github.com/TimX-21/auth-service-go/internal/auth/model"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	UserID int64  `json:"user_id"`
	RoleID int   `json:"role_id"`
}

func GenerateJWT(user model.User, isReset bool) (string, error) {
	jwtsecret, err := GetJWTSecret()
	if err != nil {
		return "", err
	}

	now := time.Now()

	expireTime := now.Add(24 * time.Hour)
	if isReset {
		expireTime = now.Add(15 * time.Minute)
	}

	sub := strconv.FormatInt(user.ID, 10)
	claims := CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ewallet",
			Subject:   sub,
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: expireTime,
			},
			Audience: jwt.ClaimStrings{"ewallet"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}