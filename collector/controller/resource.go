// Package controller is used to actually handle HTTP requests.
// This file defined functions which about kubernetes resources management.
package controller

import (
	"log-collector/module"
	"strings"
)

// CreateResource function used to add a kubernetes resource.
// - name: Resource name.
// - apiVersion: Kubernetes resource api version(e.g. v1, apps/v1 etc).
// - description: Description: Resource description.
// The initial state of the resource is module.ResourceStatusCreate. After verifying, it will be reset to
// module.ResourceStatusEnable or module.ResourceStatusDisable.
func CreateResource(name, apiVersion, description string, namespaced bool) error {
	resource := module.NewResource(name, apiVersion, description, namespaced)
	err := resource.Create()
	if err != nil {
		return err
	}

	// TODO: verify resource enable
	return nil
}

func verifyResourceEnable(resource *module.Resource) {
	//namespace := ""
	//if resource.Namespaced {
	//	namespace = "default"
	//}
	//url := getURL(resource, namespace)
	//
	//requests.New()
}

func getURL(resource *module.Resource, namespace string) string {
	urlList := make([]string, 0)
	if resource.APIVersion == "v1" {
		urlList = append(urlList, "api", resource.APIVersion)
	} else {
		urlList = append(urlList, "apis", resource.APIVersion)
	}

	if resource.Namespaced {
		urlList = append(urlList, "namespace", namespace)
	}
	urlList = append(urlList, resource.Name)
	return strings.Join(urlList, "/")
}
