// Package controller is used to actually handle HTTP requests.
// This file defined functions which about api-server service.
package controller

import (
	"api-server/config"
	"api-server/module"
	"api-server/requests"
	"encoding/json"
)

// This group of constant define service health status.
const (
	ServiceRunningStatus = "running"
	ServiceFaultStatus   = "fault"
)

// ServiceHealthCheck function is the health check in controller level.
// Health status will be set in this function.
func ServiceHealthCheck() *module.Health {
	result := new(module.Health)
	result.Status = ServiceRunningStatus
	return result
}

// GetVersion function used to return a version information which contains each component.
func GetVersion() *module.Version {
	version := module.NewVersion()
	getCollectorServiceVersion(version)
	return version
}

func getCollectorServiceVersion(version *module.Version) {
	req := requests.New(
		"get version",
		config.CollectorServiceIP,
		"/collector/v1/version",
		requests.GET,
		config.CollectorServicePort,
	)
	req.Mode = requests.HTTP
	resp, err := req.Send()
	if err != nil {
		return
	}

	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp.Body), &result)
	if err != nil {
		return
	}
	code, ok := result["code"]
	if !ok || int(code.(float64)) != 200 {
		return
	}

	msg, ok := result["message"]
	if !ok {
		return
	}
	message := msg.(map[string]interface{})

	v, _ := message["collector_version"].(string)
	version.Service.Collector.Version = v

	buildVersion, _ := message["build"].(string)
	version.Service.Collector.GoVersion = buildVersion
	return
}
