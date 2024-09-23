package errors

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code     string      `json:"code"`
	Message  interface{} `json:"message"`
	HTTPCode int         `json:"-"`
}

// GetHTTPError get http error
func GetHTTPError(err error) HTTPError {
	e, ok := err.(*_err)
	if !ok {
		return HTTPError{
			Code:     ErrInternalServerError.Code,
			Message:  ErrInternalServerError.Message,
			HTTPCode: ErrInternalServerError.HTTPCode,
		}
	}
	return HTTPError{
		Code:     e.Code,
		Message:  e.Message,
		HTTPCode: e.HTTPCode,
	}
}
