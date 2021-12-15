package controllers

import (
	"bytes"
	"context"
	"github.com/facilittei/ecomm/internal/mocks"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	"github.com/facilittei/ecomm/internal/services/payments"
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
	r := httptest.NewRequest(http.MethodGet, "/v1/payments/juno", payload)

	ctx := context.Background()
	httpClient := &mocks.HttpClientMock{}
	httpClient.On("Post").Return(r, nil)
	junoProvider := providers.NewJuno(httpClient)

	token := "1234abc"
	authRepository := &mocks.AuthRepositoryhMock{}
	authRepository.On("Get", ctx).Return(token, nil)
	authRepository.On("Store", ctx, token).Return(nil)

	junoSrv := services.NewJuno(junoProvider, authRepository)
	junoPaymentCtrl := NewPayment(junoSrv)
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
	r := httptest.NewRequest(http.MethodGet, "/v1/payments/juno", payload)

	ctx := context.Background()
	httpClient := &mocks.HttpClientMock{}
	httpClient.On("Post").Return(r, nil)
	junoProvider := providers.NewJuno(httpClient)

	token := "1234abc"
	authRepository := &mocks.AuthRepositoryhMock{}
	authRepository.On("Get", ctx).Return(token, nil)
	authRepository.On("Store", ctx, token).Return(nil)

	junoSrv := services.NewJuno(junoProvider, authRepository)
	junoPaymentCtrl := NewPayment(junoSrv)
	junoPaymentCtrl.Charge(w, r)

	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)

	got := string(body)
	want := `{"errors":["description is required","amount must be greater than 0","credit card hash is required"],"message":"Unprocessable Entity","status":"failed"}{"charge":{"status":"pending"}}`
	assert.Equal(t, want, got)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
}
