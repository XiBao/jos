package dmp

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDmpUpdateStatusV2Request struct {
	Request *sdk.Request
}

type KuaicheDmpUpdateStatusV2RequestData struct {
	Ids         []uint64 `json:"ids"`         // 计划id集合
	AdgroupId   uint64   `json:"adGroupId"`   // 单元id
	OperateType uint     `json:"OperateType"` // 状态值 1暂停；2启动
}

// create new request
func NewKuaicheDmpUpdateStatusV2Request() (req *KuaicheDmpUpdateStatusV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.dmp.updateStatus", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDmpUpdateStatusV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpUpdateStatusV2Request) SetData(data *KuaicheDmpUpdateStatusV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheDmpUpdateStatusV2Request) GetData() *KuaicheDmpUpdateStatusV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheDmpUpdateStatusV2RequestData)
	}
	return nil
}

func (req *KuaicheDmpUpdateStatusV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheDmpUpdateStatusV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
