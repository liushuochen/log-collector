// Package exception define all self exception types.
// This file defines database exception type, functions and methods.
package exception

import "fmt"

// InvalidDatabaseTypeError structure define an error that designated unsupported database type.
// - Name: Cluster name.
type InvalidDatabaseTypeError struct {
	baseError
	ExpectedDatabaseType string
	GotDatabaseType      string
}

// NewInvalidDatabaseTypeError returns a InvalidDatabaseTypeError with expected database type and unsupported database
// type.
func NewInvalidDatabaseTypeError(expectedDatabaseType, gotDatabaseType string) InvalidDatabaseTypeError {
	e := InvalidDatabaseTypeError{
		ExpectedDatabaseType: expectedDatabaseType,
		GotDatabaseType:      gotDatabaseType,
	}
	message := fmt.Sprintf("expect database type %s, but %s got.", expectedDatabaseType, gotDatabaseType)
	e.baseError = newBaseError(message)
	return e
}
