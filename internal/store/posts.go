package store

import (
	"context"
	"database/sql"
)

type PostsStore struct {
	db *sql.DB
}

func (store *PostsStore) Create(context context.Context) error {
	return nil
}
