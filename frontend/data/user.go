package data

import(
  "time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string   
	Password  string
	CreatedAt time.Time
}
  
type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

type UserThread struct {
	User User
	Threads []Thread
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
  db := db()
  defer db.Close()  
  err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
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
// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	db := db()
	defer db.Close()
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
	  return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
	  return
	}    
	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	db := db()
	defer db.Close()  
	session = Session{}
	err = db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Id).
		  Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)  
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	db := db()
	defer db.Close()  
	user = User{}
	err = db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserId).
		  Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)  
	return  
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
  db := db()
  defer db.Close()  
  user = User{}
  err = db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)  
  if err != nil {
    return
  }  
  return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
  db := db()
  defer db.Close()
  // Postgres does not automatically return the last insert id, because it would be wrong to assume 
  // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this 
  // information from postgres.
  statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
  stmt, err := db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  
  // use QueryRow to return a row and scan the returned id into the User struct
  err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
  if err != nil {
    return
  }  
  return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
  db := db()
  defer db.Close()
  statement := "delete from sessions where uuid = $1"
  stmt, err := db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  
  _, err = stmt.Exec(session.Uuid)  
  if err != nil {
    return
  }  
  return
}

// Update user information in the database
func (user *User) ChangeName() (err error) {
  db := db()
  defer db.Close()
  statement := "update users set name = $2 where id = $1"
  stmt, err := db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  
  _, err = stmt.Exec(user.Id, user.Name)  
  if err != nil {
    return
  }  
  return
}

// Update user information in the database
func (user *User) Update() (err error) {
  db := db()
  defer db.Close()
  statement := "update users set name = $2, password = $3 where id = $1"
  stmt, err := db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  
  _, err = stmt.Exec(user.Id, user.Name, user.Password)  
  if err != nil {
    return
  }  
  return
}