package server

import (
	"github.com/gin-gonic/gin"
	"github.com/siampudan/learning/authsession/internal/database"
	"github.com/siampudan/learning/authsession/internal/route"
	"go.uber.org/zap"
)

func Run() error {
	postgres := database.GetConnection()
	defer postgres.Close()

	redis := database.GetCache()

	log, _ := zap.NewDevelopment()
	defer log.Sync()

	r := gin.Default()

	rh := route.Handler{
		DB:    postgres,
		Cache: redis,
		Log:   log,
		R:     r,
	}

	rh.SetupRoutes()

	return r.Run()
}
