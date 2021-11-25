package routes

import (
	"github.com/facilittei/ecomm/internal/controllers"
	"github.com/facilittei/ecomm/internal/services"
	paymentSrv "github.com/facilittei/ecomm/internal/services/payments"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// init settings
func init() {
	prometheus.Register(httpRequestsTotal)
	prometheus.Register(httpDuration)
}

// Api wraps http router for handler compliance
type Api struct {
	router *httprouter.Router
}

// NewApi creates an instance of Router
func NewApi() *Api {
	return &Api{
		router: httprouter.New(),
	}
}

// Expose endpoints
func (api *Api) Expose() http.Handler {
	healthcheckCtrl := controllers.NewHealthcheck(services.NewHealthcheck())
	paymentCtl := controllers.NewPayment(paymentSrv.NewJuno())

	api.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheckCtrl.Index)
	api.router.HandlerFunc(http.MethodGet, "/v1/payments/charge", paymentCtl.Charge)
	api.router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	return metrics(api.router)
}
