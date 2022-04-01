// Package module define basic structure.
// This file defined functions which about collector logs.
package module

// Log structure contains log properties which save in database.
// - UUID: An immutable ID.
// - Name: Log name.
// - Host: On which kubernetes server are the logs collected.
// - Namespace: Under which kubernetes namespace are the logs collected.
// - CollectTime: Logs collection time.
// - Downloads: Amount of downloads.
// - Description: Logs description.
type Log struct {
	UUID  string  `json:"uuid" gorm:"type:char(36);primary_key"`
	Name  string  `json:"name" gorm:"type:varchar(255);not null"`
	Host  string  `json:"host" gorm:"type:varchar(255);not null"`
	Namespace string `json:"namespace" gorm:"type:varchar(255);not null"`
	CollectTime int64 `json:"collect_time" gorm:"type:bigint unsigned;not null"`
	Downloads int `json:"downloads" gorm:"type:int unsigned;not null"`
	Description string `json:"description" gorm:"text"`
}

func initLog() {
	_ = db.CreateTable(&Log{}, "InnoDB", "utf8", "utf8_general_ci")
}

// TableName method implement schema.Tabler interface.
func (l *Log) TableName() string {
	return "log"
}
