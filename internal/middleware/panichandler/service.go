package panichandler

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/writer"
	"github.com/labstack/echo/v4"
)

var (
	errSomethingWentWrong = "terjadi kesalahan, mohon coba beberapa saat lagi"
)

type errorDetail struct {
	Title string
	Value string
}

// HandlePanic handle panic and send error message to Discord.
func (middleware *Middleware) HandlePanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			defer func() {
				// Recover the panic if exists
				r := recover()
				if r == nil {
					return
				}

				// Get the body
				body, err := io.ReadAll(c.Request().Body)
				if err != nil {
					c.String(http.StatusInternalServerError, "Error reading request body")
					return
				}

				// Convert the body to a string
				bodyString := string(body)
				if bodyString == "" {
					bodyString = "-"
				}

				// Get the query params
				queryParams := c.QueryString()
				if queryParams == "" {
					queryParams = "-"
				}

				// Get the stack trace information
				stackTrace := make([]uintptr, 1)
				n := runtime.Callers(4, stackTrace[:])
				stackFrames := runtime.CallersFrames(stackTrace[:n])
				frame, _ := stackFrames.Next()

				// Create fields for the mattermost message
				fields := []errorDetail{
					{
						Title: "IP",
						Value: c.RealIP(),
					},
					{
						Title: "Method",
						Value: c.Request().Method,
					},
					{
						Title: "Path",
						Value: c.Path(),
					},
					{
						Title: "Function",
						Value: frame.Function,
					},
					{
						Title: "Line",
						Value: fmt.Sprintf("%v", frame.Line),
					},
					{
						Title: "File",
						Value: frame.File,
					},
					{
						Title: "Query Params",
						Value: queryParams,
					},
					{
						Title: "Body Payload",
						Value: bodyString,
					},
				}
				errMsg := fmt.Sprintf("%v", r)

				logger.Error(c.Request().Context(), fields, nil, errMsg)

				c.JSON(http.StatusBadGateway, writer.APIResponse(errSomethingWentWrong, false, nil))
			}()

			return next(c)
		}
	}
}
