package tests

import "testing"

type (

	CosmosLogs []CosmosLog

	CosmosLog struct {
		MgsIndex int `json:"mgs_index"`
		Events Events `json:"events"`
	}
	Event struct {
		Type string `json:"type"`
		Attributes []Attribute `json:"attributes"`
	}

	Attribute struct {
		Key string `json:"key"`
		Value string `json:"value"`
	}
	Events []Event
	Attributes []Attribute
)

func (e Events)GetEventByType(t *testing.T,eType string)Attributes  {
	for _,ev := range e{
		if ev.Type == eType{
			return ev.Attributes
		}
	}
	t.Log("event type not found")
	t.Fail()
	return nil
}


func (as Attributes)Get(t *testing.T,key string)Attribute  {
	for _,a := range as{
		if a.Key == key{
			return a
		}
	}
	t.Log("attribute key not found")
	t.Fail()
	return Attribute{}
}