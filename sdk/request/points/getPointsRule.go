package points

import "github.com/XiBao/jos/sdk"

type GetPointsRuleRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetPointsRuleRequest() (req *GetPointsRuleRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.getPointsRule", Params: nil}
	req = &GetPointsRuleRequest{
		Request: &request,
	}
	return
}
