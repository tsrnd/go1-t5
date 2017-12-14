package repository

import (
	"database/sql"
)

// ThreadRepository interface
type ThreadRepository interface {
}

type threadRepository struct {
	DB *sql.DB
}
