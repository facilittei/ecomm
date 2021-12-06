package providers

// authError format error received from an authentication attempt
func authError(message string) JunoAuthError {
	return JunoAuthError{
		Message: message,
	}
}