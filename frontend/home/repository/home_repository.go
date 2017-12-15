package repository

import (
	"database/sql"

	model "github.com/tsrnd/goweb5/frontend/user"
)

const SALT string = "VUDANG"

// UserRepository interface
type HomeRepository interface {
	Index() (*model.Session, error)
}

type homeRepository struct {
	DB *sql.DB
}

func (m *homeRepository) Index() (session *model.Session, err error) {
	return
}

func NewHomeRepository(db *sql.DB) HomeRepository {
	return &homeRepository{db}
}
