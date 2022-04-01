// Package database contains database operations.
// This file define MySQL operations.
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDatabase structure define a group of MySQL operations.
type MySQLDatabase struct {
	*database
}

// NewMySQLDatabase returns a pointer of MySQLDatabase.
// - name: Database name.
func NewMySQLDatabase(name string) *MySQLDatabase {
	d := new(MySQLDatabase)
	d.database = newDatabase(name)
	d.t = MySQLDatabaseType
	return d
}

// Connect method used to connect MySQL server.
// - username: Username for MySQL server
// - pwd: Password for MySQL server
// - host: Hostname or ip for MySQL server
// - port: Port for MySQL server
// - timeout: Timeout seconds for connection. If the value of the timeout is less than 10, NewMySQLDatabase will reset
//            it to 10.
func (md *MySQLDatabase) Connect(username, pwd, host string, port, timeout int) error {
	if timeout < 10 {
		timeout = 10
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds", username, pwd,
		host, port, md.name, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	md.db = db
	return nil
}

// CreateTable used to create a table in database. The CreateTable method does not return error regardless of whether
// the table exists.
// - table: A pointer of table structure.
// - engine: Table engine. Use sql "show engines" to get the list of support engines.
// - charset: MySQL table charset. Use sql "show character set" to get the list of support character sets.
// - collate: MySQL table collate. Use sql "show collation" to get the list of support collates.
func (md *MySQLDatabase) CreateTable(table interface{}, engine, charset, collate string) error {
	options := fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s COLLATE=%s ROW_FORMAT=Dynamic",
		engine, charset, collate)
	return md.db.Set("gorm:table_options", options).AutoMigrate(table)
}
