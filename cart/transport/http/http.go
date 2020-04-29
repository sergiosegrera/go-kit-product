package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sergiosegrera/go-kit-product/cart/endpoints"
	"github.com/sergiosegrera/go-kit-product/cart/service"
	"github.com/sergiosegrera/go-kit-product/cart/transport/http/handlers"
)

func Serve(svc *service.Service) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	postCart := handlers.MakePostCartHandler(endpoints.MakePostCartEndpoint(svc))

	router.Post("/cart", postCart)

	return http.ListenAndServe(":8080", router)
}
