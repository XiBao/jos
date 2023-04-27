package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordDeleteV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordDeleteV2RequestData struct {
	Ids       []uint64 `json:"ids"`       // 关键词id的集合
	AdgroupId uint64   `json:"adGroupId"` // 单元id
}

// create new request
func NewKuaicheKeywordDeleteV2Request() (req *KuaicheKeywordDeleteV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.delete.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordDeleteV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordDeleteV2Request) SetData(data *KuaicheKeywordDeleteV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordDeleteV2Request) GetData() *KuaicheKeywordDeleteV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordDeleteV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordDeleteV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordDeleteV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
