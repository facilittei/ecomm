package transports

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequester_Post(t *testing.T) {
	payload := []byte(`{"id":"1234abc","amount":0.99}`)

	server := &serverTest{}
	server.mux = http.NewServeMux()
	server.mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(payload)
	})

	serverTest := httptest.NewServer(server)

	defer serverTest.Close()
	requester := NewRequester()
	res, err := requester.Post(serverTest.URL+"/payments", map[string]interface{}{
		"name":        "Test User",
		"email":       "test@facilittei.com",
		"description": "Golang API",
		"amount":      "0.99",
	}, nil)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error trying read response: %v", err)
	}

	assert.Nil(t, err)
	assert.Equal(t, payload, body)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}
