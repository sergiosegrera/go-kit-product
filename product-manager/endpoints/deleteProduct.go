package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sergiosegrera/go-kit-product/product-manager/service"
)

type DeleteProductRequest struct {
	Id int64 `json:"id"`
}

type GetProductResponse struct {
	Error string `json:"error"`
}

func MakeDeleteProductEndpoint(svc service.ProductManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := svc.DeleteProduct(ctx, req.Id)
		return GetProductResponse{Error: err.Error()}, err
	}
}
