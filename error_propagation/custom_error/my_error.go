package customerror

import (
	"fmt"
	"runtime/debug"
)

// MyError - custom error
type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

// WrapError - wrap error
func WrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}
