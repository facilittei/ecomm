package services

import (
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	repositories "github.com/facilittei/ecomm/internal/repositories/auth"
	"log"
)

// Juno handles payment transaction requests
type Juno struct {
	junoProvider   *providers.Juno
	authRepository repositories.Auth
}

// NewJuno creates an instance of Juno
func NewJuno(junoProvider *providers.Juno, authRepository repositories.Auth) *Juno {
	return &Juno{
		junoProvider:   junoProvider,
		authRepository: authRepository,
	}
}

// Charge customer using Juno payment provider
func (j *Juno) Charge() (map[string]string, error) {
	token, err := j.authRepository.Get()
	if err != nil {
		auth, err := j.junoProvider.Authenticate()
		if err != nil {
			return map[string]string{
				"status":  "failed",
				"message": "could not establish communication with payment provider",
			}, err
		}
		token = auth.AccessToken

		if err := j.authRepository.Store(token); err != nil {
			log.Printf("could not store auth token on repository: %v", err) //TODO: logging err to centralized system
		}
	}

	return map[string]string{
		"status": "pending",
	}, nil
}
