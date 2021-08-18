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

func TestUserTaskCreate(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        *services.ReqUserTaskCreate
		expectedError error
	}{
		{
			Name: "add task assignee",
			params: &services.ReqUserTaskCreate{
				Id:           "111223422342",
				Name:         "test task 4",
				Description:  "test",
				Assignee:     "boss01",
				Owner:        "boss01",
				ParentTaskId: "578560ec-b9ff-11eb-b976-0242ac110002",
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.UserTaskService.Create(tt.params)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestUserTaskGet(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expected      *entitys.RespUserTask
		expectedError error
	}{
		{
			Name:          "not found",
			params:        "1c9711cbc-a808-11eb-8191-0242ac110002",
			expected:      nil,
			expectedError: fmt.Errorf("not found"),
		},
		{
			Name:   "申请表单任务",
			params: "5270ecc9-b91b-11eb-b976-0242ac110002",
			expected: &entitys.RespUserTask{
				Id:                  "5270ecc9-b91b-11eb-b976-0242ac110002",
				Name:                "申请表单填写",
				Assignee:            "demo",
				Created:             "2021-05-20T03:27:36.000+0000",
				Due:                 "",
				ParentTaskId:        "",
				Priority:            50,
				ProcessDefinitionId: "db_auth_workflow_v3:8:e73b6ff0-a7c5-11eb-8191-0242ac110002",
				ProcessInstanceId:   "524f5b04-b91b-11eb-b976-0242ac110002",
				TaskDefinitionKey:   "Activity_03es9gt",
				Suspended:           false,
				FormKey:             ""},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.UserTaskService.Get(tt.params)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}

func TestTaskGetList(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        *services.ReqUserTaskGetList
		expected      int
		expectedError error
	}{
		{
			Name:          "all test",
			params:        &services.ReqUserTaskGetList{ProcessDefinitionKey: "more_node"},
			expected:      2,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.UserTaskService.GetList(tt.params)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, len(ret))
			t.Log(ret)
		})
	}
}

func TestGetFormKey(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expected      *entitys.RespUserTaskForm
		expectedError error
	}{
		{
			Name:   "test form key",
			params: "8095c994-b934-11eb-b976-0242ac110002",
			expected: &entitys.RespUserTaskForm{
				ContextPath: "",
				Key:         "test_form_key",
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.UserTaskService.GetFormKey(tt.params)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}

func TestSetAssignee(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		params2       string
		expectedError error
	}{
		{
			Name:          "change assignee",
			params:        "8095c994-b934-11eb-b976-0242ac110002",
			params2:       "demo",
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.UserTaskService.SetAssignee(tt.params, tt.params2)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestClaim(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		params2       string
		expectedError error
	}{
		{
			Name:          "change assignee",
			params:        "4abfd279-b9e0-11eb-b976-0242ac110002",
			params2:       "boss01",
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.UserTaskService.Claim(tt.params, tt.params2)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestUnClaim(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expectedError error
	}{
		{
			Name:          "change assignee",
			params:        "8095c994-b934-11eb-b976-0242ac110002",
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.UserTaskService.UnClaim(tt.params)
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestComplete(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		params        string
		expectedError error
	}{
		{
			Name:          "complete task",
			params:        "7f922cdc-b9fe-11eb-b976-0242ac110002",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := c.UserTaskService.Complete(tt.params, &services.ReqUserTaskComplete{
				Variables: map[string]entitys.Variable{
					"mark": {
						Value: "ok",
						Type:  "String",
					},
				},
			})
			assertions.Equal(tt.expectedError, err)
		})
	}
}

func TestCreateComment(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		Id            string
		Message       string
		expected      *entitys.RespUserTaskComment
		expectedError error
	}{
		{
			Name:    "all test",
			Id:      "5270ecc9-b91b-11eb-b976-0242ac110002",
			Message: "test comment",
			expected: &entitys.RespUserTaskComment{
				UserName: "demo", Message: "test comment",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.UserTaskService.CreateComment(tt.Id, "demo", tt.Message)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}

func TestGetCommentList(t *testing.T) {
	assertions := assert.New(t)
	var tests = []struct {
		Name          string
		Id            string
		Message       string
		expected      []entitys.RespUserTaskComment
		expectedError error
	}{
		{
			Name:          "all test",
			Id:            "5270ecc9-b91b-11eb-b976-0242ac110002",
			expected:      nil,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ret, err := c.UserTaskService.GetCommentList(tt.Id)
			t.Log(err)
			assertions.Equal(tt.expectedError, err)
			assertions.Equal(tt.expected, ret)
		})
	}
}
