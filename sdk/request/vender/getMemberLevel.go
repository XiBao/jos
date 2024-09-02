package vender

import (
	"github.com/XiBao/jos/sdk"
)

type GetMemberLevelRequest struct {
	Request *sdk.Request
}

// create new request
func NewVenderGetMemberLevelRequest() (req *GetMemberLevelRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.vender.getMemberLevel", Params: make(map[string]interface{}, 1)}
	req = &GetMemberLevelRequest{
		Request: &request,
	}
	return
}

func (req *GetMemberLevelRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *GetMemberLevelRequest) GetCustomerPin() string {
	unionId, found := req.Request.Params["customerPin"]
	if found {
		return unionId.(string)
	}
	return ""
}
