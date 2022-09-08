// Package route used to init log-collector router
package route

import (
	"log-collector/api"
	"log-collector/requests"

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

	apiServiceV1 := router.Group("/collector/v1")
	initServiceRouter(apiServiceV1)
	initClusterRouter(apiServiceV1)

	apiServiceV2 := router.Group("/collector/v2")
	initClusterRouterV2(apiServiceV2)

	return router
}

func initServiceRouter(group *gin.RouterGroup) {
	service := new(api.Service)
	group.GET("/health", service.HealthCheck)
	group.GET("/version", service.Version)
}

func initClusterRouter(group *gin.RouterGroup) {
	cluster := new(api.Cluster)
	group.POST("/cluster/create", cluster.CreateWithKubeConfigContent)
	group.DELETE("/cluster/delete", cluster.Delete)
	group.PUT("/cluster/edit", cluster.Edit)
}

func initClusterRouterV2(group *gin.RouterGroup) {
	cluster := new(api.Cluster)
	group.POST("/cluster/create", cluster.CreateWithKubeConfigFile)
}
