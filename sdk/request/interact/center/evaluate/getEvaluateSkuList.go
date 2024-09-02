package center

import (
	"github.com/XiBao/jos/sdk"
)

type GetEvaluateSkuListRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateSkuListRequest() (req *GetEvaluateSkuListRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateSkuReadService.getSkuListByPatams", Params: make(map[string]interface{})}
	req = &GetEvaluateSkuListRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateSkuListRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateSkuListRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateSkuListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateSkuListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateSkuListRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateSkuListRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *GetEvaluateSkuListRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *GetEvaluateSkuListRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *GetEvaluateSkuListRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *GetEvaluateSkuListRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *GetEvaluateSkuListRequest) SetPageNumber(pageNumber uint) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *GetEvaluateSkuListRequest) GetPageNumber() uint {
	pageNumber, found := req.Request.Params["pageNubmer"]
	if found {
		return pageNumber.(uint)
	}
	return 0
}

func (req *GetEvaluateSkuListRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *GetEvaluateSkuListRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
