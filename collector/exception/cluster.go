// Package exception define all self exception types.
// This file defines cluster self exception type, functions and methods.
package exception

import "fmt"

// EmptyClusterError structure define an error that has empty IP address and domain name.
// - Name: Cluster name.
type EmptyClusterError struct {
	Name string
	baseError
}

// NewEmptyClusterError returns a EmptyClusterError with a given name.
func NewEmptyClusterError(name string) EmptyClusterError {
	e := EmptyClusterError{
		Name: name,
	}
	message := fmt.Sprintf("Cluster %s ip and domain name both empty", name)
	e.baseError = newBaseError(message)
	return e
}
