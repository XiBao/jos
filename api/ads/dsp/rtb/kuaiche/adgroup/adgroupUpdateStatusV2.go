package adgroup

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 修改快车单元状态
type KuaicheAdgroupUpdateStatusV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupUpdateStatusV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt    `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupUpdateStatusV2Response struct {
	Responce  *KuaicheAdgroupUpdateStatusV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_group_updatestatus_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_group_updatestatus_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                   `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheAdgroupUpdateStatusV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheAdgroupUpdateStatusV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheAdgroupUpdateStatusV2Responce struct {
	Data *KuaicheAdgroupUpdateStatusV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                    `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheAdgroupUpdateStatusV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheAdgroupUpdateStatusV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheAdgroupUpdateStatusV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheAdgroupUpdateStatusV2(req *KuaicheAdgroupUpdateStatusV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupUpdateStatusV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheAdgroupUpdateStatusV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.Data.Success, nil
}
