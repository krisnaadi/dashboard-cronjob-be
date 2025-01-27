package logrequest

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (middleware *Middleware) LogRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(c.Path())
			return next(c)
		}
	}
}
