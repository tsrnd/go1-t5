package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

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

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", writeLog(indexHandler))
	mux.HandleFunc("/err", errHandler)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/register", signupAccount)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/thread/new", protect(writeLog(newThread)))
	mux.HandleFunc("/thread/create", protect(writeLog(createThread)))
	mux.HandleFunc("/thread/post", protect(writeLog(postThread)))
	mux.HandleFunc("/thread/read", writeLog(readThread))

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServeTLS("chapter2/gencert/cert.pem", "chapter2/gencert/key.pem")
}
