// Package exception define all self exception types.
// This file defines network exception type, functions and methods.
package exception

import "fmt"

// PingLostPackageError will be raised when the network ping lost package.
type PingLostPackageError struct {
	baseError
	host string
}

// NewPingLostPackageError returns a PingLostPackageError with given host address.
func NewPingLostPackageError(host string) PingLostPackageError {
	e := PingLostPackageError{
		host: host,
	}
	message := fmt.Sprintf("ping %s lost package", host)
	e.baseError = newBaseError(message)
	return e
}

// Hostname method returns hostname of PingLostPackageError type.
func (e PingLostPackageError) Hostname() string {
	return e.host
}
