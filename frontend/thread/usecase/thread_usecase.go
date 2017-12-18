package usecase

import (
	"time"

	. "github.com/tsrnd/goweb5/frontend/thread"
	repos "github.com/tsrnd/goweb5/frontend/thread/repository"
	"github.com/tsrnd/goweb5/frontend/user"
)

// ThreadUsecase interface
type ThreadUsecase interface {
	CreatedAtDate(time time.Time) string
	NumReplies(id int) (count int)
	Posts(id int) (posts []Post, err error)
	CreateThread(userId int, topic string) (conv Thread, err error)
	CreatePost(userId int, conv Thread, body string) (post Post, err error)
	Threads() (threads []Thread, err error)
	ThreadByUUID(uuid string) (conv Thread, err error)
	User(id int) *user.User
}

type threadUsecase struct {
	threadRepos repos.ThreadRepository
}

func (this *threadUsecase) CreatedAtDate(time time.Time) string {
	return this.threadRepos.CreatedAtDate(time)
}

func (this *threadUsecase) NumReplies(id int) (count int) {
	return this.threadRepos.NumReplies(id)
}
func (this *threadUsecase) Posts(id int) (posts []Post, err error) {
	return this.threadRepos.Posts(id)
}
func (this *threadUsecase) CreateThread(userId int, topic string) (conv Thread, err error) {
	return this.threadRepos.CreateThread(userId, topic)
}
func (this *threadUsecase) CreatePost(userId int, conv Thread, body string) (post Post, err error) {
	return this.threadRepos.CreatePost(userId, conv, body)
}
func (this *threadUsecase) Threads() (threads []Thread, err error) {
	return this.threadRepos.Threads()
}
func (this *threadUsecase) ThreadByUUID(uuid string) (conv Thread, err error) {
	return this.threadRepos.ThreadByUUID(uuid)
}
func (this *threadUsecase) User(id int) *user.User {
	return this.threadRepos.User(id)
}

// NewThreadUsecase func
func NewThreadUsecase(a repos.ThreadRepository) ThreadUsecase {
	return &threadUsecase{a}
}
