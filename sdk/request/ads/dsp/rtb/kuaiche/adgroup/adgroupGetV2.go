package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupGetV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupGetV2RequestData struct {
	Id uint64 `json:"id"` // 计划id
}

// create new request
func NewKuaicheAdgroupGetV2Request() (req *KuaicheAdgroupGetV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.get.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupGetV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupGetV2Request) SetData(data *KuaicheAdgroupGetV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupGetV2Request) GetData() *KuaicheAdgroupGetV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupGetV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupGetV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupGetV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
