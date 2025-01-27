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
	RegisterCronjobRoutes(router)
	RegisterAuthRoutes(router)
	return router
}

func RegisterCronjobRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.GET("api/v1/jobs", router.Handler.Cronjob.HandleGetCronjob, m...)
	router.Echo.GET("api/v1/jobs/:id", router.Handler.Cronjob.HandleShowCronjob, m...)
	router.Echo.POST("api/v1/jobs", router.Handler.Cronjob.HandleCreateCronjob, m...)
	router.Echo.PUT("api/v1/jobs/:id", router.Handler.Cronjob.HandleEditCronjob, m...)
	router.Echo.DELETE("api/v1/jobs/:id", router.Handler.Cronjob.HandleEditCronjob, m...)
}

func RegisterAuthRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.POST("api/v1/auth/login", router.Handler.Auth.HandleLogin, m...)
	router.Echo.POST("api/v1/auth/signup", router.Handler.Auth.HandleRegister, m...)
	router.Echo.GET("api/v1/auth/user", router.Handler.Auth.HandleShowUser, m...)
}
