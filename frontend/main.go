package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/tsrnd/goweb5/frontend/config"
)

func main() {
	db := config.DB()
	cache := config.Cache()
	router := config.Router(db, cache)
	port := config.Port()
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	filesDir := filepath.Join(workDir, "public")
	FileServer(router, "/static/", http.Dir(filesDir))
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
