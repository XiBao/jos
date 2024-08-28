package center

import (
	"github.com/XiBao/jos/sdk"
)

type GetEvaluateJoinedSkuListRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateJoinedSkuListRequest() (req *GetEvaluateJoinedSkuListRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateSkuReadService.getJoinedSkuListByParams", Params: make(map[string]interface{})}
	req = &GetEvaluateJoinedSkuListRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateJoinedSkuListRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateJoinedSkuListRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateJoinedSkuListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateJoinedSkuListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateJoinedSkuListRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateJoinedSkuListRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *GetEvaluateJoinedSkuListRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *GetEvaluateJoinedSkuListRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *GetEvaluateJoinedSkuListRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *GetEvaluateJoinedSkuListRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *GetEvaluateJoinedSkuListRequest) SetPageNumber(pageNumber uint) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *GetEvaluateJoinedSkuListRequest) GetPageNumber() uint {
	pageNumber, found := req.Request.Params["pageNubmer"]
	if found {
		return pageNumber.(uint)
	}
	return 0
}

func (req *GetEvaluateJoinedSkuListRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *GetEvaluateJoinedSkuListRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
