package communications

// HttpClient provides a clear interface of which methods could be performed
// when making HTTP requests
type HttpClient interface {
	Post(url string, body interface{}, headers map[string]string) ([]byte, error)
}
