package server

import (
	app "github.com/krisnaadi/dashboard-cronjob-be/internal/app"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo    *echo.Echo
	Handler *app.Handlers
}

func NewRouter(e *echo.Echo, handler *app.Handlers) *Router {
	router := &Router{
		Echo:    e,
		Handler: handler,
	}

	// // Register global middleware
	// e.Use(middleware.LogRequest.LogRequest())
	// e.Use(middleware.PanicHandler.HandlePanic())
	// e.Use(middleware.HttpWrapper.HttpWrapper())

	// // Group middleware that usually used
	// groupMiddleware := []echo.MiddlewareFunc{
	// 	middleware.Signature.SignatureCheckMiddleware(),
	// 	middleware.Prometheus.MetricCollector(),
	// }

	// Register routes
	RegisterClientRoutes(router)
	return router
}

func RegisterClientRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.GET("api/v1/cronjob", router.Handler.Cronjob.HandleGetCronjob, m...)
}
