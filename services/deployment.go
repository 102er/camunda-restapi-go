package services

import (
	"bytes"
	"github.com/blossom102er/camunda-restapi-go/entitys"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type IDeployment interface {
	Create(deploymentCreate ReqDeploymentCreate) (deployment *entitys.RespDeploymentCreate, err error)
}

type Deployment struct {
	client *BaseRestApiService
}

func NewDeployment(client *BaseRestApiService) *Deployment {
	return &Deployment{
		client: client,
	}
}

// ReqDeploymentCreate a request to deployment create
type ReqDeploymentCreate struct {
	DeploymentName           string
	EnableDuplicateFiltering *bool
	DeployChangedOnly        *bool
	DeploymentSource         *string
	Resources                *os.File
}

// Create creates a deployment
func (d *Deployment) Create(deploymentCreate ReqDeploymentCreate) (deployment *entitys.RespDeploymentCreate, err error) {
	deployment = &entitys.RespDeploymentCreate{}
	var data []byte
	body := bytes.NewBuffer(data)
	w := multipart.NewWriter(body)

	if err = w.WriteField("deployment-name", deploymentCreate.DeploymentName); err != nil {
		return nil, err
	}

	if deploymentCreate.EnableDuplicateFiltering != nil {
		if err = w.WriteField("enable-duplicate-filtering", strconv.FormatBool(*deploymentCreate.EnableDuplicateFiltering)); err != nil {
			return nil, err
		}
	}

	if deploymentCreate.DeployChangedOnly != nil {
		if err = w.WriteField("deploy-changed-only", strconv.FormatBool(*deploymentCreate.DeployChangedOnly)); err != nil {
			return nil, err
		}
	}

	if deploymentCreate.DeploymentSource != nil {
		if err = w.WriteField("deployment-source", *deploymentCreate.DeploymentSource); err != nil {
			return nil, err
		}
	}

	resource := deploymentCreate.Resources
	var fw io.Writer

	defer func(resource *os.File) {
		err := resource.Close()
		if err != nil {

		}
	}(resource)
	if fw, err = w.CreateFormFile("data", resource.Name()); err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, resource); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	res, err := d.client.do(http.MethodPost, "/deployment/create", map[string]string{}, body, w.FormDataContentType())
	if err != nil {
		return nil, err
	}

	err = d.client.readJsonResponse(res, deployment)

	return deployment, err
}

func (d *Deployment) GetList(query map[string]string) (deployments []*entitys.RespDeploymentCreate, err error) {
	res, err := d.client.doGet("/deployment", query)
	if err != nil {
		return
	}

	err = d.client.readJsonResponse(res, &deployments)
	return
}
