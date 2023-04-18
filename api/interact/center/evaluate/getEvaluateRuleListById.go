package center

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api/util"

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

type GetEvaluateRuleListByIdData struct {
	Code      string          `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string          `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    []*EvaluateRule `json:"result,omitempty" codec:"result,omitempty"`
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

func GetEvaluateRuleListById(req *GetEvaluateRuleListByIdRequest) ([]*EvaluateRule, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewGetEvaluateRuleListByIdRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetActivityId(req.ActivityId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response GetEvaluateRuleListByIdResponse
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
