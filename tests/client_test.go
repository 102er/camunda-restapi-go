package tests

import (
	"github.com/blossom102er/camunda-restapi-go/client"
	"github.com/blossom102er/camunda-restapi-go/services"
	"log"
	"testing"
)

func init() {
	client.RegisterRestApiService(
		"http://localhost:8080/engine-rest", //camunda部署的地址
		60)
	c = client.GetCamundaRestApiClient()
}

func TestExternalTaskRoutine(t *testing.T) {
	create, err := c.DeploymentService.Create(services.ReqDeploymentCreate{})
	if err != nil {
		return
	}
	log.Print(create)
}
