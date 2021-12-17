package routes

import (
	"github.com/facilittei/ecomm/internal/controllers"
	"github.com/facilittei/ecomm/internal/logging"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	repositories "github.com/facilittei/ecomm/internal/repositories/auth"
	"github.com/facilittei/ecomm/internal/services"
	paymentSrv "github.com/facilittei/ecomm/internal/services/payments"
	transports "github.com/facilittei/ecomm/internal/transports/http"
	"github.com/go-redis/redis/v8"
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
	router      *httprouter.Router
	redisClient *redis.Client
	logger      logging.Logger
}

// NewApi creates an instance of Router
func NewApi(logger logging.Logger, redisClient *redis.Client) *Api {
	return &Api{
		logger:      logger,
		redisClient: redisClient,
		router:      httprouter.New(),
	}
}

// Expose is the application entrypoint that exposes all endpoints
func (api *Api) Expose() http.Handler {
	healthcheckCtrl := controllers.NewHealthcheck(services.NewHealthcheck())

	httpClient := transports.NewRequester()
	junoProvider := providers.NewJuno(httpClient)
	authRepository := repositories.NewRedis(api.redisClient)

	paySrv := paymentSrv.NewJuno(api.logger, junoProvider, authRepository)
	paymentCtl := controllers.NewPayment(api.logger, paySrv)

	api.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheckCtrl.Index)
	api.router.HandlerFunc(http.MethodPost, "/v1/payments/charge", paymentCtl.Charge)
	api.router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	return metrics(api.router)
}
