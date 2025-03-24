package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middlewares.
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes.
	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server.
	e.Logger.Fatal(e.Start(":8080"))
}