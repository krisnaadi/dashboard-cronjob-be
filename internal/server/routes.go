package server

import (
	app "github.com/krisnaadi/dashboard-cronjob-be/internal/app"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/config"

	echojwt "github.com/labstack/echo-jwt/v4"
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
		echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte(config.Get("JWT_KEY")),
		}),
	}

	// Register routes
	RegisterCronjobRoutes(router, groupMiddleware...)
	RegisterAuthRoutes(router)
	return router
}

func RegisterCronjobRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.GET("api/v1/jobs", router.Handler.Cronjob.HandleGetCronjob, m...)
	router.Echo.GET("api/v1/jobs/:id", router.Handler.Cronjob.HandleShowCronjob, m...)
	router.Echo.POST("api/v1/jobs", router.Handler.Cronjob.HandleCreateCronjob, m...)
	router.Echo.PUT("api/v1/jobs/:id", router.Handler.Cronjob.HandleEditCronjob, m...)
	router.Echo.DELETE("api/v1/jobs/:id", router.Handler.Cronjob.HandleEditCronjob, m...)
	router.Echo.POST("api/v1/jobs/:id/run", router.Handler.Cronjob.HandleRunCronjobManualy, m...)
	router.Echo.GET("api/v1/jobs/:id/logs", router.Handler.Cronjob.HandleGetLogByCronjob, m...)
}

func RegisterAuthRoutes(router *Router, m ...echo.MiddlewareFunc) {
	router.Echo.POST("api/v1/auth/login", router.Handler.Auth.HandleLogin, m...)
	router.Echo.POST("api/v1/auth/signup", router.Handler.Auth.HandleRegister, m...)
	router.Echo.GET("api/v1/auth/user", router.Handler.Auth.HandleShowUser, m...)
}
