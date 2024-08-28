package creative

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	creative "github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/creative"
)

// 新增，修改与删除快车创意批量接口
type KuaicheCreativeAddV2Request struct {
	api.BaseRequest
	Data   *creative.KuaicheCreativeAddV2RequestData     `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCreativeAddV2Response struct {
	Responce  *KuaicheCreativeAddV2Responce `json:"jingdong_ads_dsp_rtb_kc_ad_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kc_ad_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheCreativeAddV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheCreativeAddV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheCreativeAddV2Responce struct {
	Data *KuaicheCreativeAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheCreativeAddV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheCreativeAddV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheCreativeAddV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCreativeAddV2(ctx context.Context, req *KuaicheCreativeAddV2Request) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := creative.NewKuaicheCreativeAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheCreativeAddV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Responce.Data.Data, nil
}
