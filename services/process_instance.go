package services

import (
	"github.com/blossom102er/camunda-restapi-go/entitys"
)

type IProcessInstance interface {
	GetList(req *QueryProcessInstanceParams) ([]entitys.RespProcessInstance, error)
	ActivateOrSuspendById(id string, suspended bool) error
	Delete(id string) error
	Modify(id string, req *ReqModifyProcessInstance) error
	GetVariables(id string) (map[string]entitys.Variable, error)
	GetActivityInstances(id string) (entitys.RespActivityProcessInstance, error)
}

type ProcessInstance struct {
	client *BaseRestApiService
}

func NewProcessInstance(client *BaseRestApiService) *ProcessInstance {
	return &ProcessInstance{
		client: client,
	}
}

type QueryProcessInstanceParams struct {
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	Active               bool   `json:"active"`
	Suspended            bool   `json:"suspended"`
}

func (p *ProcessInstance) GetList(req *QueryProcessInstanceParams) ([]entitys.RespProcessInstance, error) {
	if req == nil {
		req = &QueryProcessInstanceParams{}
	}
	query := make(map[string]string)
	if req.Active {
		query["active"] = "true"
	}
	if req.Suspended {
		query["suspended"] = "true"
	}
	if len(req.ProcessDefinitionKey) > 0 {
		query["processDefinitionKey"] = req.ProcessDefinitionKey
	}
	res, err := p.client.doGet("/process-instance", query)
	if err != nil {
		return nil, err
	}
	var pi []entitys.RespProcessInstance
	err = p.client.readJsonResponse(res, &pi)
	return pi, err
}

func (p *ProcessInstance) Delete(id string) error {
	_, err := p.client.doDelete("/process-instance/"+id, nil)
	return err
}

func (p *ProcessInstance) ActivateOrSuspendById(id string, suspended bool) error {
	err := p.client.doPutJson("/process-instance/"+id+"/suspended", nil, struct {
		Suspended bool `json:"suspended"`
	}{
		suspended,
	})
	return err
}

type ReqModifyProcessInstance struct {
	SkipCustomListeners bool          `json:"skipCustomListeners"`
	SkipIoMappings      bool          `json:"skipIoMappings"`
	Instructions        []Instruction `json:"instructions"`
}

type Instruction struct {
	Type       string                       `json:"type"`
	ActivityId string                       `json:"activityId"`
	Variables  *map[string]entitys.Variable `json:"variables,omitempty"`
}

func (p *ProcessInstance) Modify(id string, req *ReqModifyProcessInstance) error {
	_, err := p.client.doPostJson("/process-instance/"+id+"/modification", nil, req)
	return err
}

func (p *ProcessInstance) GetVariables(id string) (map[string]entitys.Variable, error) {
	res, err := p.client.doGet("/process-instance/"+id+"/variables", nil)
	if err != nil {
		return nil, err
	}
	var m map[string]entitys.Variable
	err = p.client.readJsonResponse(res, &m)
	return m, err
}

func (p *ProcessInstance) GetActivityInstances(id string) (entitys.RespActivityProcessInstance, error) {
	res, err := p.client.doGet("/process-instance/"+id+"/activity-instances", nil)
	if err != nil {
		return entitys.RespActivityProcessInstance{}, err
	}
	var r entitys.RespActivityProcessInstance
	err = p.client.readJsonResponse(res, &r)
	return r, err
}
