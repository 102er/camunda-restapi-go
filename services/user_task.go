package services

import (
	"encoding/json"
	"fmt"
	"github.com/blossom102er/camunda-restapi-go/entitys"
	"log"
	"time"
)

type IUserTask interface {
	Create(req *ReqUserTaskCreate) error
	Get(id string) (*entitys.RespUserTask, error)
	Complete(id string, query *ReqUserTaskComplete) error
	SubmitForm(id string, query *ReqUserTaskComplete) error
	GetList(query *ReqUserTaskGetList) ([]entitys.RespUserTask, error)
	GetFormKey(id string) (*entitys.RespUserTaskForm, error)
	SetAssignee(id, userName string) error
	Claim(id, userName string) error
	UnClaim(id string) error
	CreateComment(id, userName, message string) (*entitys.RespUserTaskComment, error)
	GetCommentList(id string) ([]entitys.RespUserTaskComment, error)
}

type UserTask struct {
	client *BaseRestApiService
}

func NewUserTask(client *BaseRestApiService) *UserTask {
	return &UserTask{
		client: client,
	}
}

type ReqUserTaskGetList struct {
	// Restrict to tasks that belong to process instances with the given id.
	ProcessInstanceId string `json:"processInstanceId,omitempty"`
	// Restrict to tasks that belong to process instances with the given business key.
	ProcessInstanceBusinessKey string `json:"processInstanceBusinessKey,omitempty"`
	// Restrict to tasks that belong to process instances with one of the give business keys. The keys need to be in a comma-separated list.
	ProcessInstanceBusinessKeyIn []string `json:"processInstanceBusinessKeyIn,omitempty"`
	// Restrict to tasks that have a process instance business key that has the parameter value as a substring.
	ProcessInstanceBusinessKeyLike string `json:"processInstanceBusinessKeyLike,omitempty"`
	// Restrict to tasks that belong to a process definition with the given id.
	ProcessDefinitionId string `json:"processDefinitionId,omitempty"`
	// Restrict to tasks that belong to a process definition with the given key.
	ProcessDefinitionKey string `json:"processDefinitionKey,omitempty"`
	// Restrict to tasks that belong to a process definition with one of the given keys. The keys need to be in a comma-separated list.
	ProcessDefinitionKeyIn []string `json:"processDefinitionKeyIn,omitempty"`
	// Restrict to tasks that belong to a process definition with the given name.
	ProcessDefinitionName string `json:"processDefinitionName,omitempty"`
	// Restrict to tasks that have a process definition name that has the parameter value as a substring.
	ProcessDefinitionNameLike string `json:"processDefinitionNameLike,omitempty"`
	// Restrict to tasks that belong to an execution with the given id.
	ExecutionId string `json:"executionId,omitempty"`
	// Restrict to tasks that belong to case instances with the given id.
	CaseInstanceId string `json:"caseInstanceId,omitempty"`
	// Restrict to tasks that belong to case instances with the given business key.
	CaseInstanceBusinessKey string `json:"caseInstanceBusinessKey,omitempty"`
	// Restrict to tasks that have a case instance business key that has the parameter value as a substring.
	CaseInstanceBusinessKeyLike string `json:"caseInstanceBusinessKeyLike,omitempty"`
	// Restrict to tasks that belong to a case definition with the given id.
	CaseDefinitionId string `json:"caseDefinitionId,omitempty"`
	// Restrict to tasks that belong to a case definition with the given key.
	CaseDefinitionKey string `json:"caseDefinitionKey,omitempty"`
	// Restrict to tasks that belong to a case definition with the given name.
	CaseDefinitionName string `json:"caseDefinitionName,omitempty"`
	// Restrict to tasks that have a case definition name that has the parameter value as a substring.
	CaseDefinitionNameLike string `json:"caseDefinitionNameLike,omitempty"`
	// Restrict to tasks that belong to a case execution with the given id.
	CaseExecutionId string `json:"caseExecutionId,omitempty"`
	// Only include tasks which belong to one of the passed and comma-separated activity instance ids.
	ActivityInstanceIdIn []string `json:"activityInstanceIdIn,omitempty"`
	// Only include tasks which belong to one of the passed and comma-separated tenant ids.
	TenantIdIn []string `json:"tenantIdIn,omitempty"`
	// Only include tasks which belong to no tenant. Value may only be true, as false is the default behavior.
	WithoutTenantId string `json:"withoutTenantId,omitempty"`
	// Restrict to tasks that the given user is assigned to.
	Assignee string `json:"assignee,omitempty"`
	// Restrict to tasks that the user described by the given expression is assigned to. See the user guide for more information on available functions.
	AssigneeExpression string `json:"assigneeExpression,omitempty"`
	// Restrict to tasks that have an assignee that has the parameter value as a substring.
	AssigneeLike string `json:"assigneeLike,omitempty"`
	// Restrict to tasks that have an assignee that has the parameter value described by the given expression as a substring. See the user guide for more information on available functions.
	AssigneeLikeExpression string `json:"assigneeLikeExpression,omitempty"`
	// Restrict to tasks that the given user owns.
	Owner string `json:"owner,omitempty"`
	// Restrict to tasks that the user described by the given expression owns. See the user guide for more information on available functions.
	OwnerExpression string `json:"ownerExpression,omitempty"`
	// Only include tasks that are offered to the given group.
	CandidateGroup string `json:"candidateGroup,omitempty"`
	// Only include tasks that are offered to the group described by the given expression. See the user guide for more information on available functions.
	CandidateGroupExpression string `json:"candidateGroupExpression,omitempty"`
	// Only include tasks that are offered to the given user or to one of his groups.
	CandidateUser string `json:"candidateUser,omitempty"`
	// Only include tasks that are offered to the user described by the given expression. See the user guide for more information on available functions.
	CandidateUserExpression string `json:"candidateUserExpression,omitempty"`
	// Also include tasks that are assigned to users in candidate queries. Default is to only include tasks that are not assigned to any user if you query by candidate user or group(s).
	IncludeAssignedTasks bool `json:"includeAssignedTasks,omitempty"`
	// Only include tasks that the given user is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee).
	InvolvedUser string `json:"involvedUser,omitempty"`
	// Only include tasks that the user described by the given expression is involved in. A user is involved in a task if an identity link exists between task and user (e.g., the user is the assignee). See the user guide for more information on available functions.
	InvolvedUserExpression string `json:"involvedUserExpression,omitempty"`
	// If set to true, restricts the query to all tasks that are assigned.
	Assigned bool `json:"assigned,omitempty"`
	// If set to true, restricts the query to all tasks that are unassigned.
	Unassigned bool `json:"unassigned,omitempty"`
	// Restrict to tasks that have the given key.
	TaskDefinitionKey string `json:"taskDefinitionKey,omitempty"`
	// Restrict to tasks that have one of the given keys. The keys need to be in a comma-separated list.
	TaskDefinitionKeyIn []string `json:"taskDefinitionKeyIn,omitempty"`
	// Restrict to tasks that have a key that has the parameter value as a substring.
	TaskDefinitionKeyLike string `json:"taskDefinitionKeyLike,omitempty"`
	// Restrict to tasks that have the given name.
	Name string `json:"name,omitempty"`
	// Restrict to tasks that do not have the given name.
	NameNotEqual string `json:"nameNotEqual,omitempty"`
	// Restrict to tasks that have a name with the given parameter value as substring.
	NameLike string `json:"nameLike,omitempty"`
	// Restrict to tasks that do not have a name with the given parameter value as substring.
	NameNotLike string `json:"nameNotLike,omitempty"`
	Priority    int64  `json:"priority,omitempty"`
	// Restrict to tasks that are due on the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueDate time.Time `json:"dueDate,omitempty"`
	// Restrict to tasks that are due on the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueDateExpression time.Time `json:"dueDateExpression,omitempty"`
	// Restrict to tasks that are due after the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueAfter time.Time `json:"dueAfter,omitempty"`
	// Restrict to tasks that are due after the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueAfterExpression string `json:"dueAfterExpression,omitempty"`
	// Restrict to tasks that are due before the given date. The date must have the format yyyy-MM-dd'T'HH:mm:ss, e.g., 2013-01-23T14:42:45.
	DueBefore time.Time `json:"dueBefore,omitempty"`
	// Restrict to tasks that are due before the date described by the given expression. See the user guide for more information on available functions. The expression must evaluate to a java.util.Date or org.joda.time.DateTime object.
	DueBeforeExpression string `json:"dueBeforeExpression,omitempty"`
	// Restrict to tasks that are offered to any of the given candidate groups.Takes a comma-separated list of group names, so for example developers, support, sales.
	CandidateGroups []string `json:"candidateGroups,omitempty"`
	// Restrict to tasks that are offered to any of the candidate groups described by the given expression.See the user guide for more information on available functions.The expression must evaluate to java.util.List of Strings.
	CandidateGroupsExpression []string `json:"candidateGroupsExpression,omitempty"`
	// Only include tasks which have a candidate group.Value may only be true, as false is the default behavior.
	WithCandidateGroups bool `json:"withCandidateGroups,omitempty"`
	// Only include tasks which have a candidate user.Value may only be true, as false is the default behavior.
	WithCandidateUsers bool `json:"withCandidateUsers,omitempty"`
	// Only include active tasks.Value may only be true, as false is the default behavior.
	Active bool `json:"active,omitempty"`
	// Only include suspended tasks.Value may only be true, as false is the default behavior.
	Suspended        bool                       `json:"suspended,omitempty"`
	TaskVariables    []VariableFilterExpression `json:"taskVariables,omitempty"`
	ProcessVariables []VariableFilterExpression `json:"processVariables,omitempty"`
	// Restrict query to all tasks that are sub tasks of the given task.Takes a task id.
	ParentTaskId string `json:"parentTaskId,omitempty"`
}

