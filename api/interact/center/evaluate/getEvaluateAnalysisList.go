package center

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api/util"

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

type GetEvaluateAnalysisListData struct {
	Code      string              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []*EvaluateAnalysis `json:"result,omitempty" codec:"result,omitempty"`
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

func GetEvaluateAnalysisList(req *GetEvaluateAnalysisListRequest) ([]*EvaluateAnalysis, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateAnalysisListRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetEvaluateAnalysisListResponse
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
