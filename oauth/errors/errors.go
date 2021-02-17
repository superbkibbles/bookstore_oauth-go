package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "Bad_request",
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "Not_found",
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Error:   "internal_server_error",
		Status:  http.StatusInternalServerError,
	}
}
