package main

import (
	"fmt"
	"log-collector/config"
	"log-collector/route"
)

func main() {
	router := route.InitRouter()
	err := router.Run(fmt.Sprintf("%s:%s", config.ServiceHost, config.ServicePort))
	if err != nil {
		fmt.Println("Start log-collector service failed: ", err.Error())
	}
}
