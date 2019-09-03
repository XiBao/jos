package crm

import "github.com/XiBao/jos/sdk"

type IsPointsEnabledRequest struct {
	Request *sdk.Request
}

func NewIsPointsEnabledRequest() (req *IsPointsEnabledRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.isPointsEnabled", Params: nil}
	req = &IsPointsEnabledRequest{
		Request: &request,
	}
	return
}
