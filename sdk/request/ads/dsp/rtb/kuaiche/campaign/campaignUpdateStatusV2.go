package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignUpdateStatusV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignUpdateStatusV2RequestData struct {
	Ids         []uint64 `json:"ids"`         // 计划id集合
	OperateType uint     `json:"OperateType"` // 状态值 1暂停；2启动
}

// create new request
func NewKuaicheCampaignUpdateStatusV2Request() (req *KuaicheCampaignUpdateStatusV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.updatestatus.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignUpdateStatusV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignUpdateStatusV2Request) SetData(data *KuaicheCampaignUpdateStatusV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignUpdateStatusV2Request) GetData() *KuaicheCampaignUpdateStatusV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignUpdateStatusV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignUpdateStatusV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignUpdateStatusV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
