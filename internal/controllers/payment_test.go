package controllers

import (
	"bytes"
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/facilittei/ecomm/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestJunoChargeEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	payload := bytes.NewReader([]byte(`{
		"description": "My awesome product",
		"amount": 0.99,
		"customer": {
			"name": "Jeff Bezos",
			"email": "jeff@amazon.com",
			"document": "11144740452",
			"address": {
				"street": "Rua Guedes Perreira",
				"number": "90",
				"city": "Recife",
				"state": "PE",
				"postCode": "52060150"
			}
		},
		"creditCard": {
			"hash": "cbdbb8f4-53a9-42e8-a077-83af06823e85"
		}
	}`))
	r := httptest.NewRequest(http.MethodGet, "/v1/payments", payload)

	logger := &mocks.LoggerMock{}
	junoSrv := &mocks.PaymentServiceMock{}
	junoSrv.On("Charge", payment.Request{
		Description: "My awesome product",
		Amount:      0.99,
		Customer: customer.Customer{
			Name:     "Jeff Bezos",
			Email:    "jeff@amazon.com",
			Document: "11144740452",
			Address: customer.Address{
				Street:   "Rua Guedes Perreira",
				Number:   "90",
				City:     "Recife",
				State:    "PE",
				PostCode: "52060150",
			},
		},
		CreditCard: payment.CreditCard{
			Hash: "cbdbb8f4-53a9-42e8-a077-83af06823e85",
		},
	}).Return(map[string]string{
		"transactionId": "123abc",
		"id":            "abc123",
		"status":        "CONFIRMED",
		"message":       "",
	}, nil)
	junoPaymentCtrl := NewPayment(logger, junoSrv)
	junoPaymentCtrl.Charge(w, r)

	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)

	rgx, err := regexp.Compile("status")
	require.Empty(t, err)

	got := string(body)
	assert.Regexp(t, rgx, got)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestJunoChargeEndpointMissingFields(t *testing.T) {
	w := httptest.NewRecorder()
	payload := bytes.NewReader([]byte(`{
		"customer": {
			"name": "Jeff Bezos",
			"email": "jeff@amazon.com",
			"document": "11144740452",
			"address": {
				"street": "Rua Guedes Perreira",
				"number": "90",
				"city": "Recife",
				"state": "PE",
				"postCode": "52060150"
			}
		}
	}`))
	r := httptest.NewRequest(http.MethodGet, "/v1/payments", payload)

	logger := &mocks.LoggerMock{}
	junoSrv := &mocks.PaymentServiceMock{}
	junoSrv.On("Charge", payment.Request{
		Customer: customer.Customer{
			Name:     "Jeff Bezos",
			Email:    "jeff@amazon.com",
			Document: "11144740452",
			Address: customer.Address{
				Street:   "Rua Guedes Perreira",
				Number:   "90",
				City:     "Recife",
				State:    "PE",
				PostCode: "52060150",
			},
		},
	}).Return(nil, nil)
	junoPaymentCtrl := NewPayment(logger, junoSrv)
	junoPaymentCtrl.Charge(w, r)

	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)

	got := string(body)
	want := `{"errors":["description is required","amount must be greater than 0","credit card hash is required"],"message":"Unprocessable Entity","status":"failed"}`
	assert.Equal(t, want, got)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}
