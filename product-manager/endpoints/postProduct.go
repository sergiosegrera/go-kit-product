package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/sergiosegrera/go-kit-product/product-manager/service"
	"github.com/sergiosegrera/go-kit-product/product/models"
)

type PostProductRequest struct {
	Product models.Product `json:"product"`
}

type PostProductResponse struct {
	Error string `json:"error"`
}

func MakePostProductEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostProductRequest)
		err := svc.PostProduct(ctx, req.Product)
		if err != nil {
			return PostProductResponse{Error: err.Error()}, err
		}
		return PostProductResponse{Error: ""}, err
	}
}
