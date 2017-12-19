package repository

import (
	"time"

	. "github.com/tsrnd/goweb5/frontend/thread"
	"github.com/tsrnd/goweb5/frontend/user"
)

// ThreadRepository interface
type ThreadRepository interface {
	CreatedAtDate(time time.Time) string
	NumReplies(id int) (count int)
	Posts(id int) (posts []Post, err error)
	CreateThread(userId int, topic string) (conv Thread, err error)
	CreatePost(userId int, conv Thread, body string) (post Post, err error)
	Threads() (threads []Thread, err error)
	ThreadByUUID(uuid string) (conv Thread, err error)
	User(id int) *user.User
	DeletePost(id int) (err error)
}
