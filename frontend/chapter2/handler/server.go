package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	h := MyHandler{}
	server := http.Server{
		Addr:    "0.0.0.0:8082",
		Handler: &h,
	}
	server.ListenAndServe()
}
