package rest_err

import (
	"net/http"
)

//this file is to define what might be the error message

type RestErr struct {
	Message string   `json:"message"` //message that the dev choose to show to the user
	Err     string   `json:"error"`
	Code    int      `json:"code"` //store code from the request that will retrieve to the client
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"` //list of wich field is incorret
	Message string `json:"message"`
}

// when receive an err from an api send the RestErr, does not need to create a varivel to the error
func (r *RestErr) Error() string {
	return r.Message
}

// NewRestErr is a constructor
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// creating methods to each error type:

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

// NewBadRequestValidationError wrong object
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// NewInternalServerError standart internal error in the application
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusNotFound,
	}
}

func NewFobiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusForbidden,
	}
}
