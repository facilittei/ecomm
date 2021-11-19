package servers

import (
	"errors"
	"fmt"
	"github.com/facilittei/ecomm/internal/config"
	"github.com/facilittei/ecomm/internal/routes"
	"github.com/facilittei/ecomm/internal/services"
	"github.com/gofiber/fiber/v2"
)

// App holds essential information
type App struct {
	Cfg    config.Config
	Router *fiber.App
}

// NewApp creates an instance of App
func NewApp(cfg config.Config) *App {
	return &App{
		Cfg:    cfg,
		Router: fiber.New(),
	}
}

// Routes register endpoints
func (app *App) Routes() *fiber.App {
	healthcheckSrv := services.NewHealthcheck()
	healthcheckRt := routes.NewHealthcheck(healthcheckSrv)

	v1 := app.Router.Group("/v1")
	v1.Get("/healthcheck", healthcheckRt.Index)

	return app.Router
}

// Listen HTTP requests
func (app *App) Listen() error {
	if app.Router == nil {
		return errors.New("router must be instantiated")
	}

	if app.Cfg.Port == "" {
		return errors.New("port is not set")
	}

	app.Routes()

	fmt.Printf("server listening on port %s", app.Cfg.Port)
	return app.Router.Listen(fmt.Sprintf(":%s", app.Cfg.Port))
}
