package entitys

// RespDeploymentCreate a JSON object corresponding to the DeploymentWithDefinitions interface in the engine
type RespDeploymentCreate struct {
	Id                         string                           `json:"id"`
	Name                       string                           `json:"name"`
	Source                     string                           `json:"source"`
	DeploymentTime             Time                             `json:"deploymentTime"`
	DeployedProcessDefinitions map[string]RespProcessDefinition `json:"deployedProcessDefinitions"`
}
