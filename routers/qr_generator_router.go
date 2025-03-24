package routers

import (
	"github.com/caiguanhao/readqr"
	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	"net/http"
	"path/filepath"
	"qr-generator/dtos"
	"strings"
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
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Bad Request", Details: err.Error()})
	}

	if err := c.Validate(qrGeneratorRequest); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Bad Request", Details: err.Error()})
	}

	qr, err := qrcode.Encode(qrGeneratorRequest.Url, qrGeneratorRequest.GetLevel(), qrGeneratorRequest.GetSize())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: "Internal Server Error", Details: err.Error()})
	}

	return c.Blob(http.StatusOK, "image/png", qr)
}

// PostDecodeQR godoc
//
//	@Summary		Decode QR code
//	@Description	Decode QR code from image file
//	@Tags			QR Generator
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file					true	"Image file"
//	@Success		200		{object}	dtos.QRDecoderResponse	"OK"
//	@Failure		400		{string}	string					"Bad Request"
//	@Router			/qr/decode [post]
func PostDecodeQR(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	filename := file.Filename

	if !strings.HasSuffix(filename, ".png") {
		extension := filepath.Ext(filename)

		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Only PNG file is supported", Details: "File extension is " + extension + " instead of .png"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Bad Request", Details: err.Error()})
	}
	defer src.Close()

	code, err := readqr.Decode(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Message: "Bad Request", Details: err.Error()})
	}

	return c.JSON(http.StatusOK, dtos.QRDecoderResponse{DecodedText: code})
}

func AddQRGeneratorRoutes(e *echo.Echo) {
	qrGroup := e.Group("/qr")
	qrGroup.POST("/generate", PostGenerateQR)
	qrGroup.POST("/decode", PostDecodeQR)
}
