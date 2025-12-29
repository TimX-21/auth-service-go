package handler

import (
	"net/http"

	"github.com/TimX-21/auth-service-go/internal/apperror"
	"github.com/TimX-21/auth-service-go/internal/auth/dto"
	"github.com/TimX-21/auth-service-go/internal/auth/model"
	"github.com/TimX-21/auth-service-go/internal/auth/service"
	"github.com/TimX-21/auth-service-go/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthServiceItf
}

func NewAuthHandler(s service.AuthServiceItf) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) GetUserDataHandler(c *gin.Context) {

	request := dto.GetUserDataRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}
	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	user := model.User{Email: request.Email}
	userData, err := h.authService.GetUserDataService(c, user)
	if err != nil {
		c.Error(err)
		return
	}

	response := dto.GetUserDataResponse{
		ID:         userData.ID,
		Email:      userData.Email,
		IsVerified: userData.IsVerified,
		CreatedAt:  userData.CreatedAt,
		UpdatedAt:  userData.UpdatedAt,
	}

	util.HandleResponse(response, http.StatusOK, c)
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	request := dto.LoginRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	RequestedUser := model.User{
		Email:    request.Email,
		Password: request.Password,
	}

	loginResponse, err := h.authService.LoginService(c, RequestedUser)
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.LoginResponse{
		AccessToken: loginResponse,
	}

	util.HandleResponse(res, http.StatusOK, c)
}

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	request := dto.RegisterRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	NewUser := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err := h.authService.RegisterService(c, NewUser)
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.RegisterResponse{
		Message: "Registration successful. Please log in to continue.",
	}

	util.HandleResponse(res, http.StatusCreated, c)
}

func (h *AuthHandler) ForgotPasswordRequestHandler(c *gin.Context) {
	request := dto.ForgotPasswordRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	RequestedUser := model.User{
		Email: request.Email,
	}

	_ = h.authService.ForgotPasswordRequestService(c, RequestedUser)

	res := dto.ForgotPasswordResponse{
		Message: "If the email exists, an OTP has been sent.",
	}

	util.HandleResponse(res, http.StatusOK, c)
}

func (h *AuthHandler) VerifyResetOTP(c *gin.Context) {
	request := dto.VerifyResetOTPRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	resetToken, err := h.authService.VerifyResetOTPService(c, model.User{Email: request.Email}, model.PasswordResetOTP{OTP: request.OTP})
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.VerifyResetOTPResponse{
		ResetToken: resetToken,
	}

	util.HandleResponse(res, http.StatusOK, c)
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	request := dto.ResetPasswordRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if err := util.ValidateStruct(request); err != nil {
		c.Error(apperror.ErrInvalidRequest)
		return
	}

	if request.NewPassword != request.ConfirmPassword {
		c.Error(apperror.ErrPasswordNotMatch)
		return
	}

	err := h.authService.ResetPasswordService(c, model.PasswordResetToken{Token: request.ResetToken}, model.User{Password: request.NewPassword})
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.ResetPasswordResponse{
		Message: "Your password has been reset successfully.",
	}

	util.HandleResponse(res, http.StatusOK, c)
}
