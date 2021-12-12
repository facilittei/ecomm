package repositories

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Redis manages authentication tokens
type Redis struct {
	ctx context.Context
	db  *redis.Client
}

// NewRedis creates an instance of Redis auth
func NewRedis(db *redis.Client) Auth {
	return &Redis{
		db: db,
	}
}

// Store authentication token
func (r Redis) Store(ctx context.Context, token string) error {
	return r.db.Set(ctx, "token", token, 50*time.Minute).Err()
}

// Get authentication token
func (r Redis) Get(ctx context.Context) (string, error) {
	result, err := r.db.Get(ctx, "token").Result()
	switch {
	case err == redis.Nil:
		return "", AuthError{Message: fmt.Sprintf("there is no token stored")}
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}
