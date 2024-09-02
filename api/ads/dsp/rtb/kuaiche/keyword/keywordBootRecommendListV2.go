package keyword

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/keyword"
)

// 快车商品核心拓词
type KuaicheKeywordBootRecommendListV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordBootRecommendListV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt         `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordBootRecommendListV2Response struct {
	Responce  *KuaicheKeywordBootRecommendListV2Responce `json:"jingdong_ads_dsp_rtb_keyword_boot_recommend_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_boot_recommend_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                        `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheKeywordBootRecommendListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordBootRecommendListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordBootRecommendListV2Responce struct {
	Data *KuaicheKeywordBootRecommendListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                         `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordBootRecommendListV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordBootRecommendListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheKeywordBootRecommendListV2ResponseData struct {
	Data []dsp.KeywordRecommend `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordBootRecommendListV2(ctx context.Context, req *KuaicheKeywordBootRecommendListV2Request) ([]dsp.KeywordRecommend, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordBootRecommendListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheKeywordBootRecommendListV2Response
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
