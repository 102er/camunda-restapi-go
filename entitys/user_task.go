package entitys

type RespUserTask struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Assignee            string `json:"assignee"`
	Created             string `json:"created"`
	Due                 string `json:"due"`
	ParentTaskId        string `json:"parentTaskId"`
	Owner               string `json:"owner"`
	Priority            int64  `json:"priority"`
	ProcessDefinitionId string `json:"processDefinitionId"`
	ProcessInstanceId   string `json:"processInstanceId"`
	TaskDefinitionKey   string `json:"taskDefinitionKey"`
	Suspended           bool   `json:"suspended"`
	FormKey             string `json:"formKey"`
}
type RespUserTaskForm struct {
	ContextPath string `json:"contextPath"`
	Key         string `json:"key"`
}

type RespUserTaskComment struct {
	UserName   string `json:"userId"`
	Message    string `json:"message"`
	CreateTime string `json:"time"`
}

type CustomMessage struct {
	UserName string `json:"userName"`
	Content  string `json:"content"`
}
