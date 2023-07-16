package error

type Error struct {
	code    int
	message string
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}

func NewError(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}
