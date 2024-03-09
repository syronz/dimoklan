package api

import (
	"dimoklan/internal/errors/errstatus"
	"errors"
	"net/http"
)

// status return back the status code based on the error
func status(err error) int {
	switch {
	case errors.Is(err, errstatus.ErrBadRequest):
		return http.StatusBadRequest

	case errors.Is(err, errstatus.ErrUnauthorized):
		return http.StatusUnauthorized

	case errors.Is(err, errstatus.ErrForbidden):
		return http.StatusForbidden

	case errors.Is(err, errstatus.ErrNotFound):
		return http.StatusNotFound

	case errors.Is(err, errstatus.ErrMethodNotAllowed):
		return http.StatusMethodNotAllowed

	case errors.Is(err, errstatus.ErrNotAcceptable):
		return http.StatusNotAcceptable

	case errors.Is(err, errstatus.ErrRequestTimeout):
		return http.StatusRequestTimeout

	case errors.Is(err, errstatus.ErrConflict):
		return http.StatusConflict

	case errors.Is(err, errstatus.ErrPreconditionFailed):
		return http.StatusPreconditionFailed

	case errors.Is(err, errstatus.ErrRequestEntityTooLarge):
		return http.StatusRequestEntityTooLarge

	case errors.Is(err, errstatus.ErrUnprocessableEntity):
		return http.StatusUnprocessableEntity

	case errors.Is(err, errstatus.ErrLocked):
		return http.StatusLocked

	case errors.Is(err, errstatus.ErrPreconditionRequired):
		return http.StatusPreconditionFailed

	case errors.Is(err, errstatus.ErrTooManyRequests):
		return http.StatusTooManyRequests

	case errors.Is(err, errstatus.ErrRequestHeaderFieldsTooLarge):
		return http.StatusRequestHeaderFieldsTooLarge

	case errors.Is(err, errstatus.ErrInternalServerError):
		return http.StatusInternalServerError

	case errors.Is(err, errstatus.ErrBadGateway):
		return http.StatusBadGateway

	case errors.Is(err, errstatus.ErrGatewayTimeout):
		return http.StatusGatewayTimeout

	default:
		return http.StatusInternalServerError
	}
}
