// Package config define global variables.
// This file used to define global variables and initialise function.
package config

import (
	"log-collector/utils/encrypt"
	"os"
	"strconv"
)

const (
	// ServiceVersion indicates the version number of log-collector. Default is "0.0.0". It will be reset in build.sh
	// file.
	ServiceVersion = "__VERSION__"
)

var (
	// ServiceHost indicate log-collector host
	ServiceHost = "0.0.0.0"

	// ServicePort indicate log-collector port. It equal to environment variable "LOG_COLLECTOR_COLLECTOR_PORT", default
	// is 9189.
	ServicePort = 9189

	// DatabaseType indicate connection the type of database. It equals to environment variable
	// "LOG_COLLECTOR_DATABASE_TYPE", default is "mysql"
	DatabaseType = ""

	// DatabaseUser indicate connection the username of database. It equals to environment variable
	// "LOG_COLLECTOR_DATABASE_USER". If the value of DatabaseUser is empty, a panic will be raised.
	DatabaseUser = ""

	// DatabasePWD indicate connection the pwd of database. It equals to environment variable
	// "LOG_COLLECTOR_DATABASE_PWD". If the value of DatabasePWD is empty, a panic will be raised.
	DatabasePWD = ""

	// DatabasePort indicate database service port. It equals to environment variable "LOG_COLLECTOR_DATABASE_PORT",
	// default is 3306.
	DatabasePort = 0

	// DatabaseHost indicate hostname or host ip for database, It equals to environment variable
	// "LOG_COLLECTOR_DATABASE_HOST". Default is "127.0.0.1".
	DatabaseHost = ""
)

// InitGlobalVariable used to init global variables.
func InitGlobalVariable() {
	var err error = nil

	// init ServicePort
	servicePort, err := strconv.Atoi(os.Getenv("LOG_COLLECTOR_COLLECTOR_PORT"))
	if err == nil {
		ServicePort = servicePort
	}

	// init DatabaseType
	DatabaseType = os.Getenv("LOG_COLLECTOR_DATABASE_TYPE")
	if DatabaseType == "" {
		DatabaseType = "mysql"
	}

	// init DatabaseUser
	DatabaseUser = os.Getenv("LOG_COLLECTOR_DATABASE_USER")
	if DatabaseUser == "" {
		panic("cannot get the username of database")
	}

	// init DatabasePWD
	encryptDatabasePWD := os.Getenv("LOG_COLLECTOR_DATABASE_PWD")
	if encryptDatabasePWD == "" {
		panic("bad pwd of database")
	}
	DatabasePWD, err = encrypt.UnEncrypt(encryptDatabasePWD, encrypt.DatabaseUserKey)
	if err != nil {
		panic(err.Error())
	}

	// init DatabasePort
	DatabasePort, err = strconv.Atoi(os.Getenv("LOG_COLLECTOR_DATABASE_PORT"))
	if err != nil {
		DatabasePort = 3306
	}

	// init DatabaseHost
	DatabaseHost = os.Getenv("LOG_COLLECTOR_DATABASE_HOST")
	if DatabaseHost == "" {
		DatabaseHost = "127.0.0.1"
	}
}