// MarshalJSON marshal to json
func (q *ReqUserTaskGetList) MarshalJSON() ([]byte, error) {
	type Alias ReqUserTaskGetList
	return json.Marshal(&struct {
		*Alias
		DueDate           string `json:"dueDate,omitempty"`
		DueDateExpression string `json:"dueDateExpression,omitempty"`
		DueAfter          string `json:"dueAfter,omitempty"`
		DueBefore         string `json:"dueBefore,omitempty"`
	}{
		Alias:             (*Alias)(q),
		DueDate:           toCamundaTime(q.DueDate),
		DueDateExpression: toCamundaTime(q.DueDateExpression),
		DueAfter:          toCamundaTime(q.DueAfter),
		DueBefore:         toCamundaTime(q.DueBefore),
	})
}

type variableFilterExpressionOperator string

type VariableFilterExpression struct {
	Name     string                           `json:"name"`
	Operator variableFilterExpressionOperator `json:"operator"`
	Value    string                           `json:"value"`
}

type ReqUserTaskCreate struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Assignee     string `json:"assignee"`
	Owner        string `json:"owner"`
	ParentTaskId string `json:"parentTaskId"`
}

func (t *UserTask) Create(req *ReqUserTaskCreate) error {
	_, err := t.client.doPostJson("/task/create", nil, req)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}
	return nil
}

