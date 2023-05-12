package dmp

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDmpUpdatePriceRequest struct {
	Request *sdk.Request
}

type KuaicheDmpUpdatePriceRequestDmpVOS struct {
	Ids          []uint64 `json:"ids"`          // 单元下人群id集合
	AdgroupId    uint64   `json:"adGroupId"`    // 单元人群溢价系数
	AdGroupPrice uint     `json:"adGroupPrice"` // 单元人群溢价系数
}

// create new request
func NewKuaicheDmpUpdatePriceRequest() (req *KuaicheDmpUpdatePriceRequest) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.batchUpdateDifferentDmpPrice", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDmpUpdatePriceRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpUpdatePriceRequest) SetDmpVOS(dmpVOS []KuaicheDmpUpdatePriceRequestDmpVOS) {
	req.Request.Params["dmpVOS"] = dmpVOS
}

func (req *KuaicheDmpUpdatePriceRequest) GetDmpVOS() []KuaicheDmpUpdatePriceRequestDmpVOS {
	dmpVOS, found := req.Request.Params["dmpVOS"]
	if found {
		return dmpVOS.([]KuaicheDmpUpdatePriceRequestDmpVOS)
	}
	return nil
}

func (req *KuaicheDmpUpdatePriceRequest) SetParamExt(paramExt *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["paramExt"] = paramExt
}

func (req *KuaicheDmpUpdatePriceRequest) GetParamExt() *dsp.JdDspPlatformGatewayApiVoParamExt {
	paramExt, found := req.Request.Params["paramExt"]
	if found {
		return paramExt.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
