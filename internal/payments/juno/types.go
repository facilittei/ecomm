package payments

// JunoAuthenticateRequest makes a request to authorize the application
// with credentials provided by the payment gateway
type JunoAuthenticateRequest struct {
	ContentType   string
	Authorization string
}

// JunoAuthenticateResponse holds response after an authenticated call
// This will be used to make sub-sequent calls to payment gateway
type JunoAuthenticateResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
	UserName    string `json:"user_name"`
	Jti         string `json:"jti"`
}
