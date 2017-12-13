package data

type User struct {
	Model
	UserName string  `schema:"username"`
	Password string  `schema:"password"`
	Email    string  `schema:"email"`
	Gender   string  `schema:"gender"`
	Role     string  `schema:"role"`
	Avatar   string  `schema:"avatar"`
	Phone    string  `schema:"phone"`
	Address  string  `schema:"address"`
	Orders   []Order //has many order
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	return
}
