package errx

var _ error = (*Error)(nil)

// Error returns the *Error message
func (e *Error) Error() string {
	return e.Message
}

// E returns the *Error as an error
func (e *Error) E() error {
	return e
}

// WithError appends an error to the *Error
func (e *Error) WithError(err error) *Error {
	e.Raw = err
	return e
}

// WithCode switches the *Error code
func (e *Error) WithCode(code Code) *Error {
	e.Code = code
	return e
}

// WithSubCode switches the *Error sub code
func (e *Error) WithSubCode(subCode SubCode) *Error {
	e.SubCode = subCode
	return e
}

// With switches code and sub code of the *Error
func (e *Error) With(code Code, subCode SubCode) *Error {
	return e.WithCode(code).WithSubCode(subCode)
}