func (t *UserTask) Get(id string) (*entitys.RespUserTask, error) {
	res, err := t.client.doGet("/task/"+id, map[string]string{})
	if err != nil {
		return nil, err
	}

	resp := &entitys.RespUserTask{}
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	return resp, nil
}

func (t *UserTask) GetList(query *ReqUserTaskGetList) ([]entitys.RespUserTask, error) {
	if query == nil {
		query = &ReqUserTaskGetList{}
	}
	queryParams := map[string]string{}

	res, err := t.client.doPostJson("/task", queryParams, query)
	if err != nil {
		return nil, err
	}

	var resp []entitys.RespUserTask
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}

	return resp, nil
}

type ReqUserTaskComplete struct {
	// A JSON object containing variable key-value pairs
	Variables map[string]entitys.Variable `json:"variables"`
}

func (t *UserTask) Complete(id string, query *ReqUserTaskComplete) error {
	_, err := t.client.doPostJson("/task/"+id+"/complete", map[string]string{}, query)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}
	return nil
}

func (t *UserTask) SubmitForm(id string, query *ReqUserTaskComplete) error {
	_, err := t.client.doPostJson("/task/"+id+"/submit-form", map[string]string{}, query)
	if err != nil {
		return fmt.Errorf("can't post json: %w", err)
	}
	return nil
}

func (t *UserTask) GetFormKey(id string) (*entitys.RespUserTaskForm, error) {
	res, err := t.client.doGet("/task/"+id+"/form", nil)
	if err != nil {
		return nil, err
	}
	resp := &entitys.RespUserTaskForm{}
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}
	return resp, nil
}

func (t *UserTask) SetAssignee(id, userName string) error {
	_, err := t.client.doPostJson("/task/"+id+"/assignee", nil, struct {
		UserId string `json:"userId"`
	}{UserId: userName})
	return err
}

func (t *UserTask) Claim(id, userName string) error {
	_, err := t.client.doPostJson("/task/"+id+"/claim", nil, struct {
		UserId string `json:"userId"`
	}{UserId: userName})
	return err
}

func (t *UserTask) UnClaim(id string) error {
	_, err := t.client.doPost("/task/"+id+"/unclaim", nil)
	return err
}

// CreateComment 不能设置用户，camunda默认是登录用户作为评论提交人 所以需要特别处理
func (t *UserTask) CreateComment(id, userName, message string) (*entitys.RespUserTaskComment, error) {

	m := entitys.CustomMessage{
		UserName: userName, Content: message,
	}
	mJson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	res, err := t.client.doPostJson("/task/"+id+"/comment/create", nil, struct {
		Message string `json:"message"`
	}{Message: string(mJson)})
	var resp *entitys.RespUserTaskComment
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}
	//从内容提取用户 放到结果中
	var mm entitys.CustomMessage
	if len(resp.Message) == 0 {
		return resp, nil
	}
	err = json.Unmarshal([]byte(resp.Message), &mm)
	if err != nil {
		return nil, err
	}
	resp.UserName = mm.UserName
	resp.Message = mm.Content
	return resp, err
}

func (t *UserTask) GetCommentList(id string) ([]entitys.RespUserTaskComment, error) {
	res, err := t.client.doGet("/task/"+id+"/comment", nil)
	if err != nil {
		return nil, err
	}

	var resp []entitys.RespUserTaskComment
	if err := t.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}
	for k, v := range resp {
		var mm entitys.CustomMessage
		err1 := json.Unmarshal([]byte(v.Message), &mm)
		if err1 != nil {
			log.Print(v.Message, err)
			continue
		}
		resp[k].Message = mm.Content
		resp[k].UserName = mm.UserName
	}
	return resp, err
}
