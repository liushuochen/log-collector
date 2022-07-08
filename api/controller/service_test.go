// Package controller_test used to test controller.
package controller_test

import (
	"api-server/controller"
	"testing"
)

func TestServiceHealthCheckStatus(t *testing.T) {
	expected := []string{controller.ServiceRunningStatus, controller.ServiceFaultStatus}

	result := controller.ServiceHealthCheck()
	find := false
	for _, status := range expected {
		if result.Status == status {
			find = true
			break
		}
	}

	if !find {
		t.Errorf("unknown health status: %s", result.Status)
	}
}
