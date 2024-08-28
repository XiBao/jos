package campaign

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 查询京东快车计划列表信息和数据
type KuaicheCampaignListV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignListV2RequestData    `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignListV2Response struct {
	Responce  *KuaicheCampaignListV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCampaignListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCampaignListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCampaignListV2Responce struct {
	Data *KuaicheCampaignListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                             `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCampaignListV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCampaignListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCampaignListV2ResponseData struct {
	Data *KuaicheCampaignListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheCampaignListV2ResponseDataData struct {
	Campaigns []dsp.CampaignData `json:"datas,omitempty" codec:"datas,omitempty"`
	Paginator *dsp.Paginator     `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheCampaignListV2(ctx context.Context, req *KuaicheCampaignListV2Request) (*KuaicheCampaignListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCampaignListV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
