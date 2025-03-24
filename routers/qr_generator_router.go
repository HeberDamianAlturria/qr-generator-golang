package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	"net/http"
	"qr-generator/dtos"
)

func PostGenerateQR(c echo.Context) error {
	var qrGeneratorRequest dtos.QRGeneratorRequest

	if err := c.Bind(&qrGeneratorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(qrGeneratorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	qr, err := qrcode.Encode(qrGeneratorRequest.Url, qrGeneratorRequest.GetLevel(), qrGeneratorRequest.GetSize())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", qr)
}

func AddQRGeneratorRoutes(e *echo.Echo) {
	qrGroup := e.Group("/qr")
	qrGroup.POST("/generate", PostGenerateQR)
}
