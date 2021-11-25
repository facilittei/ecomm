package controllers

import (
	"github.com/facilittei/ecomm/internal/services/payments"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJunoChargeEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/v1/payments/juno", nil)
	if err != nil {
		t.Fatal(err)
	}

	junoPaymentCtrl := NewPayment(services.NewJuno())
	junoPaymentCtrl.Charge(w, r)

	res := w.Result()
	assert.Equalf(t, http.StatusOK, res.StatusCode, "juno route")
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "pending"
	got := string(body)
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
