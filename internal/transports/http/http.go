//go:generate mockery --name=HttpClient --output ./../../mocks/ --filename http_client.go --structname HttpClientMock

package transports

import "net/http"

// HttpClient provides a clear interface of which methods could be performed
// when making HTTP requests
type HttpClient interface {
	Post(url string, body interface{}, headers map[string]string) (*http.Response, error)
}
