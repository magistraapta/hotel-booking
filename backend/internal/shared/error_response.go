package shared

import (
	"net/http"
	"time"
)

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Error     string `json:"error,omitempty"` // Detailed error message
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, path string, status int) *ErrorResponse {
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return &ErrorResponse{
		Success:   false,
		Message:   message,
		Path:      path,
		Status:    status,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// NewErrorResponseWithError creates an error response with detailed error message
func NewErrorResponseWithError(message string, err error, path string, status int) *ErrorResponse {
	response := NewErrorResponse(message, path, status)
	if err != nil {
		response.Error = err.Error()
	}
	return response
}

// NewBadRequestResponse is a convenience function for 400 Bad Request
func NewBadRequestResponse(message string, path string) *ErrorResponse {
	return NewErrorResponse(message, path, http.StatusBadRequest)
}

// NewNotFoundResponse is a convenience function for 404 Not Found
func NewNotFoundResponse(message string, path string) *ErrorResponse {
	return NewErrorResponse(message, path, http.StatusNotFound)
}

// NewInternalServerErrorResponse is a convenience function for 500 Internal Server Error
func NewInternalServerErrorResponse(message string, path string) *ErrorResponse {
	return NewErrorResponse(message, path, http.StatusInternalServerError)
}

// NewUnauthorizedResponse is a convenience function for 401 Unauthorized
func NewUnauthorizedResponse(message string, path string) *ErrorResponse {
	return NewErrorResponse(message, path, http.StatusUnauthorized)
}

// NewForbiddenResponse is a convenience function for 403 Forbidden
func NewForbiddenResponse(message string, path string) *ErrorResponse {
	return NewErrorResponse(message, path, http.StatusForbidden)
}
