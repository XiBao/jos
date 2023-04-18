package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignDeleteV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignDeleteV2RequestData struct {
	Ids []uint64 `json:"ids"` // 计划id集合
}

// create new request
func NewKuaicheCampaignDeleteV2Request() (req *KuaicheCampaignDeleteV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.delete.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignDeleteV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignDeleteV2Request) SetData(data *KuaicheCampaignDeleteV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignDeleteV2Request) GetData() *KuaicheCampaignDeleteV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignDeleteV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignDeleteV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignDeleteV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
