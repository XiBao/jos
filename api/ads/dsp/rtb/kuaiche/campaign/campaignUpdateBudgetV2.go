package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 修改京东快车计划预算
type KuaicheCampaignUpdateBudgetV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignUpdateBudgetV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt      `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignUpdateBudgetV2Response struct {
	Responce  *KuaicheCampaignUpdateBudgetV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_updatebudget_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_updatebudget_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                    `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCampaignUpdateBudgetV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignUpdateBudgetV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateBudgetV2Responce struct {
	Data *KuaicheCampaignUpdateBudgetV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                     `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCampaignUpdateBudgetV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCampaignUpdateBudgetV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateBudgetV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignUpdateBudgetV2(ctx context.Context, req *KuaicheCampaignUpdateBudgetV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignUpdateBudgetV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCampaignUpdateBudgetV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Data, nil
}
