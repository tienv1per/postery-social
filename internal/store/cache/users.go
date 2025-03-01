package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"postery/internal/store"
	"time"
)

type UsersStore struct {
	rdb *redis.Client
}

const UserExpTime = time.Minute

func (userStore *UsersStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%v", userID)
	data, err := userStore.rdb.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	var user *store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (userStore *UsersStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("user-%v", user.ID)
	json, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return userStore.rdb.SetEX(ctx, cacheKey, json, UserExpTime).Err()
}
