// Package api contains API functions which called by route package.
// This file defined methods which about api service.
package api

import (
	"api-server/api/resp"
	"api-server/controller"
	"github.com/gin-gonic/gin"
)

type Service struct{}

// HealthCheck function used to check service status.
// Currently, is only return a simple HTTP response without doing nothing.
// Consider using it in kubernetes pod livenessProve.
// Response (not real HTTP response) example:
// {
//     "status": "running"
// }
func (service *Service) HealthCheck(c *gin.Context) {
	response := controller.ServiceHealthCheck()
	resp.SendResponse(c, resp.Ok, response)
}

// GetVersion function used to get each microservice components version information.
// Response (not real HTTP response) example:
// {
//      "service": {
//          "api": {
//				"api_version": "0.1.0",
//				"build": "go1.16.5"
//          },
//			"collector": {
//				"collector_version": "0.1.0",
//				"build": "go1.16.5"
//			}
// }
func (service *Service) GetVersion(c *gin.Context) {
	response := controller.GetVersion()
	resp.SendResponse(c, resp.Ok, response)
}
