package controllers

import (
	services "github.com/facilittei/ecomm/internal/services/payments"
	"github.com/gofiber/fiber/v2"
)

// Payment requests for specific payment provider
type Payment struct {
	PaymentSrv services.Payment
}

// NewPayment creates an instance of Payment
func NewPayment(paymentSrv services.Payment) *Payment {
	return &Payment{
		PaymentSrv: paymentSrv,
	}
}

// Charge customer for the desired products
func (p *Payment) Charge(ctx *fiber.Ctx) error {
	return ctx.JSON(p.PaymentSrv.Charge())
}
