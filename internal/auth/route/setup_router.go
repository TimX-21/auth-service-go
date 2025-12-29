package route

import (
	"github.com/TimX-21/auth-service-go/internal/auth/handler"
	"github.com/TimX-21/auth-service-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	AuthHandler *handler.AuthHandler
}

func NewRouteConfig(
	authH *handler.AuthHandler,
) *RouteConfig {
	return &RouteConfig{
		AuthHandler: authH,
	}
}

func Setup(c *RouteConfig) *gin.Engine {
	s := gin.New()
	s.ContextWithFallback = true

	s.Use(gin.Recovery())
	s.Use(middleware.ErrorCustomRecoveryMiddleware())
	s.Use(middleware.LoggerMiddleware())
	s.Use(middleware.GeneralErrorMiddleware())

	api := s.Group("/api/v1")
	SetupAuthRoutes(api, c)
	return s
}

func SetupAuthRoutes(s *gin.RouterGroup, c *RouteConfig) {
	auth := s.Group("/auth")
	auth.GET("/", c.AuthHandler.GetUserDataHandler)
	auth.POST("/login", c.AuthHandler.LoginHandler)
	auth.POST("/register", c.AuthHandler.RegisterHandler)
    
	auth.POST("/forgot-password", c.AuthHandler.ForgotPasswordRequestHandler)
	auth.POST("/verify-reset-password", c.AuthHandler.VerifyResetOTP)
	auth.POST("/reset-password", c.AuthHandler.ResetPassword)
}
