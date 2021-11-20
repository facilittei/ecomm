package payments

import communications "github.com/facilittei/ecomm/internal/communications/http"

// Juno payment gateway
type Juno struct {
	URL        string                  `json:"url"`
	Auth       JunoAuthenticateRequest `json:"-"`
	httpClient communications.HttpClient
}

// Authenticate with provider to exchange credentials for a token
func (j *Juno) Authenticate() *JunoAuthenticateResponse {
	return nil
}
