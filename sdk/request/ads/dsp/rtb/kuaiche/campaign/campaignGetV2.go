package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignGetV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignGetV2RequestData struct {
	Id uint64 `json:"id"` // 计划id
}

// create new request
func NewKuaicheCampaignGetV2Request() (req *KuaicheCampaignGetV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.get.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignGetV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignGetV2Request) SetData(data *KuaicheCampaignGetV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignGetV2Request) GetData() *KuaicheCampaignGetV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignGetV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignGetV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignGetV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
