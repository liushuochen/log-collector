package exception

import "fmt"

type InvalidDatabaseTypeError struct {
	baseError
	ExpectedDatabaseType string
	GotDatabaseType string
}

func NewInvalidDatabaseTypeError(expectedDatabaseType, gotDatabaseType string) InvalidDatabaseTypeError {
	e := InvalidDatabaseTypeError{
		ExpectedDatabaseType: expectedDatabaseType,
		GotDatabaseType: gotDatabaseType,
	}
	message := fmt.Sprintf("expect database type %s, but %s got.", expectedDatabaseType, gotDatabaseType)
	e.baseError = newBaseError(message)
	return e
}
