package services

import (
	"fmt"
	"github.com/blossom102er/camunda-restapi-go/entitys"
)

type IExternalTask interface {
	Get(id string) (*entitys.RespExternalTask, error)
	GetList(query map[string]string) ([]*entitys.RespExternalTask, error)
	FetchAndLock(query QueryFetchAndLock) ([]*entitys.RespLockedExternalTask, error)
	Complete(id string, query QueryComplete) error
	HandleFailure(id string, query QueryHandleFailure) error
	Unlock(id string) error
	ExtendLock(id string, query QueryExtendLock) error
	SetPriority(id string, priority int) error
	SetRetries(id string, retries int) error
}

type ExternalTask struct {
	client *BaseRestApiService
}

func NewExternalTask(client *BaseRestApiService) *ExternalTask {
	return &ExternalTask{
		client: client,
	}
}

func (e *ExternalTask) Get(id string) (*entitys.RespExternalTask, error) {
	resp := &entitys.RespExternalTask{}
	res, err := e.client.doGet(
		"/external-task/"+id,
		map[string]string{},
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExternalTask) GetList(query map[string]string) ([]*entitys.RespExternalTask, error) {
	var resp []*entitys.RespExternalTask
	res, err := e.client.doGet(
		"/external-task",
		query,
	)
	if err != nil {
		return nil, err
	}

	if err := e.client.readJsonResponse(res, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type QueryFetchAndLock struct {
	WorkerId             string                    `json:"workerId"`
	MaxTasks             int                       `json:"maxTasks"`
	UsePriority          *bool                     `json:"usePriority,omitempty"`
	AsyncResponseTimeout *int                      `json:"asyncResponseTimeout,omitempty"`
	Topics               *[]QueryFetchAndLockTopic `json:"topics,omitempty"`
}
type QueryFetchAndLockTopic struct {
	TopicName              string                      `json:"topicName"`
	LockDuration           int                         `json:"lockDuration"`
	Variables              *[]string                   `json:"variables,omitempty"`
	LocalVariables         *bool                       `json:"localVariables,omitempty"`
	BusinessKey            *string                     `json:"businessKey,omitempty"`
	ProcessDefinitionId    *string                     `json:"processDefinitionId,omitempty"`
	ProcessDefinitionIdIn  *string                     `json:"processDefinitionIdIn,omitempty"`
	ProcessDefinitionKey   *string                     `json:"processDefinitionKey,omitempty"`
	ProcessDefinitionKeyIn *string                     `json:"processDefinitionKeyIn,omitempty"`
	ProcessVariables       map[string]entitys.Variable `json:"processVariables,omitempty"`
	DeserializeValues      *bool                       `json:"deserializeValues,omitempty"`
}

func (e *ExternalTask) FetchAndLock(query QueryFetchAndLock) ([]*entitys.RespLockedExternalTask, error) {
	var resp []*entitys.RespLockedExternalTask
	res, err := e.client.doPostJson(
		"/external-task/fetchAndLock",
		map[string]string{},
		&query,
	)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	if err := e.client.readJsonResponse(res, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type QueryComplete struct {
	WorkerId       *string                      `json:"workerId,omitempty"`
	Variables      *map[string]entitys.Variable `json:"variables"`
	LocalVariables *map[string]entitys.Variable `json:"localVariables"`
}

func (e *ExternalTask) Complete(id string, query QueryComplete) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/complete", map[string]string{}, &query)
	return err
}

// QueryHandleFailure a query for HandleFailure request
type QueryHandleFailure struct {
	WorkerId     *string `json:"workerId,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorDetails *string `json:"errorDetails,omitempty"`
	Retries      *int    `json:"retries,omitempty"`
	RetryTimeout *int    `json:"retryTimeout,omitempty"`
}

func (e *ExternalTask) HandleFailure(id string, query QueryHandleFailure) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/failure", map[string]string{}, &query)
	return err
}

// Unlock a unlocks an external task by id. Clears the task’s lock expiration time and worker id
func (e *ExternalTask) Unlock(id string) error {
	_, err := e.client.doPost("/external-task/"+id+"/unlock", map[string]string{})
	return err
}

// QueryExtendLock a query for ExtendLock request
type QueryExtendLock struct {
	NewDuration *int    `json:"newDuration,omitempty"`
	WorkerId    *string `json:"workerId,omitempty"`
}

// ExtendLock a extends the timeout of the lock by a given amount of time
func (e *ExternalTask) ExtendLock(id string, query QueryExtendLock) error {
	_, err := e.client.doPostJson("/external-task/"+id+"/extendLock", map[string]string{}, &query)
	return err
}
func (e *ExternalTask) SetPriority(id string, priority int) error {
	_, err := e.client.doPut("/external-task/"+id+"/priority", map[string]string{})
	return err
}
func (e *ExternalTask) SetRetries(id string, retries int) error {
	return e.client.doPutJson("/external-task/"+id+"/retries", map[string]string{}, map[string]int{
		"retries": retries,
	})
}
