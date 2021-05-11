package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/siampudan/learning/authsession/internal/product"
	"go.uber.org/zap"
)

func NewProductUseCase(repo product.ProductRepository, log *zap.Logger) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
		log:  log,
	}
}

type ProductUseCase struct {
	repo product.ProductRepository
	log  *zap.Logger
}

func (prodUc *ProductUseCase) GetProducts(c *gin.Context) ([]*product.Product, error) {
	result, err := prodUc.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}
