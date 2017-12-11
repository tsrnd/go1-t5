package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "Goweb5. This is the frontend system")
}

// func testHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	setCors(w)
// 	fmt.Fprint(w, "Hello world")
// }

func setCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5001")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {

	// add router and routes
	router := httprouter.New()
	router.GET("/", indexHandler)

	http.ListenAndServe(":8082", router)
}
