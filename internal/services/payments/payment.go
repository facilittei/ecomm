package services

// Payment provides clear interface for payment processing
type Payment interface {
	Charge() (map[string]string, error)
}
