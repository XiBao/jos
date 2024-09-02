package creative

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	creative "github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/creative"
)

// 快车创意管理列表
type KuaicheCreativeListV2Request struct {
	api.BaseRequest
	Data   *creative.KuaicheCreativeListV2RequestData    `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCreativeListV2Response struct {
	Responce  *KuaicheCreativeListV2Responce `json:"jingdong_ads_dsp_rtb_kc_ad_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kc_ad_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCreativeListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCreativeListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCreativeListV2Responce struct {
	Data *KuaicheCreativeListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                             `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCreativeListV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCreativeListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCreativeListV2ResponseData struct {
	Data *KuaicheCreativeListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheCreativeListV2ResponseDataData struct {
	Creatives []dsp.CreativeData `json:"datas,omitempty" codec:"datas,omitempty"`
	Paginator *dsp.Paginator     `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheCreativeListV2(ctx context.Context, req *KuaicheCreativeListV2Request) (*KuaicheCreativeListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := creative.NewKuaicheCreativeListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCreativeListV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
