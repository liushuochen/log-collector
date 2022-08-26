// Package req defines request body.
// This file defines some cluster requests body.
package req

// CreateClusterRequest defines cluster create body in JSON.
type CreateClusterRequest struct {
	IP          string `json:"IP" binding:"required"`
	KubeConfig  string `json:"kube_config" binding:"required"`
	Description string `json:"description" binding:""`
}

// DeleteClusterRequest defines cluster delete body in JSON.
type DeleteClusterRequest struct {
	UUID string `json:"uuid" binding:"required"`
}
