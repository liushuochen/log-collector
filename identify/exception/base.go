// Package exception define all self exception types.
// This file defines basic self exception type, functions and methods.
package exception


type baseError struct {
	message  string
}

func newBaseError(message string) baseError {
	return baseError{message: message}
}

// String method satisfied fmt.Stringer
func (e baseError) String() string {
	return e.message
}

// Error method satisfied exception
func (e baseError) Error() string {
	return e.message
}
