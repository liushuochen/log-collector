package main

import (
	"api-server/config"
	"api-server/route"
	"fmt"
)

func init() {
	config.InitGlobalVariable()
}

func main() {
	router := route.InitRouter()
	err := router.Run(fmt.Sprintf("%s:%d", config.ServiceHost, config.ServicePort))
	if err != nil {
		fmt.Println("Start api-server service failed: ", err.Error())
	}
}
