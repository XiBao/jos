package keyword

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 删除关键词（支持批量）
type KuaicheKeywordDeleteV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordDeleteV2RequestData    `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordDeleteV2Response struct {
	Responce  *KuaicheKeywordDeleteV2Responce `json:"jingdong_ads_dsp_rtb_keyword_delete_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_delete_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse             `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordDeleteV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordDeleteV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordDeleteV2Responce struct {
	Data *KuaicheKeywordDeleteV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                              `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordDeleteV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordDeleteV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheKeywordDeleteV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordDeleteV2(ctx context.Context, req *KuaicheKeywordDeleteV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordDeleteV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheKeywordDeleteV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
