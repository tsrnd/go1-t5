package main

import (
  "net/http"
)

func main() {
  // handle static assets
  mux := http.NewServeMux()
  files := http.FileServer(http.Dir("public"))
  mux.Handle("/static/", http.StripPrefix("/static/", files))

  //
  // all route patterns matched here
  // route handler functions defined in other files
  //

  // index
  mux.HandleFunc("/", index)
  // error

  // defined in route_auth.go
  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", signupAccount)
  mux.HandleFunc("/authenticate", authenticate)

  // update user
  mux.HandleFunc("/update/", showUpdate)
  mux.HandleFunc("/update", Update)

  // defined in route-thread.go
  mux.HandleFunc("/thread/new", newThread)
  mux.HandleFunc("/thread/read", readThread)
  mux.HandleFunc("/thread/delete", deleteThread)
  mux.HandleFunc("/post/create", postThread)
  
  // starting up the server
  server := &http.Server{
    Addr:           "127.0.0.1:8080",
    Handler:        mux,
  }
  server.ListenAndServe()
}