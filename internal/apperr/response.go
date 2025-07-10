package apperr

import "net/http"

type ErrorResponse struct {
	Code    int    `json:"-"`
	Message string `json:"error"`
}

var (
	ErrResBadRequest = &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	ErrResInternalServerError = &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
	ErrResUnauthorized = &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	ErrResRequireAuthHeader = &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "authorization header is required",
	}
	ErrResInvalidToken = &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "invalid token",
	}
)
