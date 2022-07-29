// Package database contains database operations.
// This file define basic structure of database and it init functions.
package database

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	UnknownDatabaseType = "unknown"
	MySQLDatabaseType = "mysql"
)

// Database structure define a group of common fields and methods for database operation.
// - t: Database type.
// - name: Database name.
// - db: A object for *gorm.DB. Create in Database.Connect method.
type database struct {
	t  string
	name  string
	db    *gorm.DB
}

func newDatabase(name string) *database {
	d := new(database)
	d.name = name
	return d
}

// Type returns the type of database. If the value of the d.t is empty, Type will return UnknownDatabaseType.
func (d *database) Type() string {
	if d.t == "" {
		return UnknownDatabaseType
	}
	return d.t
}

// Name returns the database name.
func (d *database) Name() string {
	return d.name
}

// String method implement fmt.Stringer
func (d *database) String() string {
	return fmt.Sprintf("%s for %s", d.Name(), d.Type())
}
