package communications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Requester client to make HTTP interactions
type Requester struct{}

// NewRequester creates an instance of Requester
func NewRequester() HttpClient {
	return &Requester{}
}

// Post requests with body and headers
func (r *Requester) Post(url string, body interface{}, headers map[string]string) (res *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error trying to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error trying to create an request instance: %v", err)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	return client.Do(req)
}
