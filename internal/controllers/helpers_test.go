package controllers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendJSON(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	err := SendJSON(w, Envelope{"status": "OK"}, http.StatusOK, nil)
	require.Empty(t, err)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)
	got := string(body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, `{"status":"OK"}`, got)
}

func TestSendOkJSON(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	err := SendOkJSON(w, Envelope{"status": "OK"}, nil)
	require.Empty(t, err)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)
	got := string(body)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, `{"status":"OK"}`, got)
}

func TestSendInternalErrorJSON(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	err := SendInternalErrorJSON(w, Envelope{"status": "Internal Server Error"}, nil)
	require.Empty(t, err)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)
	got := string(body)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Equal(t, `{"status":"Internal Server Error"}`, got)
}

func TestSendUnprocessableEntityJSON(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	err := SendUnprocessableEntityJSON(w, Envelope{"status": "Unprocessable Entity"}, nil)
	require.Empty(t, err)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.Empty(t, err)
	got := string(body)
	assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	assert.Equal(t, `{"status":"Unprocessable Entity"}`, got)
}

func TestDisplayErrors(t *testing.T) {
	errs := []error{
		errors.New("my"),
		errors.New("displayed"),
		errors.New("errors"),
	}

	got := DisplayErrors(errs)
	want := "my displayed errors"
	assert.Equal(t, want, strings.Join(got, " "))
}
