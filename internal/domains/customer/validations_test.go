package customer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomer_Validate(t *testing.T) {
	name := "Bill Gates"
	email := "bill@microsoft.com"
	document := "73379578070"
	address := Address{
		Street:   "Rua ABC",
		Number:   "1234",
		City:     "Recife",
		State:    "PE",
		PostCode: "51000000",
	}

	tests := []struct {
		name string
		args Customer
		want []error
	}{
		{
			name: "No error should returned",
			args: Customer{
				Name:     name,
				Email:    email,
				Document: document,
				Address:  address,
			},
			want: []error{},
		},
		{
			name: "Name, email, and document are missing",
			args: Customer{
				Address: address,
			},
			want: []error{ErrName, ErrEmail, ErrDocument},
		},
		{
			name: "Email is invalid",
			args: Customer{
				Name:     name,
				Email:    "bill$microsoft@",
				Document: document,
				Address:  address,
			},
			want: []error{ErrEmailInvalid},
		},
		{
			name: "Address street, number, city, state and post code are missing",
			args: Customer{
				Name:     name,
				Email:    email,
				Document: document,
				Address:  Address{},
			},
			want: []error{
				ErrAddressStreet,
				ErrAddressNumber,
				ErrAddressCity,
				ErrAddressState,
				ErrAddressPostCode,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.args.Validate()
			assert.Equal(t, tt.want, got)
			assert.Len(t, got, len(tt.want))
		})
	}
}

func TestIsCPF(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{args: "52998224725", want: true},
		{args: "73379578070", want: true},
		{args: "38553699099", want: true},
		{args: "41051378087", want: true},
		{args: "65746868060", want: true},
		{args: "65746868060", want: true},
		{args: "00011122233", want: false},
		{args: "00000000000", want: false},
		{args: "12345678912", want: false},
	}

	for _, tt := range tests {
		got := isCPF(tt.args)
		assert.Equal(t, tt.want, got)
	}
}
