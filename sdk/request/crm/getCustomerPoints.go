package crm

import (
	"github.com/XiBao/jos/sdk"
)

type GetCustomerPointsRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetCustomerPointsRequest() (req *GetCustomerPointsRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.getCustomerPoints", Params: make(map[string]interface{}, 1)}
	req = &GetCustomerPointsRequest{
		Request: &request,
	}
	return
}

func (req *GetCustomerPointsRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *GetCustomerPointsRequest) GetCustomerPin() string {
	pin, found := req.Request.Params["customerPin"]
	if found {
		return pin.(string)
	}
	return ""
}
