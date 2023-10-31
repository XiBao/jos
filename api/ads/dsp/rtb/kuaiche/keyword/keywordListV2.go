package keyword

import (
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

func (r KuaicheKeywordListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheKeywordListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheKeywordListV2Responce struct {
	Data *KuaicheKeywordListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheKeywordListV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheKeywordListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
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

	var response KuaicheKeywordListV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
