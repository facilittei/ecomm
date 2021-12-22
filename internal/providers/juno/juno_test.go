package providers

import (
	"bytes"
	"errors"
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/facilittei/ecomm/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestJuno_Authenticate_env_required(t *testing.T) {
	tests := []struct {
		name string
		args map[string]string
	}{
		{
			name: "JUNO_API_URL is missing",
			args: map[string]string{"JUNO_API_URL": ""},
		},
		{
			name: "JUNO_AUTHORIZATION_BASIC is missing",
			args: map[string]string{
				"JUNO_API_URL":             "https://url.test",
				"JUNO_AUTHORIZATION_BASIC": "",
			},
		},
	}

	httpClient := &mocks.HttpClientMock{}
	httpClient.On("Post").Return(nil, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args {
				err := os.Setenv(k, v)
				if err != nil {
					t.Fatalf("os.Setenv(%s, %s) error: %v", k, v, err)
				}

				junoProvider := NewJuno(httpClient)
				_, got := junoProvider.Authenticate()
				assert.True(t, errors.As(got, &JunoError{}))
			}
		})
	}
}

func TestJuno_Charge(t *testing.T) {
	if err := os.Setenv(envApiUrl, "https://secure.pay"); err != nil {
		t.Fatalf("os.Setenv(%s, %s) error: %v", envApiUrl, "https://secure.pay", err)
	}

	if err := os.Setenv(envApiVersion, "2"); err != nil {
		t.Fatalf("os.Setenv(%s, %s) error: %v", envApiVersion, "2", err)
	}

	if err := os.Setenv(envApiResourceToken, "1234abc"); err != nil {
		t.Fatalf("os.Setenv(%s, %s) error: %v", envApiResourceToken, "1234abc", err)
	}

	req := payment.Request{
		Description: "Super new course",
		Amount:      10,
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
			Hash: "eb3cb818-28bf-41db-8370-0cc1c9fabc67",
		},
	}

	httpClient := &mocks.HttpClientMock{}

	chargeResBody := bytes.NewReader([]byte(`{
		"_embedded": {
			"charges": [
			  {
				"id": "chr_F93724733AB07AF48EA63853CB7210AA",
				"code": 136708989,
				"reference": "",
				"dueDate": "2021-09-18",
				"checkoutUrl": "https://pay-sandbox.juno.com.br/checkout/84E4A237CF1A30999AA4117434DE017CC6BFD32A82B18BAC",
				"amount": 0.52,
				"status": "ACTIVE",
				"_links": {
				  "self": {
					"href": "https://sandbox.boletobancario.com/api-integration/charges/chr_F93724733AB07AF48EA63853CB7210AA"
				  }
				}
			  }
			]
		  }
	}`))
	chargeRes := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(chargeResBody),
		Close:      true,
	}

	payResBody := bytes.NewReader([]byte(`{
	  "transactionId": "309edf109ced8b",
	  "installments": 1,
	  "payments": [
		{
		  "id": "pay_B8D2E9B1FD04C040B5EDA951F4F7B1D1",
		  "chargeId": "chr_9768455F07BFCFC64083A1CB61C960B6",
		  "createdOn": "2021-12-21 21:02:06",
		  "updatedOn": "2021-12-21 21:02:06",
		  "date": "2021-12-21",
		  "releaseDate": "2022-01-22",
		  "amount": 0.52,
		  "fee": 0.52,
		  "type": "CREDIT_CARD",
		  "status": "CONFIRMED",
		  "failReason": null
		}
	  ]
	}`))
	payRes := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(payResBody),
		Close:      true,
	}

	junoProvider := NewJuno(httpClient)
	junoProvider.UseToken("abcdef123")

	headers, err := junoProvider.requestHeader()
	require.Nil(t, err)

	httpClient.On("Post",
		mock.AnythingOfType("string"),
		&JunoChargeRequest{
			Charge: JunoCharge{
				Description: req.Description,
				Amount:      req.Amount,
				Methods:     []string{"CREDIT_CARD"},
			},
			Billing: JunoBilling{
				Name:     req.Customer.Name,
				Document: req.Customer.Document,
				Email:    req.Customer.Email,
				Address:  req.Customer.Address,
			},
		},
		headers,
	).Return(chargeRes, nil).Once()

	httpClient.On("Post",
		mock.AnythingOfType("string"),
		&JunoPayRequest{
			ChargeID: "chr_F93724733AB07AF48EA63853CB7210AA",
			Billing: struct {
				Email   string           `json:"email"`
				Address customer.Address `json:"address"`
			}{
				Email:   req.Customer.Email,
				Address: req.Customer.Address,
			},
			CreditCard: struct {
				CreditCardHash string `json:"creditCardHash"`
			}{
				CreditCardHash: req.CreditCard.Hash,
			},
		},
		headers,
	).Return(payRes, nil).Once()

	res, err := junoProvider.Charge(req)
	assert.Nil(t, err)

	pay, ok := res.(*JunoPayResponse)
	assert.True(t, ok)
	assert.Equal(t, "309edf109ced8b", pay.TransactionID)
	assert.Equal(t, "pay_B8D2E9B1FD04C040B5EDA951F4F7B1D1", pay.Payments[0].ID)
	assert.Equal(t, "CONFIRMED", pay.Payments[0].Status)
}
