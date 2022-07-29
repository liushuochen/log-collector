package main

import (
	"fmt"
	"identify/config"
	"identify/route"
)

func init() {
	config.InitGlobalVariable()
}

func main() {
	router := route.InitRouter()
	err := router.Run(fmt.Sprintf("%s:%d", config.ServiceHost, config.ServicePort))
	if err != nil {
		fmt.Println("Start identify service failed: ", err.Error())
	}
}
