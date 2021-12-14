package payment

import "github.com/facilittei/ecomm/internal/domains/customer"

// Request receives a payment intention for buying a product
// with all necessary information: customer, address, product and payment method
type Request struct {
	Description string            `json:"description"`
	Amount      float64           `json:"amount"`
	Customer    customer.Customer `json:"customer"`
	CreditCard  CreditCard        `json:"creditCard"`
}

// CreditCard holds credit card hash generated by payment provider
type CreditCard struct {
	Hash string `json:"hash"`
}