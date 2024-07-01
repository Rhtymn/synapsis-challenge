package apperror

func NewInvalidToken(err error) error {
	return NewAppError(
		CodeInvalidToken,
		"invalid token",
		err,
	)
}
