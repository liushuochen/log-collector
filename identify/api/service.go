// Package api contains API functions which called by route package.
// This file defined methods which about identify service.
package api

import (
	"github.com/gin-gonic/gin"
	"identify/api/resp"
	"identify/controller"
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
