package tests

import (
	"github.com/blossom102er/camunda-restapi-go/client"
	"github.com/blossom102er/camunda-restapi-go/entitys"
	"github.com/blossom102er/camunda-restapi-go/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	client.RegisterRestApiService(
		"http://localhost:8080/engine-rest",
		60)
	c = client.GetCamundaRestApiClient()
}

func TestActivateOrSuspendById(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		p2            bool
		expectedError error
	}{
		{
			Name:          "not found",
			params:        "1111",
			p2:            true,
			expectedError: nil,
		},
		{
			Name:          "ok",
			params:        "57838c25-b9ff-11eb-b976-0242ac110002",
			p2:            false,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.ProcessInstanceService.ActivateOrSuspendById(tt.params, tt.p2)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestModify(t *testing.T) {
	assertions := assert.New(t)
	a := map[string]entitys.Variable{
		"assignee": {
			Value: "demo",
			Type:  "String",
		},
	}
	var tests = []struct {
		Name          string
		params        string
		p2            *services.ReqModifyProcessInstance
		expectedError error
	}{
		{
			Name:   "not found",
			params: "07269a85-ba12-11eb-b976-0242ac110002",
			p2: &services.ReqModifyProcessInstance{
				SkipCustomListeners: false,
				SkipIoMappings:      false,
				Instructions: []services.Instruction{
					{
						Type:       "startBeforeActivity",
						ActivityId: "Activity_1eeq9ev",
						Variables:  &a,
					},
					{
						Type:       "cancel",
						ActivityId: "Activity_0h041lu",
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.ProcessInstanceService.Modify(tt.params, tt.p2)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestGetVariables(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expectedError error
	}{
		{
			Name:          "ok",
			params:        "57838c25-b9ff-11eb-b976-0242ac110002",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m, err := c.ProcessInstanceService.GetVariables(tt.params)
			assertions.Equal(tt.expectedError, err)
			t.Log(m)
		})
	}
}

func TestGetActivity(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expectedError error
	}{
		{
			Name:          "ok",
			params:        "57838c25-b9ff-11eb-b976-0242ac110002",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m, err := c.ProcessInstanceService.GetActivityInstances(tt.params)
			assertions.Equal(tt.expectedError, err)
			t.Log(m)
		})
	}
}
