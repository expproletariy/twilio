package errors

type HttpErrorContainer interface {
	error
	Status() int
	FullError() string
}

type HttpErrorComparator interface {
	Is(err HttpErrorContainer) bool
}

type HttpError interface {
	HttpErrorContainer
	HttpErrorComparator
}

type httpError struct {
	msg      string
	status   int
	resource string
}

func (e httpError) Error() string {
	return e.msg
}

func (e httpError) FullError() string {
	return e.msg + ": " + e.resource
}

func (e httpError) Status() int {
	return e.status
}

func (e httpError) Is(err HttpErrorContainer) bool {
	if e.status == err.Status() && e.msg == err.Error() {
		return true
	}
	return false
}
