package center

import "github.com/XiBao/jos/sdk"

type CancelEvaluateActivityRequest struct {
	Request *sdk.Request
}

// create new request
func NewCancelEvaluateActivityRequest() (req *CancelEvaluateActivityRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.write.EvaluateActivityWriteService.cancelActivity", Params: make(map[string]interface{})}
	req = &CancelEvaluateActivityRequest{
		Request: &request,
	}
	return
}

func (req *CancelEvaluateActivityRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *CancelEvaluateActivityRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *CancelEvaluateActivityRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *CancelEvaluateActivityRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *CancelEvaluateActivityRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *CancelEvaluateActivityRequest) GetId() uint64 {
	id, found := req.Request.Params["id"]
	if found {
		return id.(uint64)
	}
	return 0
}
