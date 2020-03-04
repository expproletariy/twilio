package errors

var (
	httpErrorNotFound    = NewHttpErrorNotFound()
	httpErrorBuildInType = NewHttpErrorBuildInType()
)

func IsHttpErrorNotFound(err HttpError) bool {
	return httpErrorNotFound.Is(err)
}

func IsHttpErrorBuildInType(err HttpError) bool {
	return httpErrorBuildInType.Is(err)
}
