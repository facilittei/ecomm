package payment

import (
	"errors"
)

var (
	ErrDescription    = errors.New("description is required")
	ErrAmount         = errors.New("amount must be greater than 0")
	ErrCreditCardHash = errors.New("credit card hash is required")
	ErrChargeID       = errors.New("charge ID is required")
)

// Validate checks whether the Request payload
// has all required fields for making a payment transaction
// this means personal details, address, product and payment method
func (req *Request) Validate() []error {
	var errs []error

	if req.Description == "" {
		errs = append(errs, ErrDescription)
	}

	if req.Amount <= 0 {
		errs = append(errs, ErrAmount)
	}

	if err := req.Customer.Validate(); err != nil {
		errs = append(errs, err...)
	}

	if err := req.CreditCard.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

// Validate checks whether the CreditCard has required fields
func (req *CreditCard) Validate() []error {
	var errs []error

	if req.Hash == "" {
		errs = append(errs, ErrCreditCardHash)
	}

	return errs
}

// Validate checks whether the Charge has required fields
func (c Charge) Validate() []error {
	var errs []error

	if c.ID.String() == "00000000-0000-0000-0000-000000000000" {
		errs = append(errs, ErrChargeID)
	}

	if err := c.Customer.Validate(); err != nil {
		errs = append(errs, err...)
	}

	return errs
}
