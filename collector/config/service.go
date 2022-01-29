package config

import "os"

var (
	ServiceHost = "0.0.0.0"
	ServicePort = "9188"
)

func init() {
	ServicePort = os.Getenv("LOG_COLLECTOR_PORT")
	if ServicePort == "" {
		ServicePort = "9188"
	}
}
