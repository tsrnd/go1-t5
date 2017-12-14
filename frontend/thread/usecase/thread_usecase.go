package usecase

import (
	repos "github.com/tsrnd/goweb5/frontend/thread/repository"
)

// ThreadUsecase interface
type ThreadUsecase interface {
}

type threadUsecase struct {
	threadRepos repos.ThreadRepository
}
