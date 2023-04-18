package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignUpdateNameV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignUpdateNameV2RequestData struct {
	Id   uint64 `json:"id"`   // 快车计划id
	Name string `json:"name"` // 计划名称
}

// create new request
func NewKuaicheCampaignUpdateNameV2Request() (req *KuaicheCampaignUpdateNameV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.updatename.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignUpdateNameV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignUpdateNameV2Request) SetData(data *KuaicheCampaignUpdateNameV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignUpdateNameV2Request) GetData() *KuaicheCampaignUpdateNameV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignUpdateNameV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignUpdateNameV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignUpdateNameV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
