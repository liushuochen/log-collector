package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log-collector/requests"
)


func InitRouter() *gin.Engine {
    router := gin.New()
    router.Use(cors.New(cors.Config{
    	AllowOrigins: []string{"*"},
    	AllowMethods: []string{requests.GET, requests.POST, requests.PUT, requests.DELETE},
    	// AllowHeaders: []string{"token"},
    	ExposeHeaders: []string{"Content-Length"},
    	AllowCredentials: true,
	}))

    // apiServiceV1 := router.Group("/logcollector/v1")
    return router
}
