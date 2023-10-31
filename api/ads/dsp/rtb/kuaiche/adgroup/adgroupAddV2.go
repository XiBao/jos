package adgroup

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	adgroup "github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 新建快车商品，店铺类型单元
type KuaicheAdgroupAddV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupAddV2RequestData       `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupAddV2Response struct {
	Responce  *KuaicheAdgroupAddV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_productgroup_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_productgroup_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheAdgroupAddV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheAdgroupAddV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheAdgroupAddV2Responce struct {
	Data *KuaicheAdgroupAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                           `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheAdgroupAddV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheAdgroupAddV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheAdgroupAddV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheAdgroupAddV2(req *KuaicheAdgroupAddV2Request) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheAdgroupAddV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Responce.Data.Data, nil
}
