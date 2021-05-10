package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	tableName struct{} `pg:"users"`
	ID        int      `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Password  string   `json:"password" pg:"password"`
}

func (user User) UserValidation() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Password, validation.Required),
	)
}

type UserRepository interface {
	CreateUser(*User) (*User, error)
	GetUserByID(ID int) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type UserUseCase interface {
	SignUp(*gin.Context) error
	Login(*gin.Context) (*http.Cookie, error)
}
