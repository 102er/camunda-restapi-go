package services

import (
	"github.com/blossom102er/camunda-restapi-go/entitys"
	"io/ioutil"
)

type IProcessDefinition interface {
	GetActivityInstanceStatistics(by QueryProcessDefinitionBy, query map[string]string) (statistic []*entitys.RespActivityInstanceStatistics, err error)
	GetDiagram(by QueryProcessDefinitionBy) (data []byte, err error)
	GetStartFormVariables(by QueryProcessDefinitionBy, query map[string]string) (variables map[string]entitys.Variable, err error)
	GetList(query map[string]string) (processDefinitions []*entitys.RespProcessDefinition, err error)
	GetProcessInstanceStatistics(query map[string]string) (statistic []*entitys.RespInstanceStatistics, err error)
	GetXML(by QueryProcessDefinitionBy) (resp *entitys.RespBPMNProcessDefinition, err error)
	Get(by QueryProcessDefinitionBy) (processDefinition *entitys.RespProcessDefinition, err error)
	StartInstance(by QueryProcessDefinitionBy, req ReqStartInstance) (processDefinition *entitys.RespStartedProcessDefinition, err error)
	ActivateOrSuspendById(by QueryProcessDefinitionBy, req ReqActivateOrSuspendById) error
	ActivateOrSuspendByKey(req ReqActivateOrSuspendByKey) error
	Delete(by QueryProcessDefinitionBy, query map[string]string) error
	RestartProcessInstance(id string, req ReqRestartInstance) error
	RestartProcessInstanceAsync(id string, req ReqRestartInstance) (resp *entitys.RespBatch, err error)
}

type ReqRestartInstance struct {
	ProcessInstanceIds           string                  `json:"processInstanceIds,omitempty"`
	HistoricProcessInstanceQuery string                  `json:"historicProcessInstanceQuery,omitempty"`
	StartInstructions            *[]ReqStartInstructions `json:"startInstructions,omitempty"`
	SkipCustomListeners          bool                    `json:"skipCustomListeners,omitempty"`
	SkipIoMappings               bool                    `json:"skipIoMappings,omitempty"`
	InitialVariables             bool                    `json:"initialVariables,omitempty"`
	WithVariablesInReturn        bool                    `json:"withoutBusinessKey,omitempty"`
}

type ReqActivateOrSuspendById struct {
	Suspended               bool `json:"suspended,omitempty"`
	IncludeProcessInstances bool `json:"includeProcessInstances,omitempty"`
	// e.g., 2013-01-23T14:42:45
	ExecutionDate *entitys.Time `json:"executionDate,omitempty"`
}

type ReqActivateOrSuspendByKey struct {
	ProcessDefinitionKey    string `json:"processDefinitionKey"`
	Suspended               *bool  `json:"suspended,omitempty"`
	IncludeProcessInstances *bool  `json:"includeProcessInstances,omitempty"`
	// e.g., 2013-01-23T14:42:45
	ExecutionDate *entitys.Time `json:"executionDate,omitempty"`
}

type ReqStartInstance struct {
	Variables             *map[string]entitys.Variable `json:"variables,omitempty"`
	BusinessKey           string                       `json:"businessKey,omitempty"`
	CaseInstanceId        string                       `json:"caseInstanceId,omitempty"`
	StartInstructions     *[]ReqStartInstructions      `json:"startInstructions,omitempty"`
	SkipCustomListeners   bool                         `json:"skipCustomListeners,omitempty"`
	SkipIoMappings        bool                         `json:"skipIoMappings,omitempty"`
	WithVariablesInReturn bool                         `json:"withVariablesInReturn,omitempty"`
}

type ReqStartInstructions struct {
	Type         string                          `json:"type"`
	ActivityId   string                          `json:"activityId,omitempty"`
	TransitionId string                          `json:"transitionId,omitempty"`
	Variables    *map[string]entitys.VariableSet `json:"variables,omitempty"`
}

// ProcessDefinition a client for ProcessDefinition
type ProcessDefinition struct {
	client *BaseRestApiService
	path   string
}

func NewProcessDefinition(client *BaseRestApiService) *ProcessDefinition {
	return &ProcessDefinition{
		client: client,
		path:   "/process-definition",
	}
}

