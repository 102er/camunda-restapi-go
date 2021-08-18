package services

import (
	"github.com/blossom102er/camunda-restapi-go/entitys"
)

type IMessage interface {
	SendMessage(query *ReqMessage) (resp []entitys.RespMessage, err error)
}

// Message a client for Message API
type Message struct {
	client *BaseRestApiService
}

// ReqMessage a request to send a message
type ReqMessage struct {
	MessageName      string                       `json:"messageName"`
	BusinessKey      string                       `json:"businessKey"`
	ProcessVariables *map[string]entitys.Variable `json:"processVariables,omitempty"`
	ResultEnabled    bool                         `json:"resultEnabled"`
}

func NewMessage(client *BaseRestApiService) *Message {
	return &Message{
		client: client,
	}
}

// SendMessage sends message to a process
func (m *Message) SendMessage(query *ReqMessage) (resp []entitys.RespMessage, err error) {
	res, err := m.client.doPostJson("/message", map[string]string{}, query)
	//不接受返回值
	if !query.ResultEnabled {
		return
	}
	if err != nil {
		return
	}
	err = m.client.readJsonResponse(res, &resp)
	return
}
