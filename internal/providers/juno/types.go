package providers

import (
	"fmt"
	"github.com/facilittei/ecomm/internal/domains/customer"
)

// JunoAuth holds response after an authenticated call
// This will be used to make sub-sequent calls to payment gateway
type JunoAuth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
	UserName    string `json:"user_name"`
	Jti         string `json:"jti"`
}

// JunoError is an error that occurred when trying to communicate with Juno
type JunoError struct {
	Message string
}

// Error raised by a communication attempt
func (e JunoError) Error() string {
	return fmt.Sprintf("could not communcation with Juno: %s", e.Message)
}

// JunoChargeRequest creates a new charge that'll be used later to make the payment
type JunoChargeRequest struct {
	Charge  JunoCharge  `json:"charge"`
	Billing JunoBilling `json:"billing"`
}

// JunoCharge has product-related information
type JunoCharge struct {
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	Methods     []string `json:"paymentTypes"`
}

// JunoBilling has payment-related information
type JunoBilling struct {
	Name     string           `json:"name"`
	Document string           `json:"document"`
	Email    string           `json:"email"`
	Address  customer.Address `json:"address"`
}

// JunoChargeResponse has charge ID that is required to make the payment
type JunoChargeResponse struct {
	Embedded struct {
		Charges []struct {
			ID string `json:"id"`
		} `json:"charges"`
	} `json:"_embedded"`
}

// JunoPayRequest makes the payment request for specified charge
type JunoPayRequest struct {
	ChargeID string `json:"chargeId"`
	Billing  struct {
		Email   string           `json:"email"`
		Address customer.Address `json:"address"`
	} `json:"billing"`
	CreditCard struct {
		CreditCardHash string `json:"creditCardHash"`
	} `json:"creditCardDetails"`
}

// JunoErrorResponse is returned when an error occur
type JunoErrorResponse struct {
	Timestamp string            `json:"timestamp"`
	Status    int32             `json:"status"`
	Error     string            `json:"error"`
	Details   []JunoErrorDetail `json:"details"`
}

// JunoErrorDetail has descriptive error information
type JunoErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"errorCode"`
}

// JunoPayResponse has payment status after payment processing
type JunoPayResponse struct {
	TransactionID string          `json:"transactionId"`
	Payments      []JunoPayDetail `json:"payments"`
}

// JunoPayDetail has descriptive payment information
type JunoPayDetail struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"failReason"`
}
