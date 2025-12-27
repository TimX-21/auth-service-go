package service

import (
	"context"

	"github.com/TimX-21/auth-service-go/internal/apperror"
	"github.com/TimX-21/auth-service-go/internal/auth/model"
	"github.com/TimX-21/auth-service-go/internal/auth/repository"
	"github.com/TimX-21/auth-service-go/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository repository.AuthRepositoryItf
	txManager      repository.TransactionManager
}

func NewAuthService(authRepository repository.AuthRepositoryItf, txManager repository.TransactionManager) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		txManager:      txManager,
	}
}

func (s *AuthService) GetUserDataService(ctx context.Context, user model.User) (*model.User, error) {
	userData, err := s.authRepository.GetUserByEmail(ctx, user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *AuthService) LoginService(ctx context.Context, user model.User) (string, error) {

	DbUserData, err := s.authRepository.GetUserByEmail(ctx, user)
	if err != nil {
		return "", apperror.ErrUserNotFound
	}

	InputPassword := user.Password
	DbPassword := DbUserData.Password

	err = bcrypt.CompareHashAndPassword([]byte(DbPassword), []byte(InputPassword))
	if err != nil {
		return "", apperror.ErrUnauthorized
	}

	// jwtsecret, err := util.GetJWTSecret()
	// if err != nil {
	// 	return "", apperror.ErrInternalServer
	// }

	token, err := util.GenerateJWT(user, false)
	if err != nil {
		return "", apperror.ErrInternalServer
	}

	return token, nil
}
