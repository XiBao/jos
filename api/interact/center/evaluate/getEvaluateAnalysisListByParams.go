package center

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api/util"

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

type GetEvaluateAnalysisListByParamsData struct {
	Code      string              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []*EvaluateAnalysis `json:"result,omitempty" codec:"result,omitempty"`
}

func GetEvaluateAnalysisListByParams(req *GetEvaluateAnalysisListByParamsRequest) ([]*EvaluateAnalysis, error) {
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetEvaluateAnalysisListByParamsResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No result.")
	}

	return response.Data.Result, nil
}
