package main

import (
	"fmt"
	"net/http"
	"goweb5/frontend/data"
	"strconv"
)

type SingleThread struct {
	User data.User
	Thread *data.Thread
}

func newThread(w http.ResponseWriter, r * http.Request) {
	sess, _ :=  session(w,r)
	user, err := sess.User()
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		userthread := &data.UserThread{}
		userthread.User = user
		generateHTML(w, userthread, "layout", "private.navbar", "user", "thread")
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
	  fmt.Println(writer, request, "Cannot read thread")
	} else {
	  fmt.Println(thread)
	  sess, err := session(writer, request)
	  if err != nil {
		generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
	  } else {
		user,_ := sess.User()
		userthread := &SingleThread{User: user, Thread: &thread}
		generateHTML(writer, userthread, "layout", "private.navbar", "private.thread","user")
	  }
	}
}

// DELETE /thread/{id}/delete
func deleteThread(w http.ResponseWriter, r *http.Request)  {
	_, err := session(w,r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Don't get this thread")
		}
		id,_ := strconv.Atoi(r.PostFormValue("id"))
		fmt.Println(id)
		result := data.DeleteThread(id)
		if result != nil {
			fmt.Println("Delete fail!")
			http.Redirect(w,r,"/",302)
		} else {
			fmt.Println("Delete Success!")
			http.Redirect(w, r, "/", 302)
		}
	}
}

// POST /thread/create
// Create a thread
func createThread(writer http.ResponseWriter, request *http.Request) {
  sess, err := session(writer, request)
  if err != nil {
    http.Redirect(writer, request, "/login", 302)
  } else {    
    err = request.ParseForm()
    if err != nil {
      fmt.Println(err, "Cannot parse form")
    }  
    user, err := sess.User(); if err != nil {
      fmt.Println(err, "Cannot get user from session")
    }
    topic := request.PostFormValue("topic")
    if _, err := user.CreateThread(topic); err != nil {
      fmt.Println(err, "Cannot create thread")
    }
    http.Redirect(writer, request, "/", 302)        
  }
}

// POST /post/create
// create post thread
func postThread(writer http.ResponseWriter, request *http.Request) {
  sess, err := session(writer, request)
  if err != nil {
    http.Redirect(writer, request, "/login", 302)
  } else {
    err = request.ParseForm()
    if err != nil {
      fmt.Println(err, "Cannot parse form")
    }  
    user, err := sess.User(); if err != nil {
      fmt.Println(err, "Cannot get user from session")
    }
    body := request.PostFormValue("body")
    uuid := request.PostFormValue("uuid")
    thread, err := data.ThreadByUUID(uuid); if err != nil {
      fmt.Println(writer, request, "Cannot read thread")
    }
    if _, err := user.CreatePost(thread, body); err != nil {
      fmt.Println(err, "Cannot create post")
    }
    url := fmt.Sprint("/thread/read?id=", uuid)
    http.Redirect(writer, request, url , 302)        
  }
}