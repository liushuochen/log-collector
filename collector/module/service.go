// Package module define basic structure.
// This file defined functions which about collector service.
package module

import "runtime"

const (
	// ServiceVersion indicates the version number of log-collector. Default is "0.0.0". It will be reset in build.sh
	ServiceVersion = "__VERSION__"
)

// Health structure contains service health information.
// - Status: service status. Defined in controller/service.go
type Health struct {
	Status string `json:"status"`
}

// Version structure contains service version information
// - ServiceVersion: log-collector version.
// - GoVersion: go build version.
type Version struct {
	ServiceVersion string `json:"service_version"`
	GoVersion      string `json:"go_version"`
}

// GetVersion method is used to set the fields in the Version pointer object.
func (v *Version) GetVersion() {
	v.ServiceVersion = ServiceVersion
	v.GoVersion = runtime.Version()
}
