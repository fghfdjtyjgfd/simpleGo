package route

import (
	"fmt"
	"net/http"
	r "testapi/response"
	"time"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success("Hello, Docker! <3\n"))
	})

	e.GET("/name", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(
			fmt.Sprintf("hostname"),
		))
	})

	e.GET("/man", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(
			fmt.Sprintf("/, /man, /healthcheck, /db, /page"),
		))
	})

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r.Success(time.Now().String()))
	})

	e.GET("/db", func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, r.InternalServerError())
	})

	e.GET("/page", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, r.NotFound())
	})
}
