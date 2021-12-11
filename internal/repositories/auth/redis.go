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
func NewRedis(ctx context.Context, db *redis.Client) Auth {
	return &Redis{
		ctx: ctx,
		db:  db,
	}
}

// Store authentication token
func (r Redis) Store(token string) error {
	return r.db.Set(r.ctx, "token", token, 50*time.Minute).Err()
}

// Get authentication token
func (r Redis) Get() (string, error) {
	result, err := r.db.Get(r.ctx, "token").Result()
	switch {
	case err == redis.Nil:
		return "", AuthError{Message: fmt.Sprintf("there is no token stored")}
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}
