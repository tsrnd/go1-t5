package psql

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/tsrnd/goweb5/frontend/services/crypto"
	"github.com/tsrnd/goweb5/frontend/services/util"
	model "github.com/tsrnd/goweb5/frontend/user"
	"github.com/tsrnd/goweb5/frontend/user/repository"
)

type userRepository struct {
	DB *sql.DB
}

func (m *userRepository) CreateSession(email string, id int) (*model.Session, error) {
	const statement = `
  insert into sessions 
  (uuid, email, user_id, created_at) 
  values ($1, $2, $3, $4) 
  returning id, uuid, email, user_id, created_at
  `
	var session = model.Session{}
	var err error
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return &session, err
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(utils.CreateUUID(), email, id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return &session, err
}

func (m *userRepository) SessionByID(id int) (*model.Session, error) {
	session := model.Session{}
	err := m.DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return &session, err
}
func (m *userRepository) SessionByCookie(cookie *http.Cookie) (sess model.Session, err error) {
	if cookie == nil {
		sess = model.Session{}
		err = errors.New("Invalid session")
	} else {
		sess = model.Session{Uuid: cookie.Value}
		if ok, _ := m.Check(&sess); !ok {
			err = errors.New("Invalid session")
		}
	}
	return sess, err
}
func (m *userRepository) Check(session *model.Session) (valid bool, err error) {
	err = m.DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}
func (m *userRepository) DeleteByUUID(UUID string) (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(UUID)
	return
}

func (m *userRepository) User(userID int) (*model.User, error) {
	user := model.User{}
	err := m.DB.QueryRow("SELECT id, uuid, name, password, email, created_at FROM users WHERE id = $1", userID).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Password, &user.Email, &user.CreatedAt)
	return &user, err
}

func (m *userRepository) SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = m.DB.Exec(statement)
	return
}

func (m *userRepository) Create(name string, email string, password string) (int, error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var id int
	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(utils.CreateUUID(), name, email, crypto.HashPassword(password, crypto.SALT), time.Now()).Scan(&id)
	return id, err
}

func (m *userRepository) Delete(id int) error {
	statement := "delete from users where id = $1"
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (m *userRepository) Update(id int, name string, password string) error {
	statement := "update users set name = $2, password = $3 where id = $1"
	stmt, err := m.DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, password)
	return err
}

func (m *userRepository) UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = m.DB.Exec(statement)
	return
}

func (m *userRepository) Users() (users []model.User, err error) {
	rows, err := m.DB.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := model.User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Get a single user given the email
func (m *userRepository) UserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := m.DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return &user, err
}

// Get a single user given the UUID
func (m *userRepository) UserByUUID(uuid string) (*model.User, error) {
	user := model.User{}
	err := m.DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return &user, err
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db}
}
