package center

import "github.com/XiBao/jos/sdk"

type GetEvaluateRuleListByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateRuleListByIdRequest() (req *GetEvaluateRuleListByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateRuleReadService.getRuleListByActivityId", Params: make(map[string]interface{})}
	req = &GetEvaluateRuleListByIdRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateRuleListByIdRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateRuleListByIdRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateRuleListByIdRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateRuleListByIdRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateRuleListByIdRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateRuleListByIdRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}
