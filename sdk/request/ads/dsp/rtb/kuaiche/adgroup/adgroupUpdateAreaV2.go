package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupUpdateAreaV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupUpdateAreaV2RequestData struct {
	Ids        []uint64 `json:"ids"`        // 单元id集合
	NewAreaIds []string `json:"newAreaIds"` // 区域码集合
}

// create new request
func NewKuaicheAdgroupUpdateAreaV2Request() (req *KuaicheAdgroupUpdateAreaV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.updateArea.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupUpdateAreaV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupUpdateAreaV2Request) SetData(data *KuaicheAdgroupUpdateAreaV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupUpdateAreaV2Request) GetData() *KuaicheAdgroupUpdateAreaV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupUpdateAreaV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupUpdateAreaV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupUpdateAreaV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
