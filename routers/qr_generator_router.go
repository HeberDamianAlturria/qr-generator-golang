package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	"net/http"
	"qr-generator/dtos"
)

// PostGenerateQR godoc
//
//	@Summary		Generate QR code
//	@Description	Generate QR code from URL
//	@Tags			QR Generator
//	@Accept			json
//	@Produce		png
//	@Param			qrGeneratorRequest	body		dtos.QRGeneratorRequest	true	"QR Generator Request"
//	@Success		200					{file}		blob					"OK"
//	@Failure		400					{string}	string					"Bad Request"
//	@Failure		500					{string}	string					"Internal Server Error"
//	@Router			/qr/generate [post]
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
