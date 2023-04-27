package creative

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCreativeUpdateStatusV2Request struct {
	Request *sdk.Request
}

type KuaicheCreativeUpdateStatusV2RequestData struct {
	Ids         []uint64 `json:"ids"`         // 计划id集合
	OperateType uint     `json:"OperateType"` // 状态值 1暂停；2启动
}

// create new request
func NewKuaicheCreativeUpdateStatusV2Request() (req *KuaicheCreativeUpdateStatusV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.featured.updateAdStatus.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCreativeUpdateStatusV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCreativeUpdateStatusV2Request) SetData(data *KuaicheCreativeUpdateStatusV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCreativeUpdateStatusV2Request) GetData() *KuaicheCreativeUpdateStatusV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCreativeUpdateStatusV2RequestData)
	}
	return nil
}

func (req *KuaicheCreativeUpdateStatusV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCreativeUpdateStatusV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
