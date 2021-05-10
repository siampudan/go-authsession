package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	auth "github.com/siampudan/learning/authsession/internal/auth/handler"
	authRepo "github.com/siampudan/learning/authsession/internal/auth/repo"
	authUsecase "github.com/siampudan/learning/authsession/internal/auth/usecase"
)

type Handler struct {
	Log   *zap.Logger
	DB    *pg.DB
	Cache *redis.Client
	R     *gin.Engine
}

func NewHandler(DB *pg.DB, Log *zap.Logger, Cache *redis.Client) *Handler {
	return &Handler{
		DB:    DB,
		Log:   Log,
		Cache: Cache,
	}
}

func (routeHandler *Handler) SetupRoutes() {
	authRepo := authRepo.NewAuthRepository(routeHandler.DB, routeHandler.Log)
	authUc := authUsecase.NewAuthUseCase(authRepo, routeHandler.Log, routeHandler.Cache)

	route := routeHandler.R.Group("/api")
	auth.AuthRoute(authUc, route)
}
