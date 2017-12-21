package main

import (
	"fmt"
	"html/template"
	"net/http"
	"goweb5/frontend/data"
	"errors"
)
// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request)(sess data.Session, err error){
	cookie, err := request.Cookie("_cookie")
	if err == nil {
	  sess = data.Session{Uuid: cookie.Value}
	  if ok, _ := sess.Check(); !ok {
		err = errors.New("Invalid session")
	  }
	}  
	return
}
  
// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(writer http.ResponseWriter, data1 interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data1)
}