package providers

import (
	"encoding/json"
	"fmt"
	transports "github.com/facilittei/ecomm/internal/transports/http"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	envAuthUrl   = "JUNO_AUTHORIZATION_URL"
	envAuthBasic = "JUNO_AUTHORIZATION_BASIC"

	authUrlPath = "/authorization-server/oauth/token?grant_type=client_credentials"
)

// Juno payment gateway
type Juno struct {
	httpClient transports.HttpClient
}

// NewJuno creates an instance of Juno
func NewJuno(httpClient transports.HttpClient) *Juno {
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

	res, err := j.httpClient.Post(endpoint+authUrlPath, nil, map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Basic %s", authBasic),
	})
	if err != nil {
		return nil, authError(fmt.Sprintf("http post request has failed: %v", err))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, authError(fmt.Sprintf("error trying read response: %v", err))
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("could not close request body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, authError(string(body))
	}

	auth := &JunoAuth{}
	err = json.Unmarshal(body, auth)
	if err != nil {
		return nil, authError(fmt.Sprintf("error trying unmarshal response: %v", err))
	}

	return auth, nil
}
