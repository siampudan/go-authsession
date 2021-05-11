package repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/siampudan/learning/authsession/internal/product"
	"go.uber.org/zap"
)

func NewProductRepo(DB *pg.DB, log *zap.Logger) *ProductRepository {
	return &ProductRepository{
		DB:  DB,
		Log: log,
	}
}

type ProductRepository struct {
	DB  *pg.DB
	Log *zap.Logger
}

func (prodRepo *ProductRepository) GetProducts() ([]*product.Product, error) {
	var result []*product.Product
	err := prodRepo.DB.Model(&result).Select()
	if err != nil {
		return nil, err
	}
	return result, nil
}
