// Package module define basic structure.
// This file defined functions which about collector kubernetes clusters.
package module

import "time"

const (
	ClusterStatusCreating = "creating"
	ClusterStatusActive   = "active"
	ClusterStatusFault    = "fault"
)

// Cluster means kubernetes cluster.
// UUID: Cluster uuid. It generates by collector service.
// IP: Cluster IP address. Support IPv4 and IPv6.
// DomainName: Cluster domain name.
// Name: Cluster name. Defined by user.
// Status: Cluster status. Dot a kubernetes argument.
// Description: Cluster description or comment.
// KubeConfig: Content of kube config file for this cluster.
// CreateTime: Cluster create time in collector service.
type Cluster struct {
	UUID        string `json:"uuid" gorm:"type:char(36);primary_key"`
	IP          string `json:"ip" gorm:"type:varchar(255);unique_index"`
	DomainName  string `json:"domain_name" gorm:"type:varchar(255);unique_index"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Status      string `json:"status" gorm:"type:varchar(30);not null"`
	Description string `json:"description" gorm:"type:text"`
	KubeConfig  string `json:"kube_config" gorm:"type:longtext"`
	CreateTime  int64  `json:"create_time" gorm:"type:bigint unsigned;not null"`
}

func initCluster() {
	_ = db.CreateTable(&Cluster{}, "InnoDB", "utf8", "utf8_general_ci")
}

// NewCluster returns a pointer of Cluster.
func NewCluster(uuid, ip, domainName, name, description, kubeConfig string) *Cluster {
	cluster := new(Cluster)
	cluster.UUID = uuid
	cluster.IP = ip
	cluster.DomainName = domainName
	cluster.Name = name
	cluster.Status = ClusterStatusCreating
	cluster.Description = description
	cluster.KubeConfig = kubeConfig
	cluster.CreateTime = time.Now().Unix()
	return cluster
}

// TableName method implement schema.Tabler interface.
func (cluster *Cluster) TableName() string {
	return "cluster"
}

// Create method used to create a data to cluster table.
func (cluster *Cluster) Create() error {
	return db.Create(cluster)
}

// Delete method used to create a data to cluster table.
func (cluster *Cluster) Delete() error {
	return db.Delete(cluster)
}

// Update method used to update a data to cluster table.
func (cluster *Cluster) Update() error {
	return db.Update(cluster)
}

// Search method used to search target cluster based on designated uuid.
func (cluster *Cluster) Search() error {
	return db.Search(cluster, "uuid=?", cluster.UUID)
}
