package errx

// New returns a new *Error
func New(code Code, subCode SubCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		SubCode: subCode,
	}
}

// Wrap returns a *Error from an error
func Wrap(err error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return New(
		GetCode(""),
		GetSubCode(""),
		err.Error(),
	)
}
