// Package controller is used to actually handle HTTP requests.
// This file defined functions which about collector cluster.
package controller

import (
	"gopkg.in/yaml.v2"
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
