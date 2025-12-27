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
	service service.AuthServiceItf
}

func NewAuthHandler(s service.AuthServiceItf) *AuthHandler {
	return &AuthHandler{service: s}
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
	userData, err := h.service.GetUserDataService(c, user)
	if err != nil {
		c.Error(err)
		return
	}

	response := dto.GetUserDataResponse{
		ID:        userData.ID,
		Email:     userData.Email,
		IsActive:  userData.IsActive,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
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

	loginResponse, err := h.service.LoginService(c, RequestedUser)
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

	err := h.service.RegisterService(c, NewUser)
	if err != nil {
		c.Error(err)
		return
	}

	res := dto.RegisterResponse{}

	util.HandleResponse(res, http.StatusCreated, c)
}
