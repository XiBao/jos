package crm

import (
	"github.com/XiBao/jos/sdk"
)

type SetMemberGradeRequest struct {
	Request *sdk.Request
}

func NewSetMemberGradeRequest() (req *SetMemberGradeRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.setMemberGrade", Params: make(map[string]interface{}, 2)}
	req = &SetMemberGradeRequest{
		Request: &request,
	}
	return
}

func (req *SetMemberGradeRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SetMemberGradeRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SetMemberGradeRequest) SetGrade(grade uint8) {
	req.Request.Params["grade"] = grade
}

func (req *SetMemberGradeRequest) GetGrade() uint8 {
	grade, found := req.Request.Params["grade"]
	if found {
		return grade.(uint8)
	}
	return 0
}
