package entitys

type RespProcessInstance struct {
	Id        string `json:"id"`
	Suspended string `json:"suspended"`
}

type RespActivityProcessInstance struct {
	ActivityProcessInstance
	ChildActivityInstances []ActivityProcessInstance `json:"childActivityInstances"`
}

type ActivityProcessInstance struct {
	ID           string `json:"id"`
	ActivityId   string `json:"activityId"`
	ActivityName string `json:"activityName"`
	ActivityType string `json:"activityType"`
}
