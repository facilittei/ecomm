//go:generate mockery --name=Auth --output ./../../mocks/ --filename auth_repository.go --structname AuthRepositoryMock

package repositories

import "context"

// Auth token is stored in a database for reuse
type Auth interface {
	Store(ctx context.Context, token string) error
	Get(ctx context.Context) (string, error)
}
