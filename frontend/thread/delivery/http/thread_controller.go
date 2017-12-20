package http

import (
	"fmt"
	"net/http"

	"github.com/tsrnd/goweb5/frontend/services/util"

	"github.com/go-chi/chi"
	"github.com/tsrnd/goweb5/frontend/services/cache"
	threadUC "github.com/tsrnd/goweb5/frontend/thread/usecase"
	userUC "github.com/tsrnd/goweb5/frontend/user/usecase"
	user "github.com/tsrnd/goweb5/frontend/user"
)

// ThreadController type
type ThreadController struct {
	ThreadUC threadUC.ThreadUsecase
	UserUC   userUC.UserUsecase
	Cache    cache.Cache
}

type UserThread struct {
	User *user.User
	Threads []ShowThread
}

// NewThreadController func
func NewThreadController(r *chi.Mux, threadUC threadUC.ThreadUsecase, userUC userUC.UserUsecase, c cache.Cache) *ThreadController {
	handler := &ThreadController{
		ThreadUC: threadUC,
		UserUC:   userUC,
		Cache:    c,
	}
	r.Get("/", handler.Index)
	r.Get("/threads/{uuid}", handler.Show)
	r.Post("/threads/posts", handler.StorePost)
	r.Get("/threads/create", handler.Create)
	r.Post("/threads/store", handler.Store)
	r.Post("/threads/delete", handler.Delete)
	return handler
}
func (this *ThreadController) StorePost(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	sess, err := this.UserUC.SessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := this.ThreadUC.ThreadByUUID(uuid)
		if err != nil {
			utils.Error_message(writer, request, "Cannot read thread")
		}
		if _, err := this.ThreadUC.CreatePost(sess.UserId, thread, body); err != nil {
			utils.Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/threads/", uuid)
		http.Redirect(writer, request, url, 302)
	}
}

func (this *ThreadController) Store(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	sess, err := this.UserUC.SessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		topic := request.PostFormValue("topic")
		if _, err := this.ThreadUC.CreateThread(sess.UserId, topic); err != nil {
			utils.Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func (this *ThreadController) Create(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	_, err = this.UserUC.SessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread", "userlogin")
	}
}
func (this *ThreadController) Show(writer http.ResponseWriter, request *http.Request) {
	uuid := chi.URLParam(request, "uuid")
	thread, err := this.ThreadUC.ThreadByUUID(uuid)
	posts, err := this.ThreadUC.Posts(thread.Id)
	showPosts := make([]ShowPost, 0)
	for _, post := range posts {
		showPosts = append(showPosts, ShowPost{
			Id:        post.Id,
			Uuid:      post.Uuid,
			Body:      post.Body,
			User:      this.ThreadUC.User(post.UserId),
			CreatedAt: this.ThreadUC.CreatedAtDate(post.CreatedAt),
		})
	}

	showThread := ShowThread{
		Id:         thread.Id,
		Uuid:       thread.Uuid,
		Topic:      thread.Topic,
		User:       this.ThreadUC.User(thread.UserId),
		CreatedAt:  this.ThreadUC.CreatedAtDate(thread.CreatedAt),
		NumReplies: this.ThreadUC.NumReplies(thread.Id),
		Posts:      showPosts,
	}
	fmt.Println(showThread.CreatedAt)
	if err != nil {
		utils.Error_message(writer, request, "Cannot read thread")
	} else {
		cookie, err := request.Cookie("_cookie")
		_, err = this.UserUC.SessionByCookie(cookie)
		if err != nil {
			utils.GenerateHTML(writer, showThread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(writer, showThread, "layout", "private.navbar", "private.thread", "userlogin")
		}
	}
}

func (this *ThreadController) Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := this.ThreadUC.Threads()
	showThreads := make([]ShowThread, 0)
	if err != nil {
		utils.Error_message(writer, request, "Cannot get threads")
	} else {
		userthread := UserThread{}
		for _, thread := range threads {
			showThreads = append(showThreads, ShowThread{
				Id:         thread.Id,
				Uuid:       thread.Uuid,
				Topic:      thread.Topic,
				User:       this.ThreadUC.User(thread.UserId),
				CreatedAt:  this.ThreadUC.CreatedAtDate(thread.CreatedAt),
				NumReplies: this.ThreadUC.NumReplies(thread.Id),
			})
		}
		userthread.Threads = showThreads
		cookie, err := request.Cookie("_cookie")
		sess, err := this.UserUC.SessionByCookie(cookie)
		if err != nil {
			utils.GenerateHTML(writer, userthread, "layout", "public.navbar", "index")
		} else {
			userLogin, _ := this.UserUC.User(sess.UserId)
			userthread.User = userLogin
			utils.GenerateHTML(writer, userthread, "layout", "private.navbar", "index", "userlogin")
		}
	}
}

//Delete thread of userlogin
func (this *ThreadController) Delete(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	sess, err := this.UserUC.SessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	}
	err = request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}
	threadId := request.PostFormValue("threadId")
	userLogin,_ := this.UserUC.User(sess.UserId)
	thread,_ := this.ThreadUC.ThreadByID(threadId)
	if userLogin.Id != thread.UserId {
		utils.Danger(err, "Fobidden! You can't delete a thread of another user!")
	} else  {
		if err := this.ThreadUC.DeleteThread(threadId); err != nil {
			utils.Danger(err, "Cannot delete thread")
		}
	}
	http.Redirect(writer, request, "/", 302)
}
