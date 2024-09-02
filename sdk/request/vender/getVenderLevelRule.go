package vender

import (
	"github.com/XiBao/jos/sdk"
)

type GetVenderLevelRuleRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetVenderLevelRuleRequest() (req *GetVenderLevelRuleRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.vender.getVenderLevelRule", Params: nil}
	req = &GetVenderLevelRuleRequest{
		Request: &request,
	}
	return
}
