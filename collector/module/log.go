// Package module define basic structure.
// This file defined functions which about collector logs.
package module

import "time"

// Log structure contains log properties which save in database.
// - UUID: An immutable ID.
// - Name: Log name.
// - Host: On which kubernetes server are the logs collected.
// - Namespace: Under which kubernetes namespace are the logs collected.
// - CollectTime: Logs collection time.
// - Downloads: Amount of downloads.
// - Description: Log description.
type Log struct {
	UUID        string `json:"uuid" gorm:"type:char(36);primary_key"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Host        string `json:"host" gorm:"type:varchar(255);not null"`
	Namespace   string `json:"namespace" gorm:"type:varchar(255);not null"`
	CollectTime int64  `json:"collect_time" gorm:"type:bigint unsigned;not null"`
	Downloads   int    `json:"downloads" gorm:"type:int unsigned;not null"`
	Description string `json:"description" gorm:"text"`
}

func initLog() {
	_ = db.CreateTable(&Log{}, "InnoDB", "utf8", "utf8_general_ci")
}

// NewLog returns a pointer of Log.
func NewLog(uuid, name, host, namespace, description string) *Log {
	log := new(Log)
	log.UUID = uuid
	log.Name = name
	log.Host = host
	log.Namespace = namespace
	log.CollectTime = time.Now().Unix()
	log.Downloads = 0
	log.Description = description
	return log
}

// TableName method implement schema.Tabler interface.
func (l *Log) TableName() string {
	return "log"
}

// Create method used to create a data to log table.
func (l *Log) Create() error {
	return db.Create(l)
}
