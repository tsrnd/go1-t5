package main

import (
	"fmt"
	"net/http"
	"goweb5/frontend/data"
)

// GET /update/{id}
// Show page update user infor
func showUpdate(w  http.ResponseWriter, r *http.Request) {
	sess, _ := session(w,r)
	user, _ := sess.User()
	userLogin := &data.UserThread{User: user}
	generateHTML(w, userLogin, "layout", "private.navbar", "user", "updateUser")
}

// PUT /update/{id}
// Update user infor
func Update(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.FormValue("name"))
	sess, err := session(w,r)
	if err != nil {
		fmt.Println("Can't get session!")
	} else {
		user, err := sess.User()
		if err != nil {
			fmt.Println("Can't get user from session!")
			http.Redirect(w, r, "/", 302)
		}
		pass := r.FormValue("password")
		if pass == "" {
			user.Name = r.FormValue("name")
			user.ChangeName()
			http.Redirect(w, r, "/", 302)
		} else {
			pass_confirm := r.FormValue("password_confirmation")
			if pass != pass_confirm {
				fmt.Println("Password confirm doesn't match!")
				http.Redirect(w, r, "/update/", 302)
			}
		user.Name = r.FormValue("name")
		user.Password = data.Encrypt(pass)
		user.Update()	
		http.Redirect(w, r, "/login", 302)
		}
	}
}