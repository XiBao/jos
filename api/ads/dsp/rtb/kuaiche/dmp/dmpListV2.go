package dmp

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 获取快车搜索人群管理列表
type KuaicheDmpListV2Request struct {
	api.BaseRequest
	Data   *dmp.KuaicheDmpListV2RequestData              `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheDmpListV2Response struct {
	Responce  *KuaicheDmpListV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_dmplist_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_dmplist_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpListV2Responce struct {
	Data *KuaicheDmpListV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                        `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpListV2Responce) IsError() bool {
	return r.Data != nil || r.Data.IsError()
}

func (r KuaicheDmpListV2Responce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpListV2ResponseData struct {
	Data *KuaicheDmpListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheDmpListV2ResponseDataData struct {
	DmpList   []dsp.DmpData   `json:"datas,omitempty" codec:"datas,omitempty"`
	Ext       *dsp.DmpExtData `json:"ext,omitempty" codec:"ext,omitempty"`
	Paginator *dsp.Paginator  `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheDmpListV2(req *KuaicheDmpListV2Request) (*KuaicheDmpListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheDmpListV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
