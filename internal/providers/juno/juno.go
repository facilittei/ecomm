package providers

import (
	"encoding/json"
	"fmt"
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/facilittei/ecomm/internal/domains/payment"
	transports "github.com/facilittei/ecomm/internal/transports/http"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	envApiUrl           = "JUNO_API_URL"
	envAuthBasic        = "JUNO_AUTHORIZATION_BASIC"
	envApiVersion       = "JUNO_API_VERSION"
	envApiResourceToken = "JUNO_API_RESOURCE_TOKEN"

	authUrlPath = "/authorization-server/oauth/token?grant_type=client_credentials"
	apiUrlPath  = "/api-integration"
)

// Juno payment gateway
type Juno struct {
	auth       *JunoAuth
	httpClient transports.HttpClient
}

// NewJuno creates an instance of Juno
func NewJuno(httpClient transports.HttpClient) *Juno {
	return &Juno{
		auth:       &JunoAuth{},
		httpClient: httpClient,
	}
}

// Authenticate authenticates with Juno to get access token
// of which will be used on all requests
func (j *Juno) Authenticate() (*JunoAuth, error) {
	endpoint := os.Getenv(envApiUrl)
	if endpoint == "" {
		return nil, junoError(fmt.Sprintf("env required: %s", envApiUrl))
	}

	authBasic := os.Getenv(envAuthBasic)
	if authBasic == "" {
		return nil, junoError(fmt.Sprintf("env required: %s", envAuthBasic))
	}

	res, err := j.httpClient.Post(endpoint+authUrlPath, nil, map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Basic %s", authBasic),
	})
	if err != nil {
		return nil, junoError(fmt.Sprintf("http post request has failed: %v", err))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, junoError(fmt.Sprintf("error trying read response: %v", err))
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("could not close request body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, junoError(string(body))
	}

	err = json.Unmarshal(body, j.auth)
	if err != nil {
		return nil, junoError(fmt.Sprintf("error trying unmarshal response: %v", err))
	}

	return j.auth, nil
}

// Charge makes a new charge request to Juno
func (j *Juno) Charge(req payment.Request) (interface{}, error) {
	endpoint := os.Getenv(envApiUrl)
	if endpoint == "" {
		return nil, junoError(fmt.Sprintf("env required: %s", envApiUrl))
	}

	headers, err := j.requestHeader()
	if err != nil {
		return nil, junoError(fmt.Sprintf("env required: %s", err))
	}

	chargeResponse, err := j.makeChargeRequest(req, endpoint, headers)
	if err != nil {
		return nil, err
	}

	payResContent, err := j.makePayRequest(req, chargeResponse, endpoint, headers)
	if err != nil {
		return nil, err
	}
	return payResContent, nil
}

// requestHeader sets required header fields to be sent
func (j *Juno) requestHeader() (map[string]string, error) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	apiVersion := os.Getenv(envApiVersion)
	if apiVersion == "" {
		return nil, junoError(fmt.Sprintf("env required: %s", envApiVersion))
	}
	headers["X-Api-Version"] = apiVersion

	apiResourceToken := os.Getenv(envApiResourceToken)
	if apiResourceToken == "" {
		return nil, junoError(fmt.Sprintf("env required: %s", envApiResourceToken))
	}
	headers["X-Resource-Token"] = apiResourceToken

	if j.auth.AccessToken == "" {
		return nil, junoError("authorization bearer token is missing")
	}
	headers["Authorization"] = fmt.Sprintf("Bearer %s", j.auth.AccessToken)

	return headers, nil
}

// makeChargeRequest sends a create charge request
func (j *Juno) makeChargeRequest(
	req payment.Request,
	endpoint string,
	headers map[string]string,
) (*JunoChargeResponse, error) {
	chargeReq := &JunoChargeRequest{
		Charge: JunoCharge{
			Description: req.Description,
			Amount:      req.Amount,
			Methods:     []string{"CREDIT_CARD"},
		},
		Billing: JunoBilling{
			Name:     req.Customer.Name,
			Email:    req.Customer.Email,
			Document: req.Customer.Document,
			Address:  req.Customer.Address,
		},
	}

	res, err := j.httpClient.Post(endpoint+apiUrlPath+"/charges", chargeReq, headers)
	if err != nil {
		return nil, junoError(fmt.Sprintf("create charge request failed: %s", err))
	}
	defer res.Body.Close()

	chargeResContent, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, junoError(fmt.Sprintf("create charge response error [io.ReadAll]: %s", err))
	}

	chargeResponse := &JunoChargeResponse{}
	err = json.Unmarshal(chargeResContent, chargeResponse)
	if err != nil {
		return nil, junoError(fmt.Sprintf("create charge response error [json.Unmarshal]: %s", err))
	}

	if len(chargeResponse.Embedded.Charges) < 1 {
		return nil, junoError("create charge didn't returned charge ID")
	}

	return chargeResponse, nil
}

// makePayRequest sends a payment request
func (j *Juno) makePayRequest(
	req payment.Request,
	chargeResponse *JunoChargeResponse,
	endpoint string,
	headers map[string]string,
) (*JunoPayResponse, error) {
	payReq := &JunoPayRequest{
		ChargeID: chargeResponse.Embedded.Charges[0].ID,
		Billing: struct {
			Email   string           `json:"email"`
			Address customer.Address `json:"address"`
		}{
			Email:   req.Customer.Email,
			Address: req.Customer.Address,
		},
		CreditCard: struct {
			CreditCardHash string `json:"creditCardHash"`
		}{
			CreditCardHash: req.CreditCard.Hash,
		},
	}

	res, err := j.httpClient.Post(endpoint+apiUrlPath+"/payments", payReq, headers)
	if err != nil {
		return nil, junoError(fmt.Sprintf("pay charge request failed: %s", err))
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, junoError(fmt.Sprintf("pay charge response error [io.ReadAll]: %s", err))
	}

	if res.StatusCode != http.StatusOK {
		resError := &JunoErrorResponse{}
		err = json.Unmarshal(content, resError)
		if err != nil {
			return nil, junoError(fmt.Sprintf("error trying unmarshal response: %v", err))
		}

		return &JunoPayResponse{
			Payments: []JunoPayDetail{
				{Status: "FAILED", Message: resError.Details[0].Message},
			},
		}, nil
	}

	payRes := &JunoPayResponse{}
	err = json.Unmarshal(content, payRes)
	if err != nil {
		return nil, junoError(fmt.Sprintf("error trying unmarshal response: %v", err))
	}

	return payRes, nil
}

// UseToken for requests related to the resource server
func (j *Juno) UseToken(token string) {
	j.auth.AccessToken = token
}
