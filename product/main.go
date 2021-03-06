package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/go-pg/pg/v9"
	"github.com/sergiosegrera/go-kit-product/product/db"
	"github.com/sergiosegrera/go-kit-product/product/service"
	"github.com/sergiosegrera/go-kit-product/product/transport/http"
)

func main() {
	// Connect to DB
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
	defer db.Close()

	// Start attach db and start http server
	go func() {
		log.Println("Started the http server")
		err := http.Serve(service.NewService(db))
		if err != nil {
			log.Println("The http server panicked:", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)
}
