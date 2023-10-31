package area

import (
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

func (r AreaQueryV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r AreaQueryV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no responce"
}

type AreaQueryV2Responce struct {
	Data *AreaQueryV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                   `json:"code,omitempty" codec:"code,omitempty"`
}

func (r AreaQueryV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r AreaQueryV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "unexpected error"
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

	var response AreaQueryV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Responce.Data.Data, nil
}
