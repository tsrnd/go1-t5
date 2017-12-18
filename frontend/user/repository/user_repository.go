package repository

import (
	"net/http"

	model "github.com/tsrnd/goweb5/frontend/user"
)

// UserRepository interface
type UserRepository interface {
	CreateSession(email string, id int) (*model.Session, error)
	SessionByID(id int) (*model.Session, error)
	SessionByCookie(cookie *http.Cookie) (model.Session, error)
	Check(session *model.Session) (valid bool, err error)
	DeleteByUUID(UUID string) (err error)
	User(userID int) (*model.User, error)
	SessionDeleteAll() (err error)
	Create(name string, email string, password string, image string) (int, error)
	Delete(id int) (err error)
	Update(id int, name string, email string) (err error)
	UserDeleteAll() (err error)
	Users() (users []model.User, err error)
	UserByEmail(email string) (*model.User, error)
	UserByUUID(uuid string) (*model.User, error)
}
