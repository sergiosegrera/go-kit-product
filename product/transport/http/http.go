package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"

	"github.com/sergiosegrera/store/product/endpoints"
	"github.com/sergiosegrera/store/product/service"
	"github.com/sergiosegrera/store/product/transport/http/handlers"
)

func Serve(db *pg.DB) error {
	svc := service.Service{}
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	getProducts := handlers.MakeGetProductsHandler(endpoints.MakeGetProductsEndpoint(svc))
	getProduct := handlers.MakeGetProductHandler(endpoints.MakeGetProductEndpoint(svc))

	router.Get("/products", getProducts)
	router.Get("/product/{id}", getProduct)

	return http.ListenAndServe(":8080", router)
}
