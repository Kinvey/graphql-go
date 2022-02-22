package errors

import "context"

// ErrorHandler is the interface used to transform errors that occur during query execution
type ErrorHandler interface {
	TransformError(context.Context, error) error
}

// DefaultErrorHandler is the default ErrorHandler
type DefaultErrorHandler struct{}

// TransformError could transform a error to a new error(hiding any sensitive information, logging the error, etc.)
func (h *DefaultErrorHandler) TransformError(ctx context.Context, e error) error {
	return e
}
