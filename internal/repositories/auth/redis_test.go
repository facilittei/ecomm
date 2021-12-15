package repositories

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedis_Get_no_record(t *testing.T) {
	t.Parallel()
	rdb, mock := redismock.NewClientMock()
	auth := NewRedis(rdb)
	ctx := context.Background()
	mock.ExpectGet("token").RedisNil()
	token, err := auth.Get(ctx)
	assert.True(t, errors.As(err, &AuthError{}))
	assert.Empty(t, token)
}

func TestRedis_Get_token(t *testing.T) {
	t.Parallel()
	rdb, mock := redismock.NewClientMock()
	auth := NewRedis(rdb)
	ctx := context.Background()
	hash := "1234abc"
	mock.ExpectGet("token").SetVal(hash)
	token, err := auth.Get(ctx)
	assert.False(t, errors.As(err, &AuthError{}))
	assert.Equal(t, hash, token)
}

func TestRedis_Store(t *testing.T) {
	t.Parallel()
	rdb, mock := redismock.NewClientMock()
	auth := NewRedis(rdb)
	ctx := context.Background()
	hash := "1234abc"
	mock.ExpectSet("token", hash, 50*time.Minute).SetVal(hash)
	err := auth.Store(ctx, hash)
	require.Empty(t, err)
}
