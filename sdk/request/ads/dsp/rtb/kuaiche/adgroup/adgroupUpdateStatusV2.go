package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupUpdateStatusV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupUpdateStatusV2RequestData struct {
	Ids         []uint64 `json:"ids"`         // 计划id集合
	OperateType uint     `json:"OperateType"` // 状态值 1暂停；2启动
}

// create new request
func NewKuaicheAdgroupUpdateStatusV2Request() (req *KuaicheAdgroupUpdateStatusV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.updatestatus.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupUpdateStatusV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupUpdateStatusV2Request) SetData(data *KuaicheAdgroupUpdateStatusV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupUpdateStatusV2Request) GetData() *KuaicheAdgroupUpdateStatusV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupUpdateStatusV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupUpdateStatusV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupUpdateStatusV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
