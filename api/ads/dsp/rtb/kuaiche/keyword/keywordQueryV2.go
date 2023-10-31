package keyword

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 查询单元下关键词投放信息
type KuaicheKeywordQueryV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordQueryV2RequestData     `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordQueryV2Response struct {
	Responce  *KuaicheKeywordQueryV2Responce `json:"jingdong_ads_dsp_rtb_keyword_query_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_query_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordQueryV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordQueryV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordQueryV2Responce struct {
	Data *KuaicheKeywordQueryV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                             `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordQueryV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordQueryV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheKeywordQueryV2ResponseData struct {
	Data []dsp.KeywordQueryData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordQueryV2(req *KuaicheKeywordQueryV2Request) ([]dsp.KeywordQueryData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordQueryV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheKeywordQueryV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
