package main

import (
	"fmt"
	"net/http"

	"github.com/user/goweb5/frontend/data"
)

func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	generateHTML(w, nil, "layout", "private.navbar", "new.thread")
}
func createThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if _, err = user.CreateThread(r.PostFormValue("topic")); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(w, r, "/", 302)
	}
}
func readThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		errorMessage(w, r, "Cannot get thread by uuid")
	}
	_, err = session(w, r)
	if err != nil {
		generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
	} else {
		generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
	}
}
func postThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	err = r.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user, err := sess.User()
	if err != nil {
		danger(err, "Cannot get user from session")
	}
	uuid := r.PostFormValue("uuid")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		danger(err, "Cannot get thread from uuid")
	}
	_, err = user.CreatePost(thread, r.PostFormValue("body"))
	if err != nil {
		danger(thread, "Cannot create post")
	}
	url := fmt.Sprint("/thread/read?id=", uuid)
	http.Redirect(w, r, url, 302)

}
