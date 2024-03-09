package errstatus

import "errors"

var (
	ErrBadRequest                  = errors.New("400")
	ErrUnauthorized                = errors.New("401")
	ErrForbidden                   = errors.New("403")
	ErrNotFound                    = errors.New("404")
	ErrMethodNotAllowed            = errors.New("405")
	ErrNotAcceptable               = errors.New("406")
	ErrRequestTimeout              = errors.New("408")
	ErrConflict                    = errors.New("409")
	ErrPreconditionFailed          = errors.New("412")
	ErrRequestEntityTooLarge       = errors.New("413")
	ErrUnprocessableEntity         = errors.New("422")
	ErrLocked                      = errors.New("423")
	ErrPreconditionRequired        = errors.New("428")
	ErrTooManyRequests             = errors.New("429")
	ErrRequestHeaderFieldsTooLarge = errors.New("431")
	ErrInternalServerError         = errors.New("500")
	ErrBadGateway                  = errors.New("502")
	ErrGatewayTimeout              = errors.New("504")
)
