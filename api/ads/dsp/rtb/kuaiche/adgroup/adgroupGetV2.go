package adgroup

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 获取快车单元信息
type KuaicheAdgroupGetV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupGetV2RequestData       `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupGetV2Response struct {
	Responce  *KuaicheAdgroupGetV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_group_get_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_group_get_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheAdgroupGetV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheAdgroupGetV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheAdgroupGetV2Responce struct {
	Data *KuaicheAdgroupGetV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                           `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheAdgroupGetV2Responce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheAdgroupGetV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheAdgroupGetV2ResponseData struct {
	Data *KuaicheAdgroupGetV2ResponseDataAdgroup `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheAdgroupGetV2ResponseDataAdgroup struct {
	Adgroup *dsp.Adgroup `json:"adGroup,omitempty" codec:"adGroup,omitempty"`
}

func KuaicheAdgroupGetV2(req *KuaicheAdgroupGetV2Request) (*dsp.Adgroup, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupGetV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheAdgroupGetV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data.Adgroup, nil
}
