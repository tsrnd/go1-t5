package http

import (
	"github.com/go-chi/chi"
	"github.com/tsrnd/goweb5/frontend/home/usecase"
	"github.com/tsrnd/goweb5/frontend/services/cache"
)

// HomeController type
type HomeController struct {
	Usecase usecase.HomeUsecase
	Cache   cache.Cache
}

// NewHomeController func
func NewHomeController(r *chi.Mux, uc usecase.HomeUsecase, c cache.Cache) *HomeController {
	handler := &HomeController{
		Usecase: uc,
		Cache:   c,
	}
	return handler
}
