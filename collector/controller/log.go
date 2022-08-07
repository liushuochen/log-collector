// Package controller is used to actually handle HTTP requests.
// This file defined functions which about collector log.
package controller

import (
	"log-collector/module"
)

// LogCreate function will collect the log and write the log information to database.
// - uuid: A UUID string for this Log.
// - name: Log name.
// - host: On which kubernetes server are the logs collected.
// - namespace: Under which kubernetes namespace are the logs collected.
// - description: Log description.
// TODO: collect the log.
func LogCreate(uuid, name, host, namespace, description string) (*module.Log, error) {
	log := module.NewLog(uuid, name, host, namespace, description)
	err := log.Create()
	if err != nil {
		return nil, err
	}
	return log, err
}
