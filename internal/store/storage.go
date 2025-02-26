package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("Resource Not Found")
	ErrConflict          = errors.New("Resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64, PaginationFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(context.Context, *User, string, time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerID int64, userID int64) error
		Unfollow(ctx context.Context, followerID int64, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db: db},
		Users:     &UserStore{db: db},
		Comments:  &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}

func withTransaction(db *sql.DB, ctx context.Context, fn func(tx *sql.Tx) error) error {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(transaction); err != nil {
		_ = transaction.Rollback()
		return err
	}

	return transaction.Commit()
}
