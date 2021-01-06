package center

import "github.com/XiBao/jos/sdk"

type GetEvaluateAnalysisListByParamsRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEvaluateAnalysisListByParamsRequest() (req *GetEvaluateAnalysisListByParamsRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.read.EvaluateAnalysisReadService.getAnalysisListByParams", Params: make(map[string]interface{})}
	req = &GetEvaluateAnalysisListByParamsRequest{
		Request: &request,
	}
	return
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetPageNumber(pageNumber uint) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetPageNumber() uint {
	pageNumber, found := req.Request.Params["pageNumber"]
	if found {
		return pageNumber.(uint)
	}
	return 0
}

func (req *GetEvaluateAnalysisListByParamsRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *GetEvaluateAnalysisListByParamsRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}
