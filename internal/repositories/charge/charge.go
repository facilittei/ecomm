//go:generate mockery --name=Charge --output ./../../mocks/ --filename charge_repository.go --structname ChargeRepositoryMock

package repository

import (
	"context"
	"github.com/facilittei/ecomm/internal/domains/payment"
)

// Charge holds requests for charges history by tracking its states
// during communication with payment providers
type Charge interface {
	Store(ctx context.Context, charge payment.Charge) error
}
