package services

import (
	"context"
	"errors"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/facilittei/ecomm/internal/logging"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	authRepository "github.com/facilittei/ecomm/internal/repositories/auth"
	chargeRepository "github.com/facilittei/ecomm/internal/repositories/charge"
	"github.com/google/uuid"
)

// Juno handles payment transaction requests
type Juno struct {
	logger           logging.Logger
	junoProvider     *providers.Juno
	authRepository   authRepository.Auth
	chargeRepository chargeRepository.Charge
}

// NewJuno creates an instance of Juno
func NewJuno(
	logger logging.Logger,
	junoProvider *providers.Juno,
	authRepository authRepository.Auth,
	chargeRepository chargeRepository.Charge,
) *Juno {
	return &Juno{
		logger:           logger,
		junoProvider:     junoProvider,
		authRepository:   authRepository,
		chargeRepository: chargeRepository,
	}
}

// Charge customer using Juno payment provider
func (j *Juno) Charge(req payment.Request) (map[string]string, error) {
	ctx := context.Background()

	err := j.chargeRepository.Store(ctx, payment.Charge{
		ID:          uuid.New(),
		SKU:         "abcd123",
		Amount:      req.Amount,
		Description: req.Description,
		Customer:    req.Customer,
		History: []payment.ChargeHistory{
			{Status: "STARTED", Description: "Transaction has started"},
		},
	})
	if err != nil {
		j.logger.Error("could not store charge transaction start: %v", err)
		return map[string]string{
			"status":  "failed",
			"message": "could not store charge transaction start",
		}, err
	}

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
		if err := j.authRepository.Store(ctx, auth.AccessToken); err != nil {
			j.logger.Warn("could not store auth token on repository: %v", err)
		}
	}

	j.junoProvider.UseToken(token)

	res, err := j.junoProvider.Charge(req)
	if err != nil {
		return nil, err
	}

	if pay, ok := res.(*providers.JunoPayResponse); ok {
		return map[string]string{
			"transactionId": pay.TransactionID,
			"id":            pay.Payments[0].ID,
			"status":        pay.Payments[0].Status,
			"message":       pay.Payments[0].Message,
		}, nil
	}

	return nil, errors.New("charge has failed")
}
