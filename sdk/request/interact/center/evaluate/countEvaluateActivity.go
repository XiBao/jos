package center

import "github.com/XiBao/jos/sdk"

type CountEvaluateActivityRequest struct {
	Request *sdk.Request
}

// create new request
func NewCountEvaluateActivityRequest() (req *CountEvaluateActivityRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateActivityReadService.countByParams", Params: make(map[string]interface{})}
	req = &CountEvaluateActivityRequest{
		Request: &request,
	}
	return
}

func (req *CountEvaluateActivityRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *CountEvaluateActivityRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *CountEvaluateActivityRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *CountEvaluateActivityRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *CountEvaluateActivityRequest) SetStatus(status string) {
	req.Request.Params["status"] = status
}

func (req *CountEvaluateActivityRequest) GetStatus() string {
	status, found := req.Request.Params["status"]
	if found {
		return status.(string)
	}
	return ""
}

func (req *CountEvaluateActivityRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *CountEvaluateActivityRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *CountEvaluateActivityRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *CountEvaluateActivityRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *CountEvaluateActivityRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *CountEvaluateActivityRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *CountEvaluateActivityRequest) SetPageNumber(pageNumber uint) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *CountEvaluateActivityRequest) GetPageNumber() uint {
	pageNumber, found := req.Request.Params["pageNubmer"]
	if found {
		return pageNumber.(uint)
	}
	return 0
}

func (req *CountEvaluateActivityRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *CountEvaluateActivityRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
