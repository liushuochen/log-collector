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

// ClusterNotFoundError structure define an error that cannot find cluster by UUID from database.
type ClusterNotFoundError struct {
	UUID string
	baseError
}

// NewClusterNotFoundError returns a ClusterNotFoundError with a given uuid.
func NewClusterNotFoundError(uuid string) ClusterNotFoundError {
	e := ClusterNotFoundError{
		UUID: uuid,
	}
	message := fmt.Sprintf("Cannot find cluster %s", e.UUID)
	e.baseError = newBaseError(message)
	return e
}
