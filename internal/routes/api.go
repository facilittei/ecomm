package routes

import (
	"database/sql"
	"github.com/facilittei/ecomm/internal/controllers"
	"github.com/facilittei/ecomm/internal/logging"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	authRepository "github.com/facilittei/ecomm/internal/repositories/auth"
	chargeRepository "github.com/facilittei/ecomm/internal/repositories/charge"
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
	sqlClient   *sql.DB
	redisClient *redis.Client
	logger      logging.Logger
}

// NewApi creates an instance of Router
func NewApi(
	sqlClient *sql.DB,
	redisClient *redis.Client,
	logger logging.Logger,
) *Api {
	return &Api{
		router:      httprouter.New(),
		sqlClient:   sqlClient,
		redisClient: redisClient,
		logger:      logger,
	}
}

// Expose is the application entrypoint that exposes all endpoints
func (api *Api) Expose() http.Handler {
	healthcheckCtrl := controllers.NewHealthcheck(services.NewHealthcheck())

	httpClient := transports.NewRequester()
	junoProvider := providers.NewJuno(httpClient)
	authRepository := authRepository.NewRedis(api.redisClient)
	chargeRepository := chargeRepository.NewChargePsql(api.sqlClient)

	paySrv := paymentSrv.NewJuno(
		api.logger,
		junoProvider,
		authRepository,
		chargeRepository,
	)
	paymentCtl := controllers.NewPayment(api.logger, paySrv)

	api.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheckCtrl.Index)
	api.router.HandlerFunc(http.MethodPost, "/v1/payments/charge", paymentCtl.Charge)
	api.router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	return metrics(api.router)
}
