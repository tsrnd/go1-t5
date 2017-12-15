package http

import (
	"fmt"
	"gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"

	"github.com/tsrnd/goweb5/frontend/services/util"

	"github.com/go-chi/chi"
	"github.com/tsrnd/goweb5/frontend/services/cache"
	"github.com/tsrnd/goweb5/frontend/thread/usecase"
)

// ThreadController type
type ThreadController struct {
	Usecase usecase.ThreadUsecase
	Cache   cache.Cache
}

// NewThreadController func
func NewThreadController(r *chi.Mux, uc usecase.ThreadUsecase, c cache.Cache) *ThreadController {
	handler := &ThreadController{
		Usecase: uc,
		Cache:   c,
	}
	r.Get("/", handler.Index)
	r.Get("/threads/{uuid}", handler.Show)
	r.Post("/threads/posts", handler.StorePost)
	r.Get("/threads/create", handler.Create)
	r.Post("/threads/store", handler.Store)
	return handler
}

func (this *ThreadController) StorePost(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			utils.Error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/threads/", uuid)
		http.Redirect(writer, request, url, 302)
	}
}

func (this *ThreadController) Store(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			utils.Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func (this *ThreadController) Create(writer http.ResponseWriter, request *http.Request) {
	_, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		utils.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}
func (this *ThreadController) Show(writer http.ResponseWriter, request *http.Request) {
	uuid := chi.URLParam(request, "uuid")
	thread, err := this.Usecase.ThreadByUUID(uuid)
	posts, err := this.Usecase.Posts(thread.Id)
	showPosts := make([]ShowPost, 0)
	for _, post := range posts {
		showPosts = append(showPosts, ShowPost{
			Id:        post.Id,
			Uuid:      post.Uuid,
			Body:      post.Body,
			User:      this.Usecase.User(post.UserId),
			CreatedAt: this.Usecase.CreatedAtDate(post.CreatedAt),
		})
	}

	showThread := ShowThread{
		Id:         thread.Id,
		Uuid:       thread.Uuid,
		Topic:      thread.Topic,
		User:       this.Usecase.User(thread.UserId),
		CreatedAt:  this.Usecase.CreatedAtDate(thread.CreatedAt),
		NumReplies: this.Usecase.NumReplies(thread.Id),
		Posts:      showPosts,
	}
	fmt.Println(showThread.CreatedAt)
	if err != nil {
		utils.Error_message(writer, request, "Cannot read thread")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, showThread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(writer, showThread, "layout", "private.navbar", "private.thread")
		}
	}
}

func (this *ThreadController) Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := this.Usecase.Threads()
	showThreads := make([]ShowThread, 0)
	for _, thread := range threads {
		showThreads = append(showThreads, ShowThread{
			Id:         thread.Id,
			Uuid:       thread.Uuid,
			Topic:      thread.Topic,
			User:       this.Usecase.User(thread.UserId),
			CreatedAt:  this.Usecase.CreatedAtDate(thread.CreatedAt),
			NumReplies: this.Usecase.NumReplies(thread.Id),
		})
	}
	fmt.Println(len(showThreads))
	if err != nil {
		utils.Error_message(writer, request, "Cannot get threads")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, showThreads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(writer, showThreads, "layout", "private.navbar", "index")
		}
	}
}
