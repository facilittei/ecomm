package controllers

import (
	"encoding/json"
	"net/http"
)

// Envelope wraps a content to be sent as a response
type Envelope map[string]interface{}

// SendJSON writes response back for client
func SendJSON(w http.ResponseWriter, content Envelope, status int, headers http.Header) error {
	jsonContent, err := json.Marshal(content)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonContent)

	return nil
}

// SendOkJSON writes a 200 response back for the client
func SendOkJSON(w http.ResponseWriter, content Envelope, headers http.Header) error {
	return SendJSON(w, content, http.StatusOK, headers)
}

// SendInternalErrorJSON writes a 500 response back for the client
func SendInternalErrorJSON(w http.ResponseWriter, content Envelope, headers http.Header) error {
	return SendJSON(w, content, http.StatusInternalServerError, headers)
}

// SendUnprocessableEntityJSON writes a 422 response back for the client
func SendUnprocessableEntityJSON(w http.ResponseWriter, content Envelope, headers http.Header) error {
	return SendJSON(w, content, http.StatusUnprocessableEntity, headers)
}

// DisplayErrors shows the error message
func DisplayErrors(errs []error) []string {
	var res []string
	for _, err := range errs {
		res = append(res, err.Error())
	}
	return res
}
