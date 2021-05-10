package apperror

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type AppError struct {
	Status  int    `json:"-"`
	Message string `json:"message,omitempty"`
}

func New(status int, msg string) *AppError {
	return &AppError{
		Status:  status,
		Message: msg,
	}
}

func (e AppError) Error() string {
	return e.Message
}

func Response(c *gin.Context, err error) {
	switch err.(type) {
	case *AppError:
		e := err.(*AppError)
		if e.Message == "" {
			c.AbortWithStatus(e.Status)
		} else {
			c.AbortWithStatusJSON(e.Status, e)
		}
		return
	case validation.Errors:
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
