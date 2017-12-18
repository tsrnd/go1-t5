package usecase

import (
	"net/http"

	model "github.com/tsrnd/goweb5/frontend/user"
	repos "github.com/tsrnd/goweb5/frontend/user/repository"
)

// UserUsecase interface
type UserUsecase interface {
	CreateSession(email string, id int) (*model.Session, error)
	SessionByID(id int) (*model.Session, error)
	SessionByCookie(cookie *http.Cookie) (model.Session, error)
	Check(session model.Session) (valid bool, err error)
	DeleteByUUID(UUID string) (err error)
	User(userID int) (*model.User, error)
	SessionDeleteAll() (err error)
	Create(name string, email string, password string) (int, error)
	Delete(id int) (err error)
	Update(id int, name string, email string) (err error)
	UserDeleteAll() (err error)
	Users() (users []model.User, err error)
	UserByEmail(email string) (*model.User, error)
	UserByUUID(uuid string) (*model.User, error)
}

type userUsecase struct {
	userRepos repos.UserRepository
}

func (a *userUsecase) CreateSession(email string, id int) (*model.Session, error) {
	return a.userRepos.CreateSession(email, id)
}

func (a *userUsecase) SessionByID(id int) (*model.Session, error) {
	return a.userRepos.SessionByID(id)
}
func (a *userUsecase) SessionByCookie(cookie *http.Cookie) (model.Session, error) {
	return a.userRepos.SessionByCookie(cookie)
}

func (a *userUsecase) Check(session model.Session) (valid bool, err error) {
	return a.userRepos.Check(&session)
}

func (a *userUsecase) DeleteByUUID(UUID string) (err error) {
	return a.userRepos.DeleteByUUID(UUID)
}

func (a *userUsecase) User(userID int) (*model.User, error) {
	return a.userRepos.User(userID)
}

func (a *userUsecase) SessionDeleteAll() (err error) {
	return a.userRepos.SessionDeleteAll()
}

func (a *userUsecase) Create(name string, email string, password string) (int, error) {
	return a.userRepos.Create(name, email, password)
}
func (a *userUsecase) Delete(id int) (err error) {
	return a.userRepos.Delete(id)
}
func (a *userUsecase) Update(id int, name string, password string) (err error) {
	return a.userRepos.Update(id, name, password)
}

func (a *userUsecase) UserDeleteAll() (err error) {
	return a.userRepos.UserDeleteAll()
}
func (a *userUsecase) Users() (users []model.User, err error) {
	return a.userRepos.Users()
}
func (a *userUsecase) UserByEmail(email string) (*model.User, error) {
	return a.userRepos.UserByEmail(email)
}
func (a *userUsecase) UserByUUID(uuid string) (*model.User, error) {
	return a.userRepos.UserByUUID(uuid)
}

// NewUserUsecase func
func NewUserUsecase(a repos.UserRepository) UserUsecase {
	return &userUsecase{a}
}
