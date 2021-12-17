package services

import (
	"context"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/facilittei/ecomm/internal/logging"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	repositories "github.com/facilittei/ecomm/internal/repositories/auth"
)

// Juno handles payment transaction requests
type Juno struct {
	logger         logging.Logger
	junoProvider   *providers.Juno
	authRepository repositories.Auth
}

// NewJuno creates an instance of Juno
func NewJuno(
	logger logging.Logger,
	junoProvider *providers.Juno,
	authRepository repositories.Auth,
) *Juno {
	return &Juno{
		logger:         logger,
		junoProvider:   junoProvider,
		authRepository: authRepository,
	}
}

// Charge customer using Juno payment provider
func (j *Juno) Charge(req payment.Request) (map[string]string, error) {
	ctx := context.Background()
	token, err := j.authRepository.Get(ctx)
	if err != nil {
		auth, err := j.junoProvider.Authenticate()
		if err != nil {
			j.logger.Error("could not establish communication with payment provider: %v", err)
			return map[string]string{
				"status":  "failed",
				"message": "could not establish communication with payment provider",
			}, err
		}

		token = auth.AccessToken

		if err := j.authRepository.Store(ctx, token); err != nil {
			j.logger.Warn("could not store auth token on repository: %v", err)
		}
	}

	return map[string]string{
		"status": "pending",
	}, nil
}
