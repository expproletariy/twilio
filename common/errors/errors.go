package errors

import "errors"

var (
	HttpNotFoundErr = errors.New("can not found requested resource")
)
