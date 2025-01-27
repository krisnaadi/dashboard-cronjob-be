package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gosimple/slug"
	_ "github.com/joho/godotenv/autoload"
	app "github.com/krisnaadi/dashboard-cronjob-be/internal/app"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/clocker"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/config"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/customvalidator"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	timeoutServer = 60
	port          = 8080
)

type Server struct {
	handler *app.Handlers
	http    *http.Server
}

func NewHTTP(ctx context.Context) *Server {

	db, err := postgresConnect()
	if err != nil {
		panic(err)
	}

	repository := app.NewRepository(db)
	resource := app.NewResource(repository)
	useCase := app.NewUseCase(resource)
	handler := app.NewHandler(useCase)

	logger.Init(slug.Make(config.Get("APP_NAME")))
	logger.Info(ctx, nil, nil, "Connecting - NewHTTP")

	return &Server{
		handler: handler,
	}
}

func (s *Server) Run() *http.Server {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			c.JSON(he.Code, nil)
			return
		}
	}

	//add cusstom validation
	customValidationMap := make(map[string]func(fl validator.FieldLevel) bool)
	customValidationMap["date"] = customvalidator.ValidateDateFormat

	customValidator := customvalidator.CustomValidaton(customValidationMap)

	e.Validator = customValidator

	// Allow CORS requests
	e.Use(middleware.CORS())

	e.GET("/", handleHelloWorld)
	NewRouter(e, s.handler)

	s.http = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      e,
		ReadTimeout:  timeoutServer * time.Second,
		WriteTimeout: timeoutServer * time.Second,
	}

	fmt.Printf("Server running on port %d\n", port)

	return s.http
}

func handleHelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World : "+clocker.Now().String())
}

func postgresConnect() (*gorm.DB, error) {
	dbConnection := config.Get("DB_GORM_CONNECTION")
	if dbConnection == "" {
		return nil, errors.New("can't connect to DB")
	}

	dsn := dbConnection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return clocker.Now()
		},
	})

	if err != nil {
		return db, err
	}

	fmt.Println("Connection to database established")
	return db, nil
}
