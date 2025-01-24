package server

import (
	app "github.com/krisnaadi/dashboard-cronjob-be/internal/app"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo       *echo.Echo
	Handler    *app.Handlers
	Middleware *app.Middleware
}

func NewRouter(e *echo.Echo, handler *app.Handlers, middleware *app.Middleware) *Router {
	router := &Router{
		Echo:       e,
		Handler:    handler,
		Middleware: middleware,
	}

	// Register global middleware
	e.Use(middleware.LogRequest.LogRequest())
	e.Use(middleware.PanicHandler.HandlePanic())
	e.Use(middleware.HttpWrapper.HttpWrapper())

	// Group middleware that usually used
	groupMiddleware := []echo.MiddlewareFunc{
		middleware.Signature.SignatureCheckMiddleware(),
		middleware.Prometheus.MetricCollector(),
	}

	// Register routes
	RegisterClientRoutes(router, groupMiddleware...)
	return router
}

func RegisterClientRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.GET("api/v1/client", router.Handler.Client.HandleShowClient, m...)
}
