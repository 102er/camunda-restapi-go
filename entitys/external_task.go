package entitys

type RespExternalTask struct {
	ActivityId           string `json:"activityId"`
	ActivityInstanceId   string `json:"activityInstanceId"`
	ErrorMessage         string `json:"errorMessage"`
	ErrorDetails         string `json:"errorDetails"`
	ExecutionId          string `json:"executionId"`
	Id                   string `json:"id"`
	LockExpirationTime   string `json:"lockExpirationTime"`
	ProcessDefinitionId  string `json:"processDefinitionId"`
	ProcessDefinitionKey string `json:"processDefinitionKey"`
	ProcessInstanceId    string `json:"processInstanceId"`
	Retries              int    `json:"retries"`
	Suspended            bool   `json:"suspended"`
	WorkerId             string `json:"workerId"`
	Priority             int    `json:"priority"`
	TopicName            string `json:"topicName"`
	BusinessKey          string `json:"businessKey"`
}

// RespLockedExternalTask a response FetchAndLock method
type RespLockedExternalTask struct {
	ActivityId           string              `json:"activityId"`
	ActivityInstanceId   string              `json:"activityInstanceId"`
	ErrorMessage         string              `json:"errorMessage"`
	ErrorDetails         string              `json:"errorDetails"`
	ExecutionId          string              `json:"executionId"`
	Id                   string              `json:"id"`
	LockExpirationTime   string              `json:"lockExpirationTime"`
	ProcessDefinitionId  string              `json:"processDefinitionId"`
	ProcessDefinitionKey string              `json:"processDefinitionKey"`
	ProcessInstanceId    string              `json:"processInstanceId"`
	Retries              int                 `json:"retries"`
	WorkerId             string              `json:"workerId"`
	Priority             int                 `json:"priority"`
	TopicName            string              `json:"topicName"`
	BusinessKey          string              `json:"businessKey"`
	Variables            map[string]Variable `json:"variables"`
}
