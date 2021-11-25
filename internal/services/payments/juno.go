package services

// Juno handles payment transaction requests
type Juno struct{}

// NewJuno creates an instance of Juno
func NewJuno() *Juno {
	return &Juno{}
}

// Charge customer using Juno payment provider
func (j *Juno) Charge() map[string]string {
	return map[string]string{
		"status": "pending",
	}
}
