package vender

import (
	"github.com/XiBao/jos/sdk"
)

type GetCustomerBasicInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetCustomerBasicInfoRequest() (req *GetCustomerBasicInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.vender.getCustomerBasicInfo", Params: make(map[string]interface{}, 1)}
	req = &GetCustomerBasicInfoRequest{
		Request: &request,
	}
	return
}

func (req *GetCustomerBasicInfoRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *GetCustomerBasicInfoRequest) GetCustomerPin() string {
	unionId, found := req.Request.Params["customerPin"]
	if found {
		return unionId.(string)
	}
	return ""
}
