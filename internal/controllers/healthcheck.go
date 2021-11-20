package controllers

import (
	"github.com/facilittei/ecomm/internal/services"
	"github.com/gofiber/fiber/v2"
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
func (h *Healthcheck) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(h.HealthcheckSrv.Index())
}
