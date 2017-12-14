package usecase

import (
	repos "github.com/tsrnd/go-clean-arch/user/repository"
)

// UserUsecase interface
type UserUsecase interface {
}

type userUsecase struct {
	userRepos repos.UserRepository
}
