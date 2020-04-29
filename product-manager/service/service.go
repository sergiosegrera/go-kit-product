package service

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/go-kit-product/product/models"
)

type ProductManagerService interface {
	// TODO: Implement
	// GetProducts(ctx context.Context) ([]*models.Product, error)
	// GetProduct(ctx context.Context, id int64) (models.Product, error)
	PostProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, id int64) error
}

type Service struct {
	db *pg.DB
}

func NewService(d *pg.DB) *Service {
	return &Service{db: d}
}

func (s *Service) PostProduct(ctx context.Context, product models.Product) error {
	err := s.db.Insert(&product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	product := &models.Product{Id: id}
	err := s.db.Delete(product)

	return err
}
