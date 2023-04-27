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

// 快车普通计划关键词管理列表
type KuaicheKeywordListV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordListV2RequestData      `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordListV2Response struct {
	Responce  *KuaicheKeywordListV2Responce `json:"jingdong_ads_dsp_rtb_keyword_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheKeywordListV2Responce struct {
	Data *KuaicheKeywordListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheKeywordListV2ResponseData struct {
	Data *KuaicheKeywordListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheKeywordListV2ResponseDataData struct {
	Keywords  []dsp.KeywordData   `json:"datas,omitempty" codec:"datas,omitempty"`
	Ext       *dsp.KeywordExtData `json:"ext,omitempty" codec:"ext,omitempty"`
	Paginator *dsp.Paginator      `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheKeywordListV2(req *KuaicheKeywordListV2Request) (*KuaicheKeywordListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordListV2Request()
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

	var response KuaicheKeywordListV2Response
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
