package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/siampudan/learning/authsession/internal/apperror"
)

func NewMiddleware(ctx context.Context, redis *redis.Client) Middleware {
	return Middleware{
		ctx:   ctx,
		redis: redis,
	}
}

type Middleware struct {
	ctx   context.Context
	redis *redis.Client
}

func (middle *Middleware) AuthCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("my-secret-cookie")
		if err != nil {
			apperror.Response(c, apperror.New(http.StatusUnauthorized, "You Are Not Authorized"))
			return
		}
		exists, _ := middle.redis.Exists(middle.ctx, cookie).Result()
		if exists == 0 {
			apperror.Response(c, apperror.New(http.StatusUnauthorized, "You Are Not Authorized"))
			return
		}
		c.Next()
	}
}
