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

// 删除关键词（支持批量）
type KuaicheKeywordAddV2Request struct {
	api.BaseRequest
	Data   *keyword.KuaicheKeywordAddV2RequestData       `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheKeywordAddV2Response struct {
	Responce  *KuaicheKeywordAddV2Responce `json:"jingdong_ads_dsp_rtb_keyword_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_keyword_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheKeywordAddV2Responce struct {
	Data *KuaicheKeywordAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                           `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheKeywordAddV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheKeywordAddV2(req *KuaicheKeywordAddV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := keyword.NewKuaicheKeywordAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}

	var response KuaicheKeywordAddV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return false, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return false, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Success, nil
}
