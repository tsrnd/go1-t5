package http

import (
	"github.com/tsrnd/goweb5/frontend/user"
)

type ShowThread struct {
	Id         int
	Uuid       string
	Topic      string
	User       *user.User
	CreatedAt  string
	NumReplies int
	Posts      []ShowPost
}
type ShowPost struct {
	Id        int
	Uuid      string
	Body      string
	User      *user.User
	CreatedAt string
}
