package routes

import (
	"context"
	communications "github.com/facilittei/ecomm/internal/communications/http"
	"github.com/facilittei/ecomm/internal/controllers"
	providers "github.com/facilittei/ecomm/internal/providers/juno"
	repositories "github.com/facilittei/ecomm/internal/repositories/auth"
	"github.com/facilittei/ecomm/internal/services"
	paymentSrv "github.com/facilittei/ecomm/internal/services/payments"
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
	ctx         context.Context
	router      *httprouter.Router
	redisClient *redis.Client
}

// NewApi creates an instance of Router
func NewApi(ctx context.Context, redisClient *redis.Client) *Api {
	return &Api{
		ctx:         ctx,
		router:      httprouter.New(),
		redisClient: redisClient,
	}
}

// Expose is the application entrypoint that exposes all endpoints
func (api *Api) Expose() http.Handler {
	healthcheckCtrl := controllers.NewHealthcheck(services.NewHealthcheck())

	httpClient := communications.NewRequester()
	junoProvider := providers.NewJuno(httpClient)
	authRepository := repositories.NewRedis(api.ctx, api.redisClient)
	paymentCtl := controllers.NewPayment(paymentSrv.NewJuno(junoProvider, authRepository))

	api.router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheckCtrl.Index)
	api.router.HandlerFunc(http.MethodPost, "/v1/payments/charge", paymentCtl.Charge)
	api.router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	return metrics(api.router)
}
