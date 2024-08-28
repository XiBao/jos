package center

import (
	"github.com/XiBao/jos/sdk"
)

type GetEvaluateActivityByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateActivityByIdRequest() (req *GetEvaluateActivityByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.vender.read.evaluate.getActivityById", Params: make(map[string]interface{})}
	req = &GetEvaluateActivityByIdRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateActivityByIdRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateActivityByIdRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateActivityByIdRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateActivityByIdRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateActivityByIdRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateActivityByIdRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}
