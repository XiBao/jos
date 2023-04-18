package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignUpdateBudgetV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignUpdateBudgetV2RequestData struct {
	Id        uint64 `json:"id"`                  // 计划id
	DateRange string `json:"dateRange,omitempty"` // 计划自定义预算
	DayBudget uint   `json:"dayBudget,omitempty"` // 计划日预算
}

// create new request
func NewKuaicheCampaignUpdateBudgetV2Request() (req *KuaicheCampaignUpdateBudgetV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.updatebudget.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignUpdateBudgetV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignUpdateBudgetV2Request) SetData(data *KuaicheCampaignUpdateBudgetV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignUpdateBudgetV2Request) GetData() *KuaicheCampaignUpdateBudgetV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignUpdateBudgetV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignUpdateBudgetV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignUpdateBudgetV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
