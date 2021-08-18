package entitys

type RespProcessDefinition struct {
	Id        string `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Version   int    `json:"Version"`
	Suspended bool   `json:"suspended"`
}

type RespActivityInstanceStatistics struct {
	Id         string                                   `json:"id"`
	Instances  int                                      `json:"instances"`
	FailedJobs int                                      `json:"failedJobs"`
	Incidents  []RespActivityInstanceStatisticsIncident `json:"incidents"`
}

type RespActivityInstanceStatisticsIncident struct {
	IncidentType  string `json:"incidentType"`
	IncidentCount int    `json:"incidentCount"`
}

type RespInstanceStatistics struct {
	Id         string                                   `json:"id"`
	Instances  int                                      `json:"instances"`
	FailedJobs int                                      `json:"failedJobs"`
	Definition RespProcessDefinition                    `json:"definition"`
	Incidents  []RespActivityInstanceStatisticsIncident `json:"incidents"`
}

type RespBPMNProcessDefinition struct {
	Id        string `json:"id"`
	Bpmn20Xml string `json:"bpmn20Xml"`
}

type RespBatch struct {
	Id                     string `json:"id"`
	Type                   string `json:"type"`
	TotalJobs              int    `json:"totalJobs"`
	BatchJobsPerSeed       int    `json:"batchJobsPerSeed"`
	InvocationsPerBatchJob int    `json:"invocationsPerBatchJob"`
	SeedJobDefinitionId    string `json:"seedJobDefinitionId"`
	MonitorJobDefinitionId string `json:"monitorJobDefinitionId"`
	BatchJobDefinitionId   string `json:"batchJobDefinitionId"`
	TenantId               string `json:"tenantId"`
}

type RespStartedProcessDefinition struct {
	Id           string              `json:"id"`
	DefinitionId string              `json:"definitionId"`
	BusinessKey  string              `json:"businessKey"`
	Suspended    bool                `json:"suspended"`
	Variables    map[string]Variable `json:"variables"`
}
