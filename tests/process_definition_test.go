package tests

import (
	"fmt"
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

func TestGet(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        *services.QueryProcessDefinitionBy
		expected      *entitys.RespProcessDefinition
		expectedError error
	}{
		{
			Name:          "not found",
			params:        &services.QueryProcessDefinitionBy{Id: "111"},
			expected:      nil,
			expectedError: fmt.Errorf("not found"),
		},
		{
			Name:   "input id test，200",
			params: &services.QueryProcessDefinitionBy{Id: "db_auth_workflow:1:ee59b29e-a249-11eb-8191-0242ac110002"},
			expected: &entitys.RespProcessDefinition{
				Id:        "db_auth_workflow:1:ee59b29e-a249-11eb-8191-0242ac110002",
				Key:       "db_auth_workflow",
				Name:      "",
				Version:   1,
				Suspended: false,
			},
			expectedError: nil,
		},
		{
			Name:   "input key test，200",
			params: &services.QueryProcessDefinitionBy{Key: "db_auth_workflow"},
			expected: &entitys.RespProcessDefinition{
				Id:        "db_auth_workflow:5:adf1296a-a356-11eb-8191-0242ac110002",
				Key:       "db_auth_workflow",
				Name:      "",
				Version:   5,
				Suspended: false,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.ProcessDefinitionService.Get(*tt.params)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}

func TestGetList(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        map[string]string
		expected      []*entitys.RespProcessDefinition
		expectedError error
	}{
		{
			Name:   "get all process definition",
			params: nil,
			expected: []*entitys.RespProcessDefinition{{
				Id:        "db_auth_workflow:5:adf1296a-a356-11eb-8191-0242ac110002",
				Key:       "db_auth_workflow",
				Name:      "",
				Version:   5,
				Suspended: false,
			}, {
				Id:        "db_auth_workflow_v3:8:e73b6ff0-a7c5-11eb-8191-0242ac110002",
				Key:       "db_auth_workflow_v3",
				Name:      "",
				Version:   8,
				Suspended: false,
			},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.ProcessDefinitionService.GetList(tt.params)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}

func TestDelete(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        *services.QueryProcessDefinitionBy
		expected      *entitys.RespProcessDefinition
		expectedError error
	}{
		{
			Name:          "not found",
			params:        &services.QueryProcessDefinitionBy{Key: "test_id"},
			expected:      nil,
			expectedError: fmt.Errorf("not found"),
		},
		{
			Name:          "input id test，200",
			params:        &services.QueryProcessDefinitionBy{Id: "test_id:2:3408c114-b9da-11eb-b976-0242ac110002"},
			expected:      nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.ProcessDefinitionService.Delete(*tt.params, nil)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestStartInstance(t *testing.T) {
	assertions := assert.New(t)
	a := map[string]entitys.Variable{
		"createor": {
			Value: "demo",
			Type:  "String",
		}, "owner": {
			Value: "demo",
			Type:  "String",
		},
	}
	var tests = []struct {
		Name          string
		params        *services.QueryProcessDefinitionBy
		params2       *services.ReqStartInstance
		expected      *entitys.RespStartedProcessDefinition
		expectedError *services.Error
	}{
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "111",
			},
			expected:      nil,
			expectedError: nil,
		},
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "222",
			},
			expected:      nil,
			expectedError: nil,
		},
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "333",
			},
			expected:      nil,
			expectedError: nil,
		},
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "444",
			},
			expected:      nil,
			expectedError: nil,
		},
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "666",
			},
			expected:      nil,
			expectedError: nil,
		},
		{
			Name:   "test start process instance by key",
			params: &services.QueryProcessDefinitionBy{Key: "external_test"},
			params2: &services.ReqStartInstance{
				Variables:             &a,
				WithVariablesInReturn: false,
				BusinessKey:           "555",
			},
			expected:      nil,
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := c.ProcessDefinitionService.StartInstance(*tt.params, *tt.params2)
			assertions.Equal(tt.expectedError, err)
			//assertions.Equal(tt.expected, ret)
		})
	}
}
