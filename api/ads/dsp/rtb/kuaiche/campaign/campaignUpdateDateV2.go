package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 修改京东快车计划投放时间
type KuaicheCampaignUpdateDateV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignUpdateDateV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt    `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignUpdateDateV2Response struct {
	Responce  *KuaicheCampaignUpdateDateV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_UpdateDate_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_UpdateDate_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCampaignUpdateDateV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignUpdateDateV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateDateV2Responce struct {
	Data *KuaicheCampaignUpdateDateV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                   `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCampaignUpdateDateV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCampaignUpdateDateV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCampaignUpdateDateV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignUpdateDateV2(ctx context.Context, req *KuaicheCampaignUpdateDateV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignUpdateDateV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCampaignUpdateDateV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
