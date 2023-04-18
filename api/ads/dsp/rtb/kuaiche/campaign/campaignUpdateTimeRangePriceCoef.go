package campaign

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 快车修改折扣时段
type KuaicheCampaignUpdateTimeRangePriceCoefRequest struct {
	api.BaseRequest
	Id                   uint64 `json:"id" codec:"id"`                                                         // 计划id
	TimeRangePriceCoef   string `json:"timeRangePriceCoef" codec:"timeRangePriceCoef"`                         // 分时段溢价，溢价系数30-500，per:投放比率（per:投放比率，计算方式：所选择溢价非0的时间/总数（7*24））,detail中0-6表示7天，数组中为24小时的折扣系数，100表示无折扣
	AccessPin            string `json:"accessPin,omitempty" codec:"accessPin,omitempty"`                       // 被免密访问的pin
	AuthType             string `json:"authType,omitempty" codec:"authType,omitempty"`                         // 授权模式
	PlatformBusinessType string `json:"platformBusinessType,omitempty" codec:"platformBusinessType,omitempty"` // 平台业务类型，DST_JZT：京准通
}

type KuaicheCampaignUpdateTimeRangePriceCoefResponse struct {
	Responce  *KuaicheCampaignUpdateTimeRangePriceCoefResponce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_updateTimeRangePriceCoef_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_updateTimeRangePriceCoef_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                              `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheCampaignUpdateTimeRangePriceCoefResponce struct {
	ReturnType *KuaicheCampaignUpdateTimeRangePriceCoefResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type KuaicheCampaignUpdateTimeRangePriceCoefResponseReturnType struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignUpdateTimeRangePriceCoef(req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignUpdateTimeRangePriceCoefRequest()
	r.SetId(req.Id)
	r.SetTimeRangePriceCoef(req.TimeRangePriceCoef)
	if req.AccessPin != "" {
		r.SetAccessPin(req.AccessPin)
	}
	if req.AuthType != "" {
		r.SetAuthType(req.AuthType)
	}
	if req.PlatformBusinessType != "" {
		r.SetPlatformBusinessType(req.PlatformBusinessType)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response KuaicheCampaignUpdateTimeRangePriceCoefResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.ReturnType == nil {
		return 0, errors.New("no result data.")
	}
	if !response.Responce.ReturnType.Success {
		return 0, errors.New(response.Responce.ReturnType.Msg)
	}

	return response.Responce.ReturnType.Data, nil
}
