package providers

import (
	"fmt"
	communications "github.com/facilittei/ecomm/internal/communications/http"
	"os"
)

const (
	envAuthUrl   = "JUNO_AUTHORIZATION_URL"
	envAuthBasic = "JUNO_AUTHORIZATION_BASIC"

	authUrlPath = "/authorization-server/oauth/token?grant_type=client_credentials"
)

// Juno payment gateway
type Juno struct {
	httpClient communications.HttpClient
}

// NewJuno creates an instance of Juno
func NewJuno(httpClient communications.HttpClient) *Juno {
	return &Juno{
		httpClient: httpClient,
	}
}

// Authenticate authenticates with Juno to get access token
// of which will be used on all requests
func (j *Juno) Authenticate() (*JunoAuth, error) {
	endpoint := os.Getenv(envAuthUrl)
	if endpoint == "" {
		return nil, authError(fmt.Sprintf("env required: %s", envAuthUrl))
	}

	authBasic := os.Getenv(envAuthBasic)
	if authBasic == "" {
		return nil, authError(fmt.Sprintf("env required: %s", envAuthBasic))
	}

	_, err := j.httpClient.Post(endpoint+authUrlPath, nil, map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": authBasic,
	})
	if err != nil {
		return nil, authError(fmt.Sprintf("http post request has failed: %v", err))
	}

	return nil, nil
}
