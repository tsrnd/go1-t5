package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/go-chi/chi"
	"github.com/user/goweb5/frontend/data"
)

func errHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		errorMessage(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		}
		h(w, r)
	}
}
func writeLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("HandlerFunc called - %v\n", runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
		info(msg)
		h(w, r)
	}
}
func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("HandlerFunc dd - %v\n", runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
		info(msg)
		h.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()
	files := http.FileServer(http.Dir(config.Static))
	r.Handle("/static/*", logHandler(http.StripPrefix("/static/", files)))
	r.Get("/", writeLog(indexHandler))
	r.Get("/err", errHandler)
	r.Get("/login", login)
	r.Post("/authenticate", authenticate)
	r.Get("/signup", signup)
	r.Post("/register", signupAccount)
	r.Post("/logout", logout)
	r.Get("/thread/new", protect(writeLog(newThread)))
	r.Post("/thread/create", protect(writeLog(createThread)))
	r.Post("/thread/post", protect(writeLog(postThread)))
	r.Get("/thread/read", writeLog(readThread))

	server := &http.Server{
		Addr:           config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServeTLS("chapter2/gencert/cert.pem", "chapter2/gencert/key.pem")
}
