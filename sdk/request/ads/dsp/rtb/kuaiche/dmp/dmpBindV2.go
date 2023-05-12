package dmp

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDmpBindV2Request struct {
	Request *sdk.Request
}

type KuaicheDmpBindV2RequestCrowdVOS struct {
	AdGroupPrice uint   `json:"adGroupPrice"` // 人群溢价
	CrowdId      uint64 `json:"crowdId"`      // 人群id,635531(智能定向人群)
	IsUsed       uint   `json:"isUsed"`       // 是否启用 1启用 0 不启用
}

type KuaicheDmpBindV2RequestData struct {
	DmpType  uint                              `json:"dmpType"`  // 1：推荐人群 ，2：搜索人群，不传默认2
	GroupId  uint64                            `json:"groupId"`  // 单元id
	CrowdVOS []KuaicheDmpBindV2RequestCrowdVOS `json:"crowdVOS"` // 人群列表
}

// create new request
func NewKuaicheDmpBindV2Request() (req *KuaicheDmpBindV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.dmp.bind.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDmpBindV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheDmpBindV2Request) SetData(data *KuaicheDmpBindV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheDmpBindV2Request) GetData() *KuaicheDmpBindV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheDmpBindV2RequestData)
	}
	return nil
}

func (req *KuaicheDmpBindV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheDmpBindV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
