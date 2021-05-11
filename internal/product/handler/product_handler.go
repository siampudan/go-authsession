package handler

import (
	"net/http"

	"github.com/siampudan/learning/authsession/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/siampudan/learning/authsession/internal/apperror"
	"github.com/siampudan/learning/authsession/internal/product"
)

func ProductRoute(usecase product.ProductUseCase, r *gin.RouterGroup, middle middleware.Middleware) {
	uc := &Handler{
		usecase: usecase,
	}

	r.GET("/products", middle.AuthCustomer(), uc.GetProducts)
}

type Handler struct {
	usecase product.ProductUseCase
}

type Response struct {
	Data    interface{}
	Status  string
	Message string
}

func (prodHandler *Handler) GetProducts(c *gin.Context) {
	result, err := prodHandler.usecase.GetProducts(c)
	if err != nil {
		apperror.Response(c, err)
		return
	}

	c.JSON(http.StatusOK, &Response{
		Data:    result,
		Status:  "Status Success",
		Message: "Berhasil mengambil data product",
	})

}
