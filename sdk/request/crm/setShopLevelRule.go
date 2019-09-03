package crm

import "github.com/XiBao/jos/sdk"

type SetShopLevelRuleRequest struct {
	Request *sdk.Request
}

func NewSetShopLevelRuleRequest() (req *SetShopLevelRuleRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.setShopLevelRule", Params: make(map[string]interface{}, 1)}
	req = &SetShopLevelRuleRequest{
		Request: &request,
	}
	return
}

func (req *SetShopLevelRuleRequest) SetCustomerLevelName(CustomerLevelName []string) {
	req.Request.Params["customerLevelName"] = CustomerLevelName
}

func (req *SetShopLevelRuleRequest) GetCustomerLevelName() []string {
	customerLevelName, found := req.Request.Params["customerLevelName"]
	if found {
		return customerLevelName.([]string)
	}
	return nil
}
