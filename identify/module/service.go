// Package module define basic structure.
// This file defined functions which about identify service.
package module

import (
	"identify/config"
	"runtime"
)

// This group of constant define service health status.
const (
	ServiceRunningStatus = "running"
	ServiceFaultStatus   = "fault"
)

// Health structure contains service health information.
// - Status: service status. Defined in controller/service.go
type Health struct {
	Status string `json:"status"`
}

// Version structure contains service version information.
// - ServiceVersion: identify version.
// - GoVersion: go build version.
type Version struct {
	ServiceVersion string `json:"identify_version"`
	GoVersion      string `json:"build"`
}

// GetVersion returns a Version pointer.
func GetVersion() *Version {
	v := new(Version)
	v.ServiceVersion = config.ServiceVersion
	v.GoVersion = runtime.Version()
	return v
}
