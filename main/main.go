package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tahamazari/echo-framework/pkg/config"
	"github.com/tahamazari/echo-framework/pkg/routes"
)

func main() {
	config.DatabaseInit()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	routes.SetupBookRoutes(e)
	e.Logger.Fatal(e.Start(":4000"))
}
