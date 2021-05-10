package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siampudan/learning/authsession/internal/apperror"
	"github.com/siampudan/learning/authsession/internal/auth"
)

func AuthRoute(authUc auth.UserUseCase, r *gin.RouterGroup) {
	handlerFunc := &Handler{
		authUc: authUc,
	}

	r.POST("/signup", handlerFunc.SignUp)
	r.POST("/login", handlerFunc.Login)
}

type Handler struct {
	authUc auth.UserUseCase
}

func (authHandler *Handler) SignUp(c *gin.Context) {
	err := authHandler.authUc.SignUp(c)
	if err != nil {
		apperror.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "status success",
		"message": "berhasil menyimpan user",
	})
}

func (authHandler *Handler) Login(c *gin.Context) {
	cookie, err := authHandler.authUc.Login(c)
	if err != nil {
		apperror.Response(c, err)
		return
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"status":  "status success",
		"message": "berhasil login",
	})
}
