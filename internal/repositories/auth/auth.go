package repositories

// Auth token is stored in a database for reuse
type Auth interface {
	Store(token string) error
	Get() (string, error)
}
