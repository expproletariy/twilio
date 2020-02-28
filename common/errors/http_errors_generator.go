package errors

func NewHttpErrorNotFound() HttpError {
	return &httpError{msg: "can not found requested resource", status: 404}
}
