package payment

import "errors"

var (
	ErrDescription    = errors.New("description is required")
	ErrAmount         = errors.New("amount must be greater than 0")
	ErrCreditCardHash = errors.New("credit card hash is required")
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
