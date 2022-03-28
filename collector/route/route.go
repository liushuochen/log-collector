// Package route used to init log-collector router
package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log-collector/api"
	"log-collector/requests"
)

// InitRouter function used to init router, include origins, methods, headers, credentials, api groups.
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{requests.GET, requests.POST, requests.PUT, requests.DELETE},
		// AllowHeaders: []string{"token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	apiServiceV1 := router.Group("/logcollector/v1")
	initServiceRouter(apiServiceV1)

	return router
}

func initServiceRouter(group *gin.RouterGroup) {
	service := new(api.Service)
	group.GET("/health", service.HealthCheck)
	group.GET("/version", service.Version)
}