package shared

import (
	"net/http"
	"time"
)

// ApiResponse represents a successful API response
type ApiResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Status    int         `json:"status"`
	Path      string      `json:"path,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// NewApiResponse creates a new successful API response
func NewApiResponse(message string, data interface{}, status int, path string) *ApiResponse {
	return &ApiResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Status:    status,
		Path:      path,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// NewSuccessResponse is a convenience function for common success responses
func NewSuccessResponse(message string, data interface{}, status int, path string) *ApiResponse {
	if status == 0 {
		status = http.StatusOK
	}
	return NewApiResponse(message, data, status, path)
}

// NewCreatedResponse is a convenience function for 201 Created responses
func NewCreatedResponse(message string, data interface{}, path string) *ApiResponse {
	return NewApiResponse(message, data, http.StatusCreated, path)
}
