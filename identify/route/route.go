// Package route used to init log-collector router
package route

import (
	"identify/api"
	"identify/requests"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	apiServiceV1 := router.Group("/identify/v1")
	initServiceRouter(apiServiceV1)

	return router
}

func initServiceRouter(group *gin.RouterGroup) {
	service := new(api.Service)
	group.GET("/health", service.HealthCheck)
}
