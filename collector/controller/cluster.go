// Package controller is used to actually handle HTTP requests.
// This file defined functions which about collector cluster.
package controller

import (
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"log-collector/exception"
	"log-collector/module"
	"log-collector/requests"
)

// ClusterCreate returns a cluster object and an error. This function will create a goroutine that check cluster status.
func ClusterCreate(uuid, ip, description, kubeConfig string) (*module.Cluster, error) {
	domainName, name := parsingKubeConfig(transKubeConfig([]byte(kubeConfig)))
	if ip == "" && domainName == "" {
		return nil, exception.NewEmptyClusterError(name)
	}

	cluster := module.NewCluster(uuid, ip, domainName, name, description, kubeConfig)
	err := cluster.Create()
	if err != nil {
		return nil, err
	}

	go func() {
		address := ""
		if ip != "" {
			address = ip
		} else {
			address = domainName
		}

		if checkCluster(address) {
			cluster.Status = module.ClusterStatusActive
		} else {
			cluster.Status = module.ClusterStatusFault
		}

		_ = cluster.Update()
	}()
	return cluster, nil
}

// ClusterDelete function returns a module.Cluster pointer. It will delete cluster in database.
// If cluster does not exist in database, an exception.ClusterNotFoundError will return.
func ClusterDelete(uuid string) (*module.Cluster, error) {
	cluster, err := searchCluster(uuid)
	if err != nil {
		return nil, err
	}

	err = cluster.Delete()
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

// ClusterEdit function used to update name, description for designated cluster with given uuid.
// It returns a pointer of module.Cluster which contain given name and description, and an error.
// If cluster does not exist in database, an exception.ClusterNotFoundError will return.
func ClusterEdit(uuid, name, description string) (*module.Cluster, error) {
	cluster, err := searchCluster(uuid)
	if err != nil {
		return nil, err
	}

	if name != "" {
		cluster.Name = name
	}
	if description != "" {
		cluster.Description = description
	}
	err = cluster.Update()
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func searchCluster(uuid string) (*module.Cluster, error) {
	cluster := new(module.Cluster)
	cluster.UUID = uuid
	err := cluster.Search()
	if err == gorm.ErrRecordNotFound {
		return nil, exception.NewClusterNotFoundError(uuid)
	}

	return cluster, nil
}

func checkCluster(address string) bool {
	req := requests.New(
		"check cluster state",
		address,
		"api/v1/namespaces/cran1/configmaps",
		requests.GET,
		6443,
	)
	req.CloseVerify()
	req.AddHeader("Token", "DEFAULT")

	resp, err := req.Send()
	if err != nil {
		return false
	}

	if resp.Code >= 500 && resp.Code != 404 {
		return false
	}
	return true
}

func parsingKubeConfig(kubeConfig map[string]interface{}) (string, string) {
	clusterInfoList := kubeConfig["clusters"].([]interface{})
	if len(clusterInfoList) <= 0 {
		return "", ""
	}
	clusterInfo := clusterInfoList[0].(map[interface{}]interface{})
	cluster := clusterInfo["cluster"].(map[interface{}]interface{})
	domain := cluster["server"].(string)
	name := clusterInfo["name"].(string)
	return domain, name
}

func transKubeConfig(kubeConfig []byte) map[string]interface{} {
	content := make(map[string]interface{})
	_ = yaml.Unmarshal(kubeConfig, content)
	return content
}
