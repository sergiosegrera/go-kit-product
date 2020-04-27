package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/go-kit-product/product/models"
	"github.com/sergiosegrera/go-kit-product/product/service"
)

type GetProductsResponse struct {
	Products []*models.Thumbnail `json:"products"`
}

func MakeGetProductsEndpoint(svc service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		thumbnails, err := svc.GetProducts(ctx)
		return GetProductsResponse{Products: thumbnails}, err
	}
}
