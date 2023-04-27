package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordUpdateTypeV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordUpdateTypeV2RequestData struct {
	Ids  []uint64 `json:"ids"`  // 关键词id列表
	Type uint     `json:"type"` // 关键词匹配类型 8 切词 4 短语 1精确
}

// create new request
func NewKuaicheKeywordUpdateTypeV2Request() (req *KuaicheKeywordUpdateTypeV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.type.update.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordUpdateTypeV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordUpdateTypeV2Request) SetData(data *KuaicheKeywordUpdateTypeV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordUpdateTypeV2Request) GetData() *KuaicheKeywordUpdateTypeV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordUpdateTypeV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordUpdateTypeV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordUpdateTypeV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
