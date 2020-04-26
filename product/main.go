package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/sergiosegrera/store/product/transport/http"
)

func main() {
	go func() {
		log.Println("Started the http server")
		err := http.Serve()
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
