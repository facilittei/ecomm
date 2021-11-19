package servers

import (
	"github.com/facilittei/ecomm/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAppWithoutRouterInstance(t *testing.T) {
	app := &App{}
	if err := app.Listen(); err == nil {
		t.Error("should return error when there isn't a router instantiated")
	}
}

func TestHealthcheckEndpoint(t *testing.T) {
	app := NewApp(config.Config{Port: "80"}).Routes()

	req, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Errorf("error making HTTP request to /v1/healthcheck %v", err)
	}

	res, err := app.Test(req, -1)
	assert.Equalf(t, http.StatusOK, res.StatusCode, "healthcheck route")
}
