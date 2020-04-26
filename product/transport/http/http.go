package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product/db"
	"github.com/sergiosegrera/store/product/endpoints"
	"github.com/sergiosegrera/store/product/models"
	"github.com/sergiosegrera/store/product/service"
	"github.com/sergiosegrera/store/product/transport/http/handlers"
)

func Serve() error {
	svc := service.Service{}
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// Create and add database connection to context
	options := &pg.Options{
		Addr:     "db:5432",
		User:     "product",
		Database: "product",
		Password: "verysecuremuchwow",
	}

	db, err := db.NewConnection(options)
	if err != nil {
		return err
	}

	// Add test data
	product := models.Product{
		Name:        "White T-Shirt",
		Thumbnail:   "https://imgur.com/qEOvdMp",
		Images:      []string{"https://imgur.com/qEOvdMp", "https://imgur.com/qEOvdMp"},
		Description: "Plain white T-Shirt",
		Price:       30,
		Public:      true,
	}

	result, err := db.Model(&product).Returning("id").Insert()
	if err != nil {
		return err
	}

	option := models.Option{
		ProductId: int64(result.RowsReturned()),
		Name:      "Small",
		Stock:     30,
	}
	err = db.Insert(&option)
	if err != nil {
		return err
	}

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
