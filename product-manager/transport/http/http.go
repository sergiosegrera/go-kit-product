package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sergiosegrera/go-kit-product/product-manager/endpoints"
	"github.com/sergiosegrera/go-kit-product/product-manager/service"
	"github.com/sergiosegrera/go-kit-product/product-manager/transport/http/handlers"
)

func Serve(svc *service.Service) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	postProduct := handlers.MakePostProductHandler(endpoints.MakePostProductEndpoint(svc))
	router.Post("/product", postProduct)

	return http.ListenAndServe(":8080", router)
}
