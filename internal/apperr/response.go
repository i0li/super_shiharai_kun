package apperr

type ErrResponse struct {
	Error string `json:"error"`
}

const (
	ErrMsgBadRequest          = "bad request"
	ErrMsgUnauthorized        = "unauthorized"
	ErrMsgInternalServerError = "internal server error"
)
