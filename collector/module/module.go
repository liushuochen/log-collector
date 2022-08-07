// Package module define basic structure.
// This file used to initialise database.
package module

import (
	"fmt"
	"log-collector/config"
	"log-collector/database"
)

var (
	// Database object
	db *database.MySQLDatabase
)

// InitDatabase used to initialise database. This function will connect database and create tables if not exist.
// A panic will be raised when connect database failed.
func InitDatabase() {
	db = database.NewMySQLDatabase("collector")
	err := db.Connect(
		config.DatabaseUser,
		config.DatabasePWD,
		config.DatabaseHost,
		config.DatabasePort,
		10,
		)
	if err != nil {
		panic(fmt.Sprintf("connect database failed: %s", err.Error()))
	}

	createTables()
}

func createTables() {
	initLog()
	initResource()
	initCluster()
}
