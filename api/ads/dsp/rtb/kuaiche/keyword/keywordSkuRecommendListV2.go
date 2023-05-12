package keyword

import (
	"encoding/json"
	"errors"

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

type KuaicheKeywordSkuRecommendListV2Responce struct {
	Data *KuaicheKeywordSkuRecommendListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                        `json:"code,omitempty" codec:"code,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response KuaicheKeywordSkuRecommendListV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return nil, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return nil, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Data, nil
}