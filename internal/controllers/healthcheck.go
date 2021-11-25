package controllers

import (
	"github.com/facilittei/ecomm/internal/services"
	"net/http"
)

// Healthcheck resource-related requests
type Healthcheck struct {
	HealthcheckSrv *services.Healthcheck
}

// NewHealthcheck creates an instance of Healthcheck
func NewHealthcheck(healthcheckSrv *services.Healthcheck) *Healthcheck {
	return &Healthcheck{
		HealthcheckSrv: healthcheckSrv,
	}
}

// Index returns system status info
func (h *Healthcheck) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.HealthcheckSrv.Index()["status"]))
}
