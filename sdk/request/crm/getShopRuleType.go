package crm

import "github.com/XiBao/jos/sdk"

type GetShopRuleTypeRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetShopRuleTypeRequest() (req *GetShopRuleTypeRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.getShopRuleType", Params: nil}
	req = &GetShopRuleTypeRequest{
		Request: &request,
	}
	return
}
