package controllers

import (
	"encoding/json"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/facilittei/ecomm/internal/logging"
	"github.com/facilittei/ecomm/internal/services/payments"
	"net/http"
)

// Payment requests for specific payment provider
type Payment struct {
	logger     logging.Logger
	PaymentSrv services.Payment
}

// NewPayment creates an instance of Payment
func NewPayment(logger logging.Logger, paymentSrv services.Payment) *Payment {
	return &Payment{
		logger:     logger,
		PaymentSrv: paymentSrv,
	}
}

// Charge customer for the desired product
func (p *Payment) Charge(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var paymentRequest payment.Request
	err := dec.Decode(&paymentRequest)
	if err != nil {
		p.logger.Error("payment request decode error [dec.Decode]: %v", err)
		if err = SendUnprocessableEntityJSON(w, Envelope{
			"status":  "failed",
			"message": http.StatusText(http.StatusUnprocessableEntity),
		}, nil); err != nil {
			p.logger.Error("payment charge response error [SendUnprocessableEntityJSON]: %v", err)
		}
		return
	}

	if errs := paymentRequest.Validate(); errs != nil {
		if err := SendUnprocessableEntityJSON(w, Envelope{
			"status":  "failed",
			"message": http.StatusText(http.StatusUnprocessableEntity),
			"errors":  DisplayErrors(errs),
		}, nil); err != nil {
			p.logger.Error("payment charge response error [SendUnprocessableEntityJSON]: %v", err)
		}
		return
	}

	charge, err := p.PaymentSrv.Charge(paymentRequest)
	if err != nil {
		p.logger.Error("payment service charge error: %v", err)
		if err := SendInternalErrorJSON(w, Envelope{
			"status":  "failed",
			"message": http.StatusText(http.StatusInternalServerError),
		}, nil); err != nil {
			p.logger.Error("payment charge response error [SendInternalErrorJSON]: %v", err)
		}
		return
	}

	if err := SendOkJSON(w, Envelope{"charge": charge}, nil); err != nil {
		p.logger.Error("payment charge response error [SendOkJSON]: %v", err)
	}
}
