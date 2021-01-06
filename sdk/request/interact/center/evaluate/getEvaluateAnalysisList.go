package center

import "github.com/XiBao/jos/sdk"

type GetEvaluateAnalysisListRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateAnalysisListRequest() (req *GetEvaluateAnalysisListRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateAnalysisReadService.getAnalysisList", Params: make(map[string]interface{})}
	req = &GetEvaluateAnalysisListRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateAnalysisListRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateAnalysisListRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateAnalysisListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateAnalysisListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateAnalysisListRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateAnalysisListRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}