func (p *ProcessDefinition) GetActivityInstanceStatistics(by QueryProcessDefinitionBy, query map[string]string) (statistic []*entitys.RespActivityInstanceStatistics, err error) {
	res, err := p.client.doGet(p.path+by.String()+"/statistics", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &statistic)
	return
}

func (p *ProcessDefinition) GetDiagram(by QueryProcessDefinitionBy) (data []byte, err error) {
	res, err := p.client.doGet(p.path+by.String()+"/diagram", map[string]string{})
	if err != nil {
		return
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (p *ProcessDefinition) GetStartFormVariables(by QueryProcessDefinitionBy, query map[string]string) (variables map[string]entitys.Variable, err error) {
	res, err := p.client.doGet(p.path+by.String()+"/form-variables", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &variables)
	return
}

func (p *ProcessDefinition) GetList(query map[string]string) (processDefinitions []*entitys.RespProcessDefinition, err error) {
	res, err := p.client.doGet(p.path, query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &processDefinitions)
	return
}

func (p *ProcessDefinition) GetProcessInstanceStatistics(query map[string]string) (statistic []*entitys.RespInstanceStatistics, err error) {
	res, err := p.client.doGet(p.path+"/statistics", query)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, &statistic)
	return
}

func (p *ProcessDefinition) GetXML(by QueryProcessDefinitionBy) (resp *entitys.RespBPMNProcessDefinition, err error) {
	resp = &entitys.RespBPMNProcessDefinition{}
	res, err := p.client.doGet(p.path+by.String()+"/xml", map[string]string{})
	if err != nil {
		return
	}
	err = p.client.readJsonResponse(res, &resp)
	return
}

func (p *ProcessDefinition) Get(by QueryProcessDefinitionBy) (processDefinition *entitys.RespProcessDefinition, err error) {
	res, err := p.client.doGet(p.path+by.String(), map[string]string{})
	if err != nil {
		return
	}
	processDefinition = &entitys.RespProcessDefinition{}
	err = p.client.readJsonResponse(res, &processDefinition)
	return
}

func (p *ProcessDefinition) StartInstance(by QueryProcessDefinitionBy, req ReqStartInstance) (processDefinition *entitys.RespStartedProcessDefinition, err error) {
	res, err := p.client.doPostJson(p.path+by.String()+"/start", map[string]string{}, &req)
	if err != nil {
		return
	}

	processDefinition = &entitys.RespStartedProcessDefinition{}
	err = p.client.readJsonResponse(res, processDefinition)
	return
}

func (p *ProcessDefinition) ActivateOrSuspendById(by QueryProcessDefinitionBy, req ReqActivateOrSuspendById) error {
	return p.client.doPutJson(p.path+by.String()+"/suspended", map[string]string{}, &req)
}

func (p *ProcessDefinition) ActivateOrSuspendByKey(req ReqActivateOrSuspendByKey) error {
	return p.client.doPutJson(p.path+"/suspended", map[string]string{}, &req)
}

func (p *ProcessDefinition) Delete(by QueryProcessDefinitionBy, query map[string]string) error {
	_, err := p.client.doDelete(p.path+by.String(), query)
	return err
}

func (p *ProcessDefinition) RestartProcessInstance(id string, req ReqRestartInstance) error {
	_, err := p.client.doPostJson(p.path+"/"+id+"/restart", map[string]string{}, &req)
	return err
}

func (p *ProcessDefinition) RestartProcessInstanceAsync(id string, req ReqRestartInstance) (resp *entitys.RespBatch, err error) {
	resp = &entitys.RespBatch{}
	res, err := p.client.doPostJson(p.path+"/"+id+"/restart-async", map[string]string{}, &req)
	if err != nil {
		return
	}

	err = p.client.readJsonResponse(res, resp)
	return
}

type QueryProcessDefinitionBy struct {
	Id       string
	Key      string
	TenantId string
}

func (q *QueryProcessDefinitionBy) String() string {
	switch {
	case len(q.Key) != 0 && len(q.TenantId) != 0:
		return "/key/" + q.Key + "/tenant-id/" + q.TenantId
	case len(q.Key) != 0:
		return "/key/" + q.Key
	default:
		return "/" + q.Id
	}
}
