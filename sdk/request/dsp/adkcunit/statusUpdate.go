package adkcunit

import "github.com/XiBao/jos/sdk"

type AdkcunitStatusUpdateRequest struct {
	Request *sdk.Request
}

func NewAdkcunitStatusUpdateRequest() (req *AdkcunitStatusUpdateRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkcunit.status.update", Params: make(map[string]interface{}, 2)}
	req = &AdkcunitStatusUpdateRequest{
		Request: &request,
	}
	return
}

func (req *AdkcunitStatusUpdateRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *AdkcunitStatusUpdateRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}

func (req *AdkcunitStatusUpdateRequest) SetAdGroupId(adGroupId string) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *AdkcunitStatusUpdateRequest) GetAdGroupId() string {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(string)
	}
	return ""
}
