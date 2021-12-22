package providers

// junoError format error received from an authentication attempt
func junoError(message string) JunoError {
	return JunoError{
		Message: message,
	}
}
