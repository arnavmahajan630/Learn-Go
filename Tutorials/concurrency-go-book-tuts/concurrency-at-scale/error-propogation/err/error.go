package something

import "runtime/debug"

type AppError struct {
	Inner error  // root error message
	Msg   string // Safe User facing message
	Op    string // Location of the error
	Stack string // Stack trace
}

func WrapE(err error, op, msg string) AppError {
	return AppError{
		Inner: err,
		Msg:   msg,
		Op:    op,
		Stack: string(debug.Stack()),
	}
}

func (e AppError) Unwrap() error {
	return e.Inner
}

func (e AppError) Error() string {
	if e.Inner != nil {
		return e.Op + ": " + e.Msg + " -> " + e.Inner.Error()
	}
	return e.Op + ": " + e.Msg
}

func (e AppError) SafeMessage() string {
	return e.Msg
}

