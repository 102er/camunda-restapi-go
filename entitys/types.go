package entitys

import (
	"strings"
	"time"
)

const DefaultDateTimeFormat = "2006-01-02T15:04:05.000-0700"

type VariableSet struct {
	// The variable's value
	Value string `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
	// A JSON object containing additional, value-type-dependent properties
	ValueInfo ValueInfo `json:"valueInfo"`
	// Indicates whether the variable should be a local variable or not. If set to true, the variable becomes a local
	// variable of the execution entering the target activity
	Local bool `json:"local"`
}

type RespCount struct {
	Count int `json:"count"`
}

type RespLink struct {
	Method string `json:"method"`
	Href   string `json:"href"`
	Rel    string `json:"rel"`
}

// Variable a variable
type Variable struct {
	// The variable's value
	Value interface{} `json:"value"`
	// The value type of the variable.
	Type string `json:"type"`
	// A JSON object containing additional, value-type-dependent properties
	ValueInfo ValueInfo `json:"valueInfo,omitempty"`
}

// ValueInfo a value info in variable
type ValueInfo struct {
	// A string representation of the object's type name
	ObjectTypeName *string `json:"objectTypeName,omitempty"`
	// The serialization format used to store the variable.
	SerializationDataFormat *string `json:"serializationDataFormat,omitempty"`
}

type RespMessage struct {
	ResultType      string                 `json:"resultType"`
	Execution       string                 `json:"execution"`
	ProcessInstance MessageProcessInstance `json:"processInstance"`
}

type MessageProcessInstance struct {
	Id           string `json:"id"`
	DefinitionId string `json:"definitionId"`
	BusinessKey  string `json:"businessKey"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	t.Time, err = time.Parse(DefaultDateTimeFormat, strings.Trim(string(b), "\""))
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	timeStr := t.Time.Format(DefaultDateTimeFormat)
	return []byte("\"" + timeStr + "\""), nil
}
