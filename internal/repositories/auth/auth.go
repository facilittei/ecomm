package repositories

import "context"

// Auth token is stored in a database for reuse
type Auth interface {
	Store(ctx context.Context, token string) error
	Get(ctx context.Context) (string, error)
}
