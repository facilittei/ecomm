package controllers

import (
	"github.com/facilittei/ecomm/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestChargeEndpoint(t *testing.T) {
	app := NewApp(config.Config{Port: "80"}).Routes()

	req, err := http.NewRequest(http.MethodPost, "/v1/payments/charge", nil)
	if err != nil {
		t.Errorf("error making HTTP request to /v1/payments/charge %v", err)
	}

	res, err := app.Test(req, -1)
	assert.Equalf(t, http.StatusOK, res.StatusCode, "payments charge route")
}
