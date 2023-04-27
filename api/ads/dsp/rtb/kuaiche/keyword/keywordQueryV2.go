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

type KuaicheKeywordQueryV2Responce struct {
	Data *KuaicheKeywordQueryV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                             `json:"code,omitempty" codec:"code,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response KuaicheKeywordQueryV2Response
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
