package apperr

type ErrResponse struct {
	Error string `json:"error"`
}

const (
	ErrMsgBadRequest            = "bad request"
	ErrMsgUnauthorized          = "unauthorized"
	ErrMsgNoAuthorizationHeader = "authorization header is required"
	ErrMsgInvalidToken          = "invalid token"
	ErrMsgInternalServerError   = "internal server error"
)
