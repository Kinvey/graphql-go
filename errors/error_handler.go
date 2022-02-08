package errors

// ErrorHandler is the interface used to transform errors that occur during query execution
type ErrorHandler interface {
	TransformError(error) error
}

// DefaultErrorHandler is the default ErrorHandler
type DefaultErrorHandler struct{}

// TransformError could transform a error to a new error(hiding any sensitive information, logging the error, etc.)
func (h *DefaultErrorHandler) TransformError(e error) error {
	return e
}
