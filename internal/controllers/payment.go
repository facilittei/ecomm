package controllers

import (
	"encoding/json"
	"github.com/facilittei/ecomm/internal/domains/payment"
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

// Charge customer for the desired product
func (p *Payment) Charge(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	w.Header().Set("Content-Type", "application/json")

	var paymentRequest payment.Request
	err := dec.Decode(&paymentRequest)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	charge, err := p.PaymentSrv.Charge()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(charge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}
