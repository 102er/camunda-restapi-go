package tests

import (
	"github.com/blossom102er/camunda-restapi-go/client"
	"github.com/blossom102er/camunda-restapi-go/services"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var c *client.CamundaRestApiClient

func init() {
	client.RegisterRestApiService(
		"http://localhost:8080/engine-rest",
		60)
	c = client.GetCamundaRestApiClient()
}

func TestDeploymentCreate(t *testing.T) {
	file, err := os.Open("easy.bpmn")
	if err != nil {
		t.Error(err)
	}
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        *services.ReqDeploymentCreate
		expectedError error
	}{
		{
			Name: "deployment bpmn",
			params: &services.ReqDeploymentCreate{
				DeploymentName: "aName",
				Resources:      file,
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			res, err := c.DeploymentService.Create(*tt.params)
			assertions.Equal(tt.expectedError, err)
			t.Log(res)
		})
	}
}
