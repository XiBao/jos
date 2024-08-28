package center

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateAnalysisListRequest struct {
	api.BaseRequest
	AppName    string `json:"appName" codec:"appName"`
	Channel    uint8  `json:"channel" codec:"channel"`
	ActivityId uint64 `json:"activityId" codec:"activityId"`
}

type GetEvaluateAnalysisListResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateAnalysisListData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateAnalysisReadService_getAnalysisList_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateAnalysisReadService_getAnalysisList_responce,omitempty"`
}

func (r GetEvaluateAnalysisListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEvaluateAnalysisListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEvaluateAnalysisListData struct {
	Code      string             `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string             `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []EvaluateAnalysis `json:"result,omitempty" codec:"result,omitempty"`
}

func (r GetEvaluateAnalysisListData) IsError() bool {
	return r.Code != "0"
}

func (r GetEvaluateAnalysisListData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type EvaluateAnalysis struct {
	ActivityId    uint64  `json:"activity_id"`    // 活动id
	PrizeRate     float64 `json:"prize_rate"`     // 获奖比率
	Count         uint    `json:"count"`          // 活动期间的评价数量
	StandardRate  float64 `json:"standard_rate"`  // 达标评价的数量
	VenderId      uint64  `json:"vender_id"`      // 商家id
	SkuId         uint64  `json:"sku_id"`         // skuId
	Id            uint64  `json:"id"`             // 业务id
	StandardCount uint    `json:"standard_count"` // 达标评价的数量
	PrizeCount    uint    `json:"prize_count"`    // 获奖评价的数量
}

func GetEvaluateAnalysisList(ctx context.Context, req *GetEvaluateAnalysisListRequest) ([]EvaluateAnalysis, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateAnalysisListRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)

	var response GetEvaluateAnalysisListResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
