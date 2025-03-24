package dtos

type QRGeneratorRequest struct {
	Url string `json:"url" validate:"required"`
}
