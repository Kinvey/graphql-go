package errors

// ErrorHandler is the interface used to create custom panic errors that occur during query execution
type ErrorHandler interface {
	TransformError(error) error
}

// DefaultErrorHandler is the default ErrorHandler
type DefaultErrorHandler struct{}

// MakePanicError creates a new QueryError from a panic that occurred during execution
func (h *DefaultErrorHandler) TransformError(e error) error {
	return e
}
