package area

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/area"
)

// 获取地域字典
type AreaQueryV2Request struct {
	api.BaseRequest
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type AreaQueryV2Response struct {
	Responce  *AreaQueryV2Responce `json:"jingdong_ads_dsp_rtb_area_query_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_area_query_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type AreaQueryV2Responce struct {
	Data *AreaQueryV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                   `json:"code,omitempty" codec:"code,omitempty"`
}

type AreaQueryV2ResponseData struct {
	Data string `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func AreaQueryV2(req *AreaQueryV2Request) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := area.NewAreaQueryV2Request()
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("no result.")
	}

	var response AreaQueryV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return "", errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return "", errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Data, nil
}
