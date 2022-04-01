// Package exception define all self exception types.
// This file define errors which about encrypt
package exception

import "fmt"

// InvalidKeyLengthError will be raised when the length of the key is not an integer multiple of 16.
type InvalidKeyLengthError struct {
	baseError
	currentLength int
	expectedLengthMin int
	expectedLengthMax int
}

// NewInvalidKeyLengthError return a InvalidKeyLengthError based on a key
func NewInvalidKeyLengthError(key string) InvalidKeyLengthError {
	e := InvalidKeyLengthError{currentLength: len(key)}
    div := e.currentLength / 16
    e.expectedLengthMin = div * 16
    e.expectedLengthMax = (div + 1) * 16
    message := fmt.Sprintf(
		"expected key length is %d or %d, but %d got",
		e.expectedLengthMin,
		e.expectedLengthMax,
		e.currentLength,
	)
    e.baseError = newBaseError(message)
    return e
}
