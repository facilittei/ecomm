package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthcheckIndex(t *testing.T) {
	healthcheck := NewHealthcheck()

	want := make(map[string]string)
	want["status"] = "available"
	got := healthcheck.Index()

	assert.Equal(t, got, want)
}
