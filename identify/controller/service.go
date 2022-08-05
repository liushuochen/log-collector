// Package controller is used to actually handle HTTP requests.
// This file defined functions which about identify service.
package controller

import "identify/module"

// ServiceHealthCheck function is the health check in controller level.
// Health status will be set in this function.
func ServiceHealthCheck() *module.Health {
	result := new(module.Health)
	result.Status = module.ServiceRunningStatus
	return result
}

// GetServiceVersion returns a *module.Version that contains version information.
func GetServiceVersion() *module.Version {
	return module.GetVersion()
}
