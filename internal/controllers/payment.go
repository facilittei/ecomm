package controllers

import (
	"github.com/facilittei/ecomm/internal/services/payments"
	"net/http"
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
func (p *Payment) Charge(w http.ResponseWriter, r *http.Request) {
	charge, err := p.PaymentSrv.Charge()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(charge["status"]))
}
