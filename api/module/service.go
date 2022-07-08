// Package module define basic structure.
// This file defined functions which about api-server service.
package module

import (
	"api-server/config"
	"runtime"
)

// Health structure contains service health information.
// - Status: service status. Defined in controller/service.go
type Health struct {
	Status string `json:"status"`
}

// Version structure contains service version information.
// - Service: component version.
type Version struct {
	Service *ServiceVersion `json:"service"`
}

// ServiceVersion structure contains each component version information.
type ServiceVersion struct {
	API       *APIVersion       `json:"api"`
	Collector *CollectorVersion `json:"collector"`
}

// CollectorVersion structure contains collector component information.
type CollectorVersion struct {
	Version   string `json:"collector_version"`
	GoVersion string `json:"build"`
}

// APIVersion structure contains api component information.
type APIVersion struct {
	Version   string `json:"api_version"`
	GoVersion string `json:"build"`
}

// NewVersion returns a Version pointer. It contains api server version and build version.
func NewVersion() *Version {
	service := new(ServiceVersion)

	api := new(APIVersion)
	api.Version = config.ServiceVersion
	api.GoVersion = runtime.Version()

	collector := new(CollectorVersion)

	v := new(Version)
	v.Service = service
	v.Service.API = api
	v.Service.Collector = collector
	return v
}
