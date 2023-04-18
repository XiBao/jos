package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignUpdateDateV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignUpdateDateV2RequestData struct {
	Id        uint64 `json:"id"`                // 计划id
	StartTime string `json:"startTime"`         // 计划开始时间
	EndTime   string `json:"endTime,omitempty"` // 计划结束时间
}

// create new request
func NewKuaicheCampaignUpdateDateV2Request() (req *KuaicheCampaignUpdateDateV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.updatedate.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignUpdateDateV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignUpdateDateV2Request) SetData(data *KuaicheCampaignUpdateDateV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignUpdateDateV2Request) GetData() *KuaicheCampaignUpdateDateV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignUpdateDateV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignUpdateDateV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignUpdateDateV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
