package repositories

import "fmt"

// AuthError is an error that occurred when
// trying to get an auth token from repository
type AuthError struct {
	Message string
}

// Error raised by an auth repository
func (e AuthError) Error() string {
	return fmt.Sprintf("auth repository resulted in: %s", e.Message)
}
