package communications

import "github.com/go-resty/resty/v2"

// Resty client to make HTTP interactions
type Resty struct {
	Client *resty.Client
}

// NewResty creates an instance of Resty
func NewResty() HttpClient {
	return &Resty{}
}

// Post requests with body and headers
func (r *Resty) Post(url string, body interface{}, headers map[string]string) ([]byte, error) {
	panic("implement me")
}
