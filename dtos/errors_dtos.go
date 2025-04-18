package dtos

type ErrorResponse struct {
	// Error message.
	// @example Bad Request
	Message string `json:"message"`

	// Detailed error message.
	// @example URL is required
	Details string `json:"details"`
}
