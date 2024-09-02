package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 新建京东快车计划
type KuaicheCampaignAddV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignAddV2RequestData     `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignAddV2Response struct {
	Responce  *KuaicheCampaignAddV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCampaignAddV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignAddV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignAddV2Responce struct {
	Data *KuaicheCampaignAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCampaignAddV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCampaignAddV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCampaignAddV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignAddV2(ctx context.Context, req *KuaicheCampaignAddV2Request) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCampaignAddV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Responce.Data.Data, nil
}
