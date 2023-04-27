package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupUpdateNameV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupUpdateNameV2RequestData struct {
	Id   uint64 `json:"id"`   // 快车单元id
	Name string `json:"name"` // 单元名称
}

// create new request
func NewKuaicheAdgroupUpdateNameV2Request() (req *KuaicheAdgroupUpdateNameV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.updatename.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupUpdateNameV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupUpdateNameV2Request) SetData(data *KuaicheAdgroupUpdateNameV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupUpdateNameV2Request) GetData() *KuaicheAdgroupUpdateNameV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupUpdateNameV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupUpdateNameV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupUpdateNameV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
