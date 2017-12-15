package usecase

import (
	repos "github.com/tsrnd/goweb5/frontend/home/repository"
	model "github.com/tsrnd/goweb5/frontend/user"
)

// UserUsecase interface
type HomeUsecase interface {
	Index() (*model.Session, error)
}

type homeUsecase struct {
	homeRepos repos.HomeRepository
}

func (a *homeUsecase) Index() (*model.Session, error) {
	return a.homeRepos.Index()
}

// NewUserUsecase func
func NewHomeUsecase(a repos.HomeRepository) HomeUsecase {
	return &homeUsecase{a}
}
