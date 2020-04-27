package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sergiosegrera/go-kit-product/product/endpoints"
	"github.com/sergiosegrera/go-kit-product/product/service"
	"github.com/sergiosegrera/go-kit-product/product/transport/http/handlers"
)

func Serve(svc *service.Service) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	getProducts := handlers.MakeGetProductsHandler(endpoints.MakeGetProductsEndpoint(svc))
	getProduct := handlers.MakeGetProductHandler(endpoints.MakeGetProductEndpoint(svc))

	router.Get("/products", getProducts)
	router.Get("/product/{id}", getProduct)

	return http.ListenAndServe(":8080", router)
}
