package controllers

import (
	"github.com/facilittei/ecomm/internal/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheckEndpoint(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := NewHealthcheck(services.NewHealthcheck())
	h.Index(w, r)

	res := w.Result()
	assert.Equalf(t, http.StatusOK, res.StatusCode, "healthcheck route")
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "available"
	got := string(body)
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
