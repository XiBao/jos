package points

import "github.com/XiBao/jos/sdk"

type GetPointsRuleRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetPointsRuleRequest() (req *GetUserBaseInfoByPinRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.getPointsRule", Params: nil}
	req = &GetUserBaseInfoByPinRequest{
		Request: &request,
	}
	return
}
