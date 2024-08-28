package center

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
)

type GetEvaluateRuleListByIdRequest struct {
	api.BaseRequest
	AppName    string `json:"appName" codec:"appName"`
	Channel    uint8  `json:"channel" codec:"channel"`
	ActivityId uint64 `json:"activityId" codec:"activityId"`
}

type GetEvaluateRuleListByIdResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEvaluateRuleListByIdData `json:"jingdong_com_jd_interact_center_api_service_read_EvaluateRuleReadService_getRuleListByActivityId_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_service_read_EvaluateRuleReadService_getRuleListByActivityId_responce,omitempty"`
}

func (r GetEvaluateRuleListByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEvaluateRuleListByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEvaluateRuleListByIdData struct {
	Code      string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []EvaluateRule `json:"result,omitempty" codec:"result,omitempty"`
}

func (r GetEvaluateRuleListByIdData) IsError() bool {
	return r.Code != "0"
}

func (r GetEvaluateRuleListByIdData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type EvaluateRule struct {
	VenderId     uint64  `json:"vender_id"`
	AssetItemId  uint    `json:"asset_item_id"`
	Type         uint8   `json:"type"`
	CouponId     uint64  `json:"coupon_id"`
	Discount     float64 `json:"discount"`
	Id           uint64  `json:"id"`
	ActivityId   uint64  `json:"activity_id"`
	AwardType    uint8   `json:"award_type"`
	Quota        uint    `json:"quota"`
	Price        float64 `json:"price"`
	FloatRatio   float64 `json:"float_ratio"`
	Nums         uint    `json:"nums"`
	CollectTimes uint    `json:"collect_times"`
	ExpireType   uint8   `json:"expire_type"`
}

func GetEvaluateRuleListById(ctx context.Context, req *GetEvaluateRuleListByIdRequest) ([]EvaluateRule, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateRuleListByIdRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)

	var response GetEvaluateRuleListByIdResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.Result, nil
}
