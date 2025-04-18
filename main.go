package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"qr-generator/routers"
	"qr-generator/validator"
	_ "qr-generator/docs"
)

//	@title			QR Generator API
//	@description	This is a QR code generator API. It is just an example for demonstrating how to use Swagger and Echo in Go.
//	@version		1
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routers.AddQRGeneratorRoutes(e)

	// Start server.
	e.Logger.Fatal(e.Start(":8080"))
}
