package payment

import (
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequest_Validate(t *testing.T) {
	customer := customer.Customer{
		Name:     "Bill Gates",
		Email:    "bill@microsoft.com",
		Document: "00011122233",
		Address: customer.Address{
			Street:   "Rua ABC",
			Number:   "1234",
			City:     "Recife",
			State:    "PE",
			PostCode: "51000000",
		},
	}
	creditCard := CreditCard{Hash: "1234abc"}

	tests := []struct {
		name string
		args Request
		want []error
	}{
		{
			name: "Description is missing",
			args: Request{
				Amount:     10,
				Customer:   customer,
				CreditCard: creditCard,
			},
			want: []error{ErrDescription},
		},
		{
			name: "Amount is missing",
			args: Request{
				Description: "My awesome product",
				Customer:    customer,
				CreditCard:  creditCard,
			},
			want: []error{ErrAmount},
		},
		{
			name: "Amount is less than 0",
			args: Request{
				Description: "My awesome product",
				Amount:      -10,
				Customer:    customer,
				CreditCard:  creditCard,
			},
			want: []error{ErrAmount},
		},
		{
			name: "Description, amount and credit card hash are missing",
			args: Request{
				Description: "",
				Amount:      0,
				Customer:    customer,
				CreditCard:  CreditCard{},
			},
			want: []error{
				ErrDescription,
				ErrAmount,
				ErrCreditCardHash,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Validate()
			assert.Equal(t, tt.want, got)
			assert.Len(t, got, len(tt.want))
		})
	}
}
