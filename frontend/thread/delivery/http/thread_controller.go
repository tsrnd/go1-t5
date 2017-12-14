package http

import (
	"github.com/tsrnd/goweb5/frontend/services/cache"
	"github.com/tsrnd/goweb5/frontend/thread/usecase"
)

// ThreadController type
type ThreadController struct {
	UseCase *usecase.ThreadUsecase
	Cache   *cache.Cache
}
