package services

// Healthcheck exposes application operational status
type Healthcheck struct{}

// NewHealthcheck creates an instance of Healthcheck
func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

// Index returns system status info
func (h *Healthcheck) Index() map[string]string {
	return map[string]string{
		"status": "available",
	}
}
