package middleware

import (
	"errors"
	"net/http"

	"github.com/TimX-21/auth-service-go/internal/apperror"
	"github.com/TimX-21/auth-service-go/internal/dto"

	"github.com/gin-gonic/gin"
)

func GeneralErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]

			HandleSingleError(c, err)
		}
	}
}

func ErrorCustomRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				rString, ok := r.(string)

				var err error
				if ok {
					err = errors.New(rString)
				} else {
					err = apperror.ErrInternalServer
				}

				c.Error(err)
				e := &dto.Error{
					Message: err.Error(),
				}
				c.JSON(http.StatusInternalServerError, dto.Response[any]{
					Success: false,
					Error:   e,
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}

func HandleSingleError(c *gin.Context, err *gin.Error) {
	metaMap, ok := err.Meta.(map[string]int)
	if !ok {
		// Fallback: if error doesn't have Meta, treat as internal server error
		e := &dto.Error{
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dto.Response[any]{
			Success: false,
			Error:   e,
		})
		return
	}

	statusCode := metaMap["status_code"]

	e := &dto.Error{
		Message: err.Error(),
	}
	c.JSON(statusCode, dto.Response[any]{
		Success: false,
		Error:   e,
	})
}
