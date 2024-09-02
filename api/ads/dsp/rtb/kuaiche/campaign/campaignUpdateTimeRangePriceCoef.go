package campaign

import (
	"context"

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

func (r KuaicheCampaignUpdateTimeRangePriceCoefResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignUpdateTimeRangePriceCoefResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateTimeRangePriceCoefResponce struct {
	ReturnType *KuaicheCampaignUpdateTimeRangePriceCoefResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r KuaicheCampaignUpdateTimeRangePriceCoefResponce) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r KuaicheCampaignUpdateTimeRangePriceCoefResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateTimeRangePriceCoefResponseReturnType struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignUpdateTimeRangePriceCoef(ctx context.Context, req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) (uint64, error) {
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

	var response KuaicheCampaignUpdateTimeRangePriceCoefResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Responce.ReturnType.Data, nil
}
