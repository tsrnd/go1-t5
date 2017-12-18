package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/tsrnd/goweb5/frontend/services/crypto"

	"github.com/tsrnd/goweb5/frontend/services/util"
	"github.com/tsrnd/goweb5/frontend/user"

	"github.com/go-chi/chi"
	"github.com/tsrnd/goweb5/frontend/services/cache"
	"github.com/tsrnd/goweb5/frontend/user/usecase"
)

// UserController type
type UserController struct {
	Usecase usecase.UserUsecase
	Cache   cache.Cache
}

// NewUserController func
func NewUserController(r *chi.Mux, uc usecase.UserUsecase, c cache.Cache) *UserController {
	handler := &UserController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/user", handler.Update)
	r.Post("/user/update", handler.UpdateUser)
	r.Get("/users", handler.GetAllUser)
	r.Post("/users", handler.signupAccount)
	r.Get("/logout", handler.Logout)
	r.Get("/login", handler.LoginPage)
	r.Post("/login", handler.Login)
	r.Get("/session", handler.Session)
	r.Get("/signup", handler.SignUp)
	return handler
}

// POST /signup
// Create the user account
func (this *UserController) signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}
	if id, err := this.Usecase.Create(request.PostFormValue("name"), request.PostFormValue("email"), request.PostFormValue("password")); err != nil {
		utils.Danger(err, "Cannot create user")
	} else {
		utils.Info(err, fmt.Sprint("Create user", id, "successful"))
	}
	http.Redirect(writer, request, "/login", 302)
}
func (this *UserController) SignUp(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func (this *UserController) LoginPage(writer http.ResponseWriter, request *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

func (this *UserController) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot pasre form")
	}
	cookie, err := request.Cookie("_cookie")
	sess, err := this.Usecase.SessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	}
	user, _ := this.Usecase.User(sess.UserId)
	user.Name = request.FormValue("name")
	if request.FormValue("password") == "" {
		err = this.Usecase.Update(user.Id, user.Name, user.Password)
		if err != nil {
			fmt.Println("Failed update user")
		}
		http.Redirect(writer, request, "/", 302)
	} else {
		user.Password = crypto.HashPassword(request.FormValue("password"), crypto.SALT)
		err = this.Usecase.Update(user.Id, user.Name, user.Password)
		if err != nil {
			fmt.Println("Failed update user")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func (this *UserController) Update(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	sess, err := this.Usecase.SessionByCookie(cookie)
	user, _ := this.Usecase.User(sess.UserId)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, user, "layout", "private.navbar", "update.user")
	}
}

func (this *UserController) Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning(err, "Failed to get cookie")
		err1 := this.Usecase.DeleteByUUID(cookie.Value)
		if err1 != nil {
			utils.Warning(err, "Logout fail")
		}
	}
	http.Redirect(writer, request, "/", 302)
}

func (this *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	p := map[string]string{
		"token": "ss22",
	}
	users, err := this.Usecase.Users()
	if err != nil {
		json.NewEncoder(w).Encode(p)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (this *UserController) Login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := this.Usecase.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		utils.Danger(err, "Cannot find user")
	}
	if user.Password == crypto.HashPassword(request.PostFormValue("password"), crypto.SALT) {
		session, err := this.Usecase.CreateSession(user.Email, user.Id)
		if err != nil {
			utils.Danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
func (this *UserController) Session(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess := user.Session{Uuid: cookie.Value}
		if ok, _ := this.Usecase.Check(sess); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
