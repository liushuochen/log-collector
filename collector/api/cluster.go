// Package api contains API functions which called by route package.
// This file defined methods which about cluster resource.
package api

import (
	"io/ioutil"
	"log-collector/api/req"
	"log-collector/api/resp"
	"log-collector/controller"
	"log-collector/exception"
	"log-collector/utils/collector_uuid"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type Cluster struct{}

// CreateWithKubeConfigContent method used to create a Cluster data by given KubeConfig content string.
func (cluster *Cluster) CreateWithKubeConfigContent(c *gin.Context) {
	var request req.CreateClusterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	cluster.create(c, request)
}

// CreateWithKubeConfigFile method used to create a Cluster data by given KubeConfig file.
func (cluster *Cluster) CreateWithKubeConfigFile(c *gin.Context) {
	var request req.CreateClusterRequest
	request.IP = c.PostForm("ip")
	request.Description = c.PostForm("description")

	fileHeader, err := c.FormFile("kube_config")
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	request.KubeConfig = string(content)
	cluster.create(c, request)
}

func (cluster *Cluster) create(c *gin.Context, request req.CreateClusterRequest) {
	validate := validator.New(&validator.Config{TagName: "binding"})
	err := validate.Struct(request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	uuid := collector_uuid.New()
	response, err := controller.ClusterCreate(uuid, request.IP, request.Description, request.KubeConfig)
	if err != nil {
		switch err.(type) {
		case exception.EmptyClusterError:
			resp.SendResponse(c, resp.Forbidden, err.Error())
		default:
			resp.SendResponse(c, resp.InternalError, err.Error())
		}
		return
	}

	resp.SendResponse(c, resp.Created, response)
	return
}

// Delete method used to delete a cluster resource.
func (cluster *Cluster) Delete(c *gin.Context) {
	var request req.DeleteClusterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	validate := validator.New(&validator.Config{TagName: "binding"})
	err = validate.Struct(request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	response, err := controller.ClusterDelete(request.UUID)
	if err != nil {
		switch err.(type) {
		case exception.ClusterNotFoundError:
			resp.SendResponse(c, resp.ResourceNotFound, err.Error())
		default:
			resp.SendResponse(c, resp.InternalError, err.Error())
		}
		return
	}

	resp.SendResponse(c, resp.Ok, response)
}

// Edit method used to edit a cluster resource.
func (cluster *Cluster) Edit(c *gin.Context) {
	var request req.EditClusterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	validate := validator.New(&validator.Config{TagName: "binding"})
	err = validate.Struct(request)
	if err != nil {
		resp.SendResponse(c, resp.BadRequest, err.Error())
		return
	}

	if request.Name == "" && request.Description == "" {
		resp.SendResponse(c, resp.BadRequest, "must designated new value of name or description for cluster")
		return
	}

	response, err := controller.ClusterEdit(request.UUID, request.Name, request.Description)
	if err != nil {
		switch err.(type) {
		case exception.ClusterNotFoundError:
			resp.SendResponse(c, resp.ResourceNotFound, err.Error())
		default:
			resp.SendResponse(c, resp.InternalError, err.Error())
		}
		return
	}

	resp.SendResponse(c, resp.Ok, response)
}
