package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/go-kit-product/product/models"
	"github.com/sergiosegrera/go-kit-product/product/service"
)

type GetProductRequest struct {
	Id int64 `json:"id"`
}

type GetProductResponse struct {
	Product *models.Product `json:"product"`
}

func MakeGetProductEndpoint(svc service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		product, err := svc.GetProduct(ctx, req.Id)
		return GetProductResponse{Product: product}, err
	}
}
