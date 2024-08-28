package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 查询京东快车计划基本信息
type KuaicheCampaignGetV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignGetV2RequestData     `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignGetV2Response struct {
	Responce  *KuaicheCampaignGetV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_get_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_get_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCampaignGetV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignGetV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignGetV2Responce struct {
	Data *KuaicheCampaignGetV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCampaignGetV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCampaignGetV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCampaignGetV2ResponseData struct {
	Data *dsp.Campaign `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignGetV2(ctx context.Context, req *KuaicheCampaignGetV2Request) (*dsp.Campaign, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignGetV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCampaignGetV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
