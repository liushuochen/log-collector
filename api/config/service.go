// Package config define global variables.
// This file used to define global variables and initialise function.
package config

import (
	"os"
	"strconv"
)

const (
	// ServiceVersion indicates the version number of api-server. Default is "0.0.0". It will be reset in build.sh file.
	ServiceVersion = "__VERSION__"
)

var (
	// ServiceHost indicate api-server host.
	ServiceHost = "0.0.0.0"

	// ServicePort indicate api-server port. It equals to environment variable "LOG_COLLECTOR_API_SERVER_PORT", default
	// is 9188.
	ServicePort = 9188

	// CollectorServicePort indicate log-collector port. It equals to environment variable
	// "LOG_COLLECTOR_COLLECTOR_PORT", default is 9189.
	CollectorServicePort = 9189

	// IdentifyServicePort indicate log-collector port. It equals to environment variable
	// "LOG_COLLECTOR_IDENTIFY_PORT", default is 9190.
	IdentifyServicePort = 9190

	// CollectorServiceIP indicate collector IP address. It will return a panic if the value is empty. It equals to
	// environment variable "LOG_COLLECTOR_COLLECTOR_SERVICE_IP"
	CollectorServiceIP = ""

	// IdentifyServiceIP indicate identify IP address. It will return a panic if the value is empty. It equals to
	// environment variable "LOG_COLLECTOR_IDENTIFY_SERVICE_IP"
	IdentifyServiceIP = ""
)

// InitGlobalVariable used to init global variables.
func InitGlobalVariable() {
	// init ServicePort
	servicePort, err := strconv.Atoi(os.Getenv("LOG_COLLECTOR_API_SERVER_PORT"))
	if err == nil {
		ServicePort = servicePort
	}

	// init CollectorServicePort
	collectorServicePort, err := strconv.Atoi(os.Getenv("LOG_COLLECTOR_COLLECTOR_PORT"))
	if err == nil {
		CollectorServicePort = collectorServicePort
	}

	// init IdentifyServicePort
	identifyServicePort, err := strconv.Atoi(os.Getenv("LOG_COLLECTOR_IDENTIFY_PORT"))
	if err == nil {
		IdentifyServicePort = identifyServicePort
	}

	// init CollectorServiceIP
	CollectorServiceIP = os.Getenv("LOG_COLLECTOR_COLLECTOR_SERVICE_IP")
	if CollectorServiceIP == "" {
		panic("empty collector service ip")
	}

	// init IdentifyServiceIP
	IdentifyServiceIP = os.Getenv("LOG_COLLECTOR_IDENTIFY_SERVICE_IP")
	if IdentifyServiceIP == "" {
		panic("empty identify service ip")
	}
}
