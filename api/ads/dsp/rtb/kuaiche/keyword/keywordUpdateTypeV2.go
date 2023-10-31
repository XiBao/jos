package keyword

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 修改关键词匹配类型
type KuaicheKeywordUpdateTypeV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordUpdateTypeV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt  `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordUpdateTypeV2Response struct {
	Responce  *KuaicheKeywordUpdateTypeV2Responce `json:"jingdong_ads_dsp_rtb_keyword_type_update_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_type_update_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                 `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordUpdateTypeV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordUpdateTypeV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordUpdateTypeV2Responce struct {
	Data *KuaicheKeywordUpdateTypeV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                  `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordUpdateTypeV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordUpdateTypeV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheKeywordUpdateTypeV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordUpdateTypeV2(req *KuaicheKeywordUpdateTypeV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordUpdateTypeV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheKeywordUpdateTypeV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
