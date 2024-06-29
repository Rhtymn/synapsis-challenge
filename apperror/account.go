package apperror

func NewWrongPassword(err error, message string) error {
	return NewAppError(CodeWrongPassword, message, err)
}

func NewAlreadyVerified(message string) error {
	return NewAppError(CodeAlreadyVerified, message, nil)
}

func NewInvalidVerifyEmailToken(err error) error {
	return NewAppError(CodeBadRequest, "invalid token", err)
}
