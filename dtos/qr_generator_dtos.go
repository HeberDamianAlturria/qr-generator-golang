package dtos

import (
	"github.com/skip2/go-qrcode"
)

// QRGeneratorRequest represents the request body for generating QR code.
// @description QR Generator Request
type QRGeneratorRequest struct {
	// URL to be encoded into QR code.
	// @required
	// @example https://example.com
	Url string `json:"url" validate:"required"`

	// Level of error correction used for QR code.
	// @example Low
	// @example Medium
	// @example High
	// @example Highest
	Level string `json:"level" validate:"omitempty,oneof=Low Medium High Highest"`

	// Size of QR code.
	// @example 256
	Size int `json:"size" validate:"omitempty,gte=256,lte=4096"`
}

func (qr *QRGeneratorRequest) GetLevel() qrcode.RecoveryLevel {
	switch qr.Level {
	case "Low":
		return qrcode.Low
	case "Medium":
		return qrcode.Medium
	case "High":
		return qrcode.High
	case "Highest":
		return qrcode.Highest
	default:
		return qrcode.Medium
	}
}

func (qr *QRGeneratorRequest) GetSize() int {
	if qr.Size == 0 {
		return 256
	}
	return qr.Size
}
