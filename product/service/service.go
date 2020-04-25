package service

import (
	"context"
	"errors"

	"github.com/sergiosegrera/store/product/models"

	"github.com/go-pg/pg/v9"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]*models.Thumbnail, error)
	//	GetProduct(ctx context.Context, id string) (models.Product, error)
}

type Service struct{}

func (Service) GetProducts(ctx context.Context) ([]*models.Thumbnail, error) {
	db, exists := ctx.Value("db").(*pg.DB)
	if !exists {
		return nil, errors.New("Database not found in context")
	}

	var products []models.Product
	err := db.Model(&products).Where("public = true").Select()
	if err != nil {
		return nil, err
	}

	var thumbnails []*models.Thumbnail
	for _, product := range products {
		thumbnails = append(thumbnails, &models.Thumbnail{
			Id:        product.Id,
			Name:      product.Name,
			Thumbnail: product.Thumbnail,
			Price:     product.Price,
		})
	}

	return thumbnails, err
}
