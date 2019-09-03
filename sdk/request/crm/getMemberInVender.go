package crm

import (
	"github.com/XiBao/jos/sdk"
)

type GetMemberInVenderRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetMemberInVenderRequest() (req *GetMemberInVenderRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.getMemberInVender", Params: make(map[string]interface{}, 1)}
	req = &GetMemberInVenderRequest{
		Request: &request,
	}
	return
}

func (req *GetMemberInVenderRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *GetMemberInVenderRequest) GetCustomerPin() string {
	pin, found := req.Request.Params["customerPin"]
	if found {
		return pin.(string)
	}
	return ""
}
