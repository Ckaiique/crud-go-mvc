package rest_err

import "net/http"

type ResErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ResErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *ResErr {
	return &ResErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ResErr {
	return &ResErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Err:     "Internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
