package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	body, _ := ioutil.ReadFile("view/layout.html")
	fmt.Fprintf(w, string(body))
}

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

	// server := http.Server{
	// 	Addr: "127.0.0.1:8082",
	// }
	// http.HandleFunc("/", indexHandler)
	// server.ListenAndServe()
}
