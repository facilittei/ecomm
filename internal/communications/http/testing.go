package communications

import (
	"net/http"
)

// serverTest helps to create a custom server for testing
type serverTest struct {
	mux *http.ServeMux
}

// ServeHTTP serves and expose endpoints with pre-defined responses
func (s *serverTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
