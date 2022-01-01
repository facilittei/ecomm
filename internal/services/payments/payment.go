//go:generate mockery --name=Payment --output ./../../mocks/ --filename payment_service.go --structname PaymentServiceMock

package services

import "github.com/facilittei/ecomm/internal/domains/payment"

// Payment provides clear interface for payment processing
type Payment interface {
	Charge(req payment.Request) (map[string]string, error)
}
