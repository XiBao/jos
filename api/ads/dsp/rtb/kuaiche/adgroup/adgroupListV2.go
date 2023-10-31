package adgroup

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 查询单元列表
type KuaicheAdgroupListV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupListV2RequestData      `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupListV2Response struct {
	Responce  *KuaicheAdgroupListV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_group_list_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_group_list_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheAdgroupListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheAdgroupListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheAdgroupListV2Responce struct {
	Data *KuaicheAdgroupListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                            `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheAdgroupListV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheAdgroupListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheAdgroupListV2ResponseData struct {
	Data *KuaicheAdgroupListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheAdgroupListV2ResponseDataData struct {
	Campaigns []dsp.AdgroupData `json:"datas,omitempty" codec:"datas,omitempty"`
	Paginator *dsp.Paginator    `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheAdgroupListV2(req *KuaicheAdgroupListV2Request) (*KuaicheAdgroupListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheAdgroupListV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
