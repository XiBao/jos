package creative

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	creative "github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/creative"
)

// 新增，修改与删除快车创意批量接口
type KuaicheCreativeAddV2Request struct {
	api.BaseRequest
	Data   *creative.KuaicheCreativeAddV2RequestData     `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCreativeAddV2Response struct {
	Responce  *KuaicheCreativeAddV2Responce `json:"jingdong_ads_dsp_rtb_kc_ad_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kc_ad_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheCreativeAddV2Responce struct {
	Data *KuaicheCreativeAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheCreativeAddV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCreativeAddV2(req *KuaicheCreativeAddV2Request) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := creative.NewKuaicheCreativeAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response KuaicheCreativeAddV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return 0, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return 0, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Data, nil
}
