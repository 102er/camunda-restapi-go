package services

type IVariableInstance interface {
	GetListByPost()
}

type VariableInstance struct {
	client *BaseRestApiService
}

func NewVariableInstance(client *BaseRestApiService) *VariableInstance {
	return &VariableInstance{
		client: client,
	}
}

type ReqVariableInstance struct {
}

func (v *VariableInstance) GetListByPost() {

}
