package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Record Not Found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db: db},
		Users: &UsersStore{db: db},
	}
}
