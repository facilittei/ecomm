package providers

import "fmt"

// JunoAuth holds response after an authenticated call
// This will be used to make sub-sequent calls to payment gateway
type JunoAuth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
	UserName    string `json:"user_name"`
	Jti         string `json:"jti"`
}

// JunoAuthError is an error that occurred when trying to authenticate
type JunoAuthError struct {
	Message string
}

// Error raised by an authentication attempt
func (e JunoAuthError) Error() string {
	return fmt.Sprintf("could not authenticate: %s", e.Message)
}
