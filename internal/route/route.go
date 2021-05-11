package route

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	auth "github.com/siampudan/learning/authsession/internal/auth/handler"
	authRepo "github.com/siampudan/learning/authsession/internal/auth/repo"
	authUsecase "github.com/siampudan/learning/authsession/internal/auth/usecase"
	"github.com/siampudan/learning/authsession/internal/middleware"
	product "github.com/siampudan/learning/authsession/internal/product/handler"
	prodRepo "github.com/siampudan/learning/authsession/internal/product/repo"
	prodUsecase "github.com/siampudan/learning/authsession/internal/product/usecase"
)

type Handler struct {
	Log   *zap.Logger
	DB    *pg.DB
	Cache *redis.Client
	R     *gin.Engine
	Ctx   context.Context
}

func NewHandler(ctx context.Context, DB *pg.DB, Log *zap.Logger, Cache *redis.Client) *Handler {
	return &Handler{
		DB:    DB,
		Log:   Log,
		Cache: Cache,
		Ctx:   ctx,
	}
}

func (routeHandler *Handler) SetupRoutes() {
	authRepo := authRepo.NewAuthRepository(routeHandler.DB, routeHandler.Log)
	productRepo := prodRepo.NewProductRepo(routeHandler.DB, routeHandler.Log)

	authUc := authUsecase.NewAuthUseCase(authRepo, routeHandler.Log, routeHandler.Cache)
	prodUc := prodUsecase.NewProductUseCase(productRepo, routeHandler.Log)

	route := routeHandler.R.Group("/api")
	middle := middleware.NewMiddleware(routeHandler.Ctx, routeHandler.Cache)

	auth.AuthRoute(authUc, route)
	product.ProductRoute(prodUc, route, middle)
}
