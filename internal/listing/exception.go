package listing

// not found error
type NotFoundError struct {
	Error	string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}