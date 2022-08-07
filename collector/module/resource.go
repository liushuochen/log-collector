// Package module define basic structure.
// This file defined functions which about collector kubernetes resources.
package module

// Kubernetes resource status
const (
	// ResourceStatusCreate status indicates that resource just created and not starting verify check.
	ResourceStatusCreate = "creating"

	// ResourceStatusEnable status indicates that resource has been verified check, and it is available.
	ResourceStatusEnable = "enable"

	// ResourceStatusDisable status indicates that resource has been created, but kubernetes cluster do not register
	// it.
	ResourceStatusDisable = "disable"
)

// Resource structure defined kubernetes resource information. Although the information is the same for most resources
// in kubernetes, but doing so is compatible with user-defined resources(For example image resource in Openshift
// Container Platform). Another benefit is that it saves a lot of duplicate code.
// - Name: Kubernetes resource name(e.g. Pod, ConfigMap, etc). You can get it in "kubectl api-resources" command and in
//         "Kind" column.
// - APIVersion: Kubernetes resource api version(e.g. v1, apps/v1 etc). You can get it in "kubectl api-resources"
//               command and in "APIVERSION" column.
// - Namespaced: Is the resource a namespace resource or a cluster resource？If the resource is a namespace resource,
//               the value of the Namespaced is true.
// - Status: Resource status. Support ResourceStatusCreate, ResourceStatusEnable and ResourceStatusDisable.
// - Description: Resource description.
type Resource struct {
	Name        string `json:"name" gorm:"type:varchar(255);primary_key"`
	APIVersion  string `json:"api_version" gorm:"type:varchar(255);primary_key"`
	Namespaced  bool   `json:"namespaced" gorm:"not null"`
	Status      string `json:"status" gorm:"varchar(50);not null"`
	Description string `json:"description" gorm:"text"`
}

func initResource() {
	_ = db.CreateTable(&Resource{}, "InnoDB", "utf8", "utf8_general_ci")
}

// NewResource returns a pointer of Resource. Default status will be set to ResourceStatusCreate.
func NewResource(name, apiVersion, description string, namespaced bool) *Resource {
	return &Resource{
		Name:       name,
		APIVersion: apiVersion,
		Namespaced: namespaced,
		Status:     ResourceStatusCreate,
	}
}

// TableName method implement schema.Tabler interface.
func (r *Resource) TableName() string {
	return "resource"
}

// Create method used to create a data to resource table.
func (r *Resource) Create() error {
	return db.Create(r)
}

// Update method used to update a data to resource table.
func (r *Resource) Update() error {
	return db.Update(r)
}
