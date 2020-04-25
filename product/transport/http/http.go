package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/store/product/db"
	"github.com/sergiosegrera/store/product/endpoints"
	"github.com/sergiosegrera/store/product/models"
	"github.com/sergiosegrera/store/product/service"
)

func Server() {
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
		panic(err)
	}
	// Test
	product := models.Product{
		Name:        "White T-Shirt",
		Thumbnail:   "https://imgur.com/qEOvdMp",
		Images:      []string{"https://imgur.com/qEOvdMp", "https://imgur.com/qEOvdMp"},
		Description: "Plain white T-Shirt",
		Options: []*models.Option{
			&models.Option{
				Name:  "Medium",
				Stock: 30,
			},
		},
		Price:  30,
		Public: true,
	}

	err = db.Insert(&product)
	if err != nil {
		log.Println(err)
	}

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	getProductsHandler := httptransport.NewServer(
		endpoints.MakeGetProductsEndpoint(svc),
		decodeGetProductsRequest,
		encodeResponse,
	).ServeHTTP

	router.Get("/products", getProductsHandler)

	http.ListenAndServe(":8080", router)
}

func decodeGetProductsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

//func GetProducts(ept endpoint.Endpoint) func(w http.ResponseWriter, r *http.Request) {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		resp, err := ept(r.Context(), nil)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//		}
//		j, _ := json.Marshal(resp)
//		w.Write(j)
//	}
//}

func main() {
	fmt.Println("vim-go")
}
