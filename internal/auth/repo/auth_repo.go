package repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/siampudan/learning/authsession/internal/auth"
	"go.uber.org/zap"
)

func NewAuthRepository(DB *pg.DB, log *zap.Logger) *Auth {
	return &Auth{
		DB:  DB,
		Log: log,
	}
}

type Auth struct {
	DB  *pg.DB
	Log *zap.Logger
}

func (authRepo *Auth) CreateUser(newUser *auth.User) (*auth.User, error) {
	_, err := authRepo.DB.Model(newUser).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (authRepo *Auth) GetUserByID(ID int) (*auth.User, error) {
	var result auth.User
	err := authRepo.DB.Model(&result).Where("id = ?", ID).Select()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (authRepo *Auth) GetUserByUsername(userName string) (*auth.User, error) {
	var result auth.User
	err := authRepo.DB.Model(&result).Where("name = ?", userName).Select()
	if err != nil {
		return nil, err
	}
	return &result, nil
}
