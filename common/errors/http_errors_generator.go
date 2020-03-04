package errors

import "net/http"

func NewHttpError(err error) HttpError {
	return &httpError{msg: err.Error(), status: 0}
}

func NewHttpErrorWithStatus(message string, status int) HttpError {
	return &httpError{msg: message, status: status}
}

func NewHttpErrorNotFound() HttpError {
	return &httpError{msg: "can not found requested resource", status: 404}
}

func NewHttpErrorNotFoundWithResource(resource string) HttpError {
	return &httpError{msg: "can not found requested resource", status: 404, resource: resource}
}

func NewHttpErrorUnauthorized() HttpError {
	return &httpError{msg: "unauthorized user", status: 401}
}

func NewHttpErrorNotCreated() HttpError {
	return &httpError{msg: "can not create resource", status: http.StatusBadRequest}
}

func NewHttpErrorNotCreatedWithResource(resource string) HttpError {
	return &httpError{msg: "can not create resource", status: http.StatusBadRequest, resource: resource}
}
