package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/kevinburke/go.uuid"
	"github.com/siampudan/learning/authsession/internal/apperror"
	"github.com/siampudan/learning/authsession/internal/auth"
	"go.uber.org/zap"
)

func NewAuthUseCase(authRepo auth.UserRepository, log *zap.Logger, cache *redis.Client) *AuthUsecase {
	return &AuthUsecase{
		repo:  authRepo,
		log:   log,
		cache: cache,
	}
}

type AuthUsecase struct {
	repo  auth.UserRepository
	log   *zap.Logger
	cache *redis.Client
}

func (authUc *AuthUsecase) SignUp(c *gin.Context) error {
	var result auth.User
	err := c.ShouldBindJSON(&result)
	if err != nil {
		authUc.log.Warn("error validate struct", zap.Error(err))
		return apperror.New(http.StatusBadRequest, "error validate")
	}

	err = result.UserValidation()
	if err != nil {
		authUc.log.Warn("error validate struct", zap.Error(err))
		return apperror.New(http.StatusBadRequest, "error validate")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(result.Password), bcrypt.DefaultCost)
	if err != nil {
		authUc.log.Warn("error generate hashed password", zap.Error(err))
		return apperror.New(http.StatusBadRequest, "error request")
	}
	result.Password = string(hashedPassword)
	_, err = authUc.repo.CreateUser(&result)
	if err != nil {
		authUc.log.Warn("error create row", zap.Error(err))
		return apperror.New(http.StatusBadRequest, "error sign up")
	}

	return nil
}

func (authUc *AuthUsecase) Login(c *gin.Context) (*http.Cookie, error) {
	var result auth.User
	err := c.ShouldBindJSON(&result)
	if err != nil {
		authUc.log.Warn("error validate struct", zap.Error(err))
		return nil, apperror.New(http.StatusBadRequest, "error validation")
	}

	compUser, err := authUc.repo.GetUserByUsername(result.Name)
	if err != nil {
		authUc.log.Warn("error get user", zap.Error(err))
		return nil, apperror.New(http.StatusNotFound, "can't found user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(compUser.Password), []byte(result.Password))
	if err != nil {
		authUc.log.Warn("error compare user password and request password", zap.Error(err))
		return nil, apperror.New(http.StatusUnauthorized, "login unsuccessfull")
	}

	token := uuid.NewV4().String()
	cookie := &http.Cookie{
		Name:    "my-secret-cookie",
		Value:   token,
		Expires: time.Now().Add(240 * time.Second),
	}

	err = authUc.cacheSessionToken(token, result.Name)
	if err != nil {
		authUc.log.Warn("error cache session token", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "login unsuccessfull")
	}

	return cookie, nil
}

func (authUc *AuthUsecase) cacheSessionToken(token, userName string) error {
	ctx := context.Background()
	duration, _ := time.ParseDuration("120s")
	status, _ := authUc.cache.SetEX(ctx, token, userName, duration).Result()
	if len(status) == 0 {
		return errors.New("failed to cache sesstion token")
	}
	return nil
}
