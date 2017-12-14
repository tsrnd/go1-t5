package main

import (
	"log"
	"net/http"

	"github.com/tsrnd/goweb5/frontend/config"
)

func main() {
	db := config.DB()
	cache := config.Cache()
	router := config.Router(db, cache)
	port := config.Port()
	if err := http.ListenAndServe(p, r); err != nil {
		log.Fatal(err)
	}
}
