package user

import "time"

const IMG_BASE_URL = "public/uploads/images"
const IMG_PUBLIC_URL = "static/uploads/images"

// User struct
type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
