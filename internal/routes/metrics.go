package routes

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
)

var httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "ecomm_http_requests_total",
	Help: "Count all HTTP requests by status code, method and path",
}, []string{"status_code", "method", "path"})

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "ecomm_http_request_duration",
	Help: "Duration in seconds of all HTTP requests",
}, []string{"path"})

// timedOperation gets how long a request/response took
type timedOperation struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

// WriteHeader tracks response status code
func (t *timedOperation) WriteHeader(statusCode int) {
	if t.wroteHeader {
		return
	}

	t.statusCode = statusCode
	t.ResponseWriter.WriteHeader(statusCode)
	t.wroteHeader = true
}

// metrics capture application internal state
func metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))

		monitor := timedOperation{
			ResponseWriter: w,
			statusCode:     0,
		}

		next.ServeHTTP(&monitor, r)

		httpRequestsTotal.WithLabelValues(strconv.Itoa(monitor.statusCode), r.Method, r.URL.Path).Inc()
		timer.ObserveDuration()
	})
}
