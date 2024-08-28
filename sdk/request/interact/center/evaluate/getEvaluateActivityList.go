package center

import (
	"github.com/XiBao/jos/sdk"
)

type GetEvaluateActivityListRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateActivityListRequest() (req *GetEvaluateActivityListRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateActivityReadService.getActivityListByParams", Params: make(map[string]interface{})}
	req = &GetEvaluateActivityListRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateActivityListRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateActivityListRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateActivityListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateActivityListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateActivityListRequest) SetStatus(status string) {
	req.Request.Params["status"] = status
}

func (req *GetEvaluateActivityListRequest) GetStatus() string {
	status, found := req.Request.Params["status"]
	if found {
		return status.(string)
	}
	return ""
}

func (req *GetEvaluateActivityListRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *GetEvaluateActivityListRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *GetEvaluateActivityListRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *GetEvaluateActivityListRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *GetEvaluateActivityListRequest) SetPageNumber(pageNumber uint) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *GetEvaluateActivityListRequest) GetPageNumber() uint {
	pageNumber, found := req.Request.Params["pageNubmer"]
	if found {
		return pageNumber.(uint)
	}
	return 0
}

func (req *GetEvaluateActivityListRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *GetEvaluateActivityListRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
