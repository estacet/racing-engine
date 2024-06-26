package apperror

type EntityNotFoundError struct {
	message string
}

func NewEntityNotFoundError(msg string) *EntityNotFoundError {
	return &EntityNotFoundError{
		message: msg,
	}
}

func (e EntityNotFoundError) Error() string {
	return e.message
}
