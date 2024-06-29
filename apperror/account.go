package apperror

func NewWrongPassword(err error, message string) error {
	return NewAppError(CodeWrongPassword, message, err)
}
