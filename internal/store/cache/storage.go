package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"postery/internal/store"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
		Delete(context.Context, int64)
	}
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return Storage{
		Users: &UsersStore{rdb: rdb},
	}
}
