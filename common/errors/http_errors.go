package errors

var (
	httpErrorNotFound = NewHttpErrorNotFound()
)

func IsHttpErrorNotFound(err HttpError) bool {
	return httpErrorNotFound.Is(err)
}
