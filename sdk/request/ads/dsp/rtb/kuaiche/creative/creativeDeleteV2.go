package creative

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCreativeDeleteV2Request struct {
	Request *sdk.Request
}

type KuaicheCreativeDeleteV2RequestData struct {
	AdIds []uint64 `json:"adIds"` // 创意id列表
}

// create new request
func NewKuaicheCreativeDeleteV2Request() (req *KuaicheCreativeDeleteV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.tp.batchDeleteAd.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCreativeDeleteV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCreativeDeleteV2Request) SetData(data *KuaicheCreativeDeleteV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCreativeDeleteV2Request) GetData() *KuaicheCreativeDeleteV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCreativeDeleteV2RequestData)
	}
	return nil
}

func (req *KuaicheCreativeDeleteV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCreativeDeleteV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
