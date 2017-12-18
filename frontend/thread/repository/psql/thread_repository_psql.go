package psql

import (
	"database/sql"
	"time"

	"github.com/tsrnd/goweb5/frontend/services/util"
	. "github.com/tsrnd/goweb5/frontend/thread"
	"github.com/tsrnd/goweb5/frontend/thread/repository"
	"github.com/tsrnd/goweb5/frontend/user"
)

type threadRepository struct {
	DB *sql.DB
}

func (this *threadRepository) CreatedAtDate(time time.Time) string {
	return time.Format("Jan 2, 2006 at 3:04pm")
}

func (this *threadRepository) NumReplies(id int) (count int) {
	rows, err := this.DB.Query("SELECT count(*) FROM posts where thread_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func (this *threadRepository) Posts(id int) (posts []Post, err error) {
	rows, err := this.DB.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func (this *threadRepository) CreateThread(userId int, topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := this.DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(utils.CreateUUID(), topic, userId, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

func (this *threadRepository) CreatePost(userId int, conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := this.DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(utils.CreateUUID(), body, userId, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

func (this *threadRepository) Threads() (threads []Thread, err error) {
	rows, err := this.DB.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func (this *threadRepository) ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = this.DB.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

func (this *threadRepository) User(id int) *user.User {
	user := user.User{}
	this.DB.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return &user
}

func NewThreadRepository(db *sql.DB) repository.ThreadRepository {
	return &threadRepository{db}
}
