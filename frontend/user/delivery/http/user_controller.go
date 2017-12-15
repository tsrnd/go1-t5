package http

import (
	"encoding/json"
	"gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"

	"github.com/tsrnd/goweb5/frontend/services/util"

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
	r.Get("/users", handler.GetAllUser)
	r.Post("/users", handler.UserRegister)
	r.Get("/logout", handler.Logout)
	r.Get("/login", handler.LoginPage)
	r.Post("/login", handler.Login)
	return handler
}
func (ctrl *UserController) LoginPage(writer http.ResponseWriter, request *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}
func (ctrl *UserController) Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.Warning(err, "Failed to get cookie")
		err1 := ctrl.Usecase.DeleteByUUID(cookie.Value)
		if err1 != nil {
			utils.Warning(err, "Logout fail")
		}
	}
	http.Redirect(writer, request, "/", 302)
}

func (ctrl *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	p := map[string]string{
		"token": "ss22",
	}
	users, err := ctrl.Usecase.Users()
	if err != nil {
		json.NewEncoder(w).Encode(p)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UserRegister func
func (ctrl *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}
	var id int
	id, err = ctrl.Usecase.Create(r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("password"))
	if err != nil {
		utils.Danger(err, "Cannot create user")
	}
	utils.Info(id)
	http.Redirect(w, r, "/login", 302)
	// decoder := json.NewDecoder(r.Body)
	// var rr requests.UserRegisterRequest
	// err := decoder.Decode(&rr)
	// if err != nil {
	// 	http.Error(w, "Invalid request body", http.StatusBadRequest)
	// 	return
	// }
	// id, err := repositories.CreateUser(ctrl.DB, rr.Email, rr.Name, rr.Password)
	// if err != nil {
	// 	log.Fatalf("Add user to database error: %s", err)
	// 	http.Error(w, "", http.StatusInternalServerError)
	// 	return
	// }
	// token, err := crypto.GenerateToken()
	// if err != nil {
	// 	log.Fatalf("Generate token Error: %s", err)
	// 	http.Error(w, "", http.StatusInternalServerError)
	// 	return
	// }
	// oneMonth := time.Duration(60*60*24*30) * time.Second
	// err = ctrl.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(id), oneMonth)
	// if err != nil {
	// 	log.Fatalf("Add token to redis Error: %s", err)
	// 	http.Error(w, "", http.StatusInternalServerError)
	// 	return
	// }
	// p := map[string]string{
	// 	"token": token,
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(p)
}

func (ctrl *UserController) Login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		utils.Danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
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
