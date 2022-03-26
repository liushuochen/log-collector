// Package api contains API functions which called by route package.
// This file defined methods which about collector service.
package api

import (
	"github.com/gin-gonic/gin"
	"log-collector/api/resp"
	"log-collector/controller"
)

type Service struct{}

// HealthCheck function used to check service status.
// Currently is only return a simple HTTP response without doing nothing.
// Consider using it in kubernetes pod livenessProve.
func (service *Service) HealthCheck(c *gin.Context) {
	response := controller.ServiceHealthCheck()
	resp.SendResponse(c, resp.Ok, response)
}

// Version function used to get the version information.
// response (not real HTTP response) example:
// {
//     "service_version": "0.0.0",
//     "go_version": "go1.16.5"
// }
func (service *Service) Version(c *gin.Context) {
	response := controller.GetServiceVersion()
	resp.SendResponse(c, resp.Ok, response)
}
