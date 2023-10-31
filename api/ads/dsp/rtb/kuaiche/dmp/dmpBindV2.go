package dmp

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 快车单元绑定、修改搜索位人群信息
type KuaicheDmpBindV2Request struct {
	api.BaseRequest
	Data   *dmp.KuaicheDmpBindV2RequestData              `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheDmpBindV2Response struct {
	Responce  *KuaicheDmpBindV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_dmp_bind_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_dmp_bind_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpBindV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpBindV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpBindV2Responce struct {
	Data *KuaicheDmpBindV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                        `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpBindV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheDmpBindV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpBindV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDmpBindV2(req *KuaicheDmpBindV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpBindV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheDmpBindV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
