package client

import (
	"github.com/blossom102er/camunda-restapi-go/services"
	"time"
)

// https://docs.camunda.org/manual/7.15/reference/rest/
var restApiClient *CamundaRestApiClient

type CamundaRestApiClient struct {
	ProcessDefinitionService services.IProcessDefinition
	ProcessInstanceService   services.IProcessInstance
	UserTaskService          services.IUserTask
	DeploymentService        services.IDeployment
	MessageService           services.IMessage
	ExternalTaskService      services.IExternalTask
}

func newCamundaRestApiClient(endpointUrl string, timeoutSec time.Duration) *CamundaRestApiClient {
	camundaRestApiClient := &CamundaRestApiClient{}
	baseBaseRestApiService := services.NewBaseRestApiService(endpointUrl, timeoutSec)
	camundaRestApiClient.ProcessDefinitionService = services.NewProcessDefinition(baseBaseRestApiService)
	camundaRestApiClient.UserTaskService = services.NewUserTask(baseBaseRestApiService)
	camundaRestApiClient.DeploymentService = services.NewDeployment(baseBaseRestApiService)
	camundaRestApiClient.ProcessInstanceService = services.NewProcessInstance(baseBaseRestApiService)
	camundaRestApiClient.ExternalTaskService = services.NewExternalTask(baseBaseRestApiService)
	camundaRestApiClient.MessageService = services.NewMessage(baseBaseRestApiService)
	return camundaRestApiClient
}

func GetCamundaRestApiClient() *CamundaRestApiClient {
	return restApiClient
}

func RegisterRestApiService(endpointUrl string, timeoutSec time.Duration) {
	restApiClient = newCamundaRestApiClient(endpointUrl, timeoutSec)
}
