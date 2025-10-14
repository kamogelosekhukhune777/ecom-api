package mux

import (
	"net/http"

	"github.com/kamogelosekhukhune777/ecom-api/foundation/logger"
	"github.com/kamogelosekhukhune777/ecom-api/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, routeAdder RouteAdder) http.Handler {
	app := web.NewApp(cfg.Log.Info)

	routeAdder.Add(app, cfg)

	return app
}
