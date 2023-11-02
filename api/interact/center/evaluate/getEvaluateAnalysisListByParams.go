package center

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateAnalysisListByParamsRequest struct {
	api.BaseRequest
	AppName    string `json:"appName" codec:"appName"`
	Channel    uint8  `json:"channel" codec:"channel"`
	ActivityId uint64 `json:"activityId" codec:"activityId"`
	PageNumber uint   `json:"pageNumber" codec:"pageNumber"`
	PageSize   uint   `json:"pageSize" codec:"pageSize"`
	SkuId      uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`
}

type GetEvaluateAnalysisListByParamsResponse struct {
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateAnalysisListByParamsData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateAnalysisReadService_getAnalysisListByParams_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateAnalysisReadService_getAnalysisListByParams_responce,omitempty"`
}

func (r GetEvaluateAnalysisListByParamsResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEvaluateAnalysisListByParamsResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEvaluateAnalysisListByParamsData struct {
	Code      string             `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string             `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []EvaluateAnalysis `json:"result,omitempty" codec:"result,omitempty"`
}

func (r GetEvaluateAnalysisListByParamsData) IsError() bool {
	return r.Code != "0"
}

func (r GetEvaluateAnalysisListByParamsData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func GetEvaluateAnalysisListByParams(req *GetEvaluateAnalysisListByParamsRequest) ([]EvaluateAnalysis, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateAnalysisListByParamsRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)
	r.SetPageSize(req.PageSize)
	r.SetPageNumber(req.PageNumber)
	if req.SkuId > 0 {
		r.SetSkuId(req.SkuId)
	}

	var response GetEvaluateAnalysisListByParamsResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
