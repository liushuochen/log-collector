// Package controller is used to actually handle HTTP requests.
// This file defined functions which about collector service.
package controller

import "log-collector/module"

// This group of constant define service health status.
const (
	RUNNING = "running"
	FAULT   = "fault"
)

// ServiceHealthCheck function is the health check in controller level.
// Health status will be set in this function.
func ServiceHealthCheck() *module.Health {
	result := new(module.Health)
	result.Status = RUNNING
	return result
}

// GetServiceVersion returns a *module.Version that contains version information.
func GetServiceVersion() *module.Version {
	version := new(module.Version)
	version.GetVersion()
	return version
}
