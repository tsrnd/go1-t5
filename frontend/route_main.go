package main

import (
  "net/http"
  "fmt"
  "goweb5/frontend/data"
)


// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
  _, err := session(writer,request)
  if err == nil {
    http.Redirect(writer, request,  "/", 302)
  } else {
    t := parseTemplateFiles("login.layout", "public.navbar", "login")
    t.Execute(writer, nil)
  }
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
  generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
  err := request.ParseForm()
  if err != nil {
    fmt.Println(err, "Cannot parse form")
  }
  user := data.User{
    Name: request.PostFormValue("name"),
    Email: request.PostFormValue("email"),
    Password: request.PostFormValue("password"),    
  }
  if err := user.Create(); err != nil {
    fmt.Println(err, "Cannot create user")
  }
  http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {  
  err := request.ParseForm()
  user, err := data.UserByEmail(request.PostFormValue("email"))
  if err != nil {
    fmt.Println(err, "Cannot find user")
  }
  if user.Password == data.Encrypt(request.PostFormValue("password")) {
    session, err := user.CreateSession()
    if err != nil {
      fmt.Println(err, "Cannot create session")
    }
    cookie := http.Cookie{
      Name:      "_cookie", 
      Value:     session.Uuid,
      HttpOnly:  true,
    }
    http.SetCookie(writer, &cookie)
    http.Redirect(writer, request, "/", 302)
  } else {
    http.Redirect(writer, request, "/login", 302)
  }
  
}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
  cookie, err := request.Cookie("_cookie")
  if err != http.ErrNoCookie {
    fmt.Println(err, "Failed to get cookie")
    session := data.Session{Uuid: cookie.Value}
    session.DeleteByUUID()    
  }  
  http.Redirect(writer, request, "/", 302)
}

func update(w http.ResponseWriter, r *http.Request) {
  
}

func index(writer http.ResponseWriter, request *http.Request) {
  threads, err := data.Threads();
	if err != nil {
	  fmt.Println("Cannot get threads")
	} else {
      userthread := &data.UserThread{Threads: threads}
      sess, err := session(writer, request)
	  if err != nil {
		  generateHTML(writer, userthread, "layout", "public.navbar", "index")
	  } else {
      user,_ := sess.User()
      userthread.User = user
      generateHTML(writer, userthread, "layout", "private.navbar", "index", "user")
	  }
	}
}