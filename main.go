package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"qr-generator/routers"
	"qr-generator/validator"
)

func main() {
	e := echo.New()

	// Custom validator.
	e.Validator = validator.NewCustomValidator()

	// Middlewares.
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes.
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routers.AddQRGeneratorRoutes(e)

	// Start server.
	e.Logger.Fatal(e.Start(":8080"))
}
