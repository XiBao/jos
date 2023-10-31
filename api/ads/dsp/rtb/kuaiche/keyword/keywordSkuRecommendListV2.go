package keyword

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 快车商品推词
type KuaicheKeywordSkuRecommendListV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordSkuRecommendListV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt        `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordSkuRecommendListV2Response struct {
	Responce  *KuaicheKeywordSkuRecommendListV2Responce `json:"jingdong_ads_dsp_rtb_keyword_sku_recommend_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_sku_recommend_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                       `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordSkuRecommendListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordSkuRecommendListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordSkuRecommendListV2Responce struct {
	Data *KuaicheKeywordSkuRecommendListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                        `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordSkuRecommendListV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordSkuRecommendListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheKeywordSkuRecommendListV2ResponseData struct {
	Data []dsp.KeywordRecommend `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordSkuRecommendListV2(req *KuaicheKeywordSkuRecommendListV2Request) ([]dsp.KeywordRecommend, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordSkuRecommendListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheKeywordSkuRecommendListV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
