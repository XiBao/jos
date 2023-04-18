package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignAddV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignAddV2RequestData struct {
	PutType      uint   `json:"putType"`             // 投放类型 3商品推广 4 活动推广（腰带店铺为4）
	StartTime    string `json:"startTime"`           // 计划开始时间
	DayBudget    uint   `json:"dayBudget,omitempty"` // 统一日预算，范围10至9999999，不填写则为不限
	CampaignType uint   `json:"campaignType"`        // 计划类型 2普通快车 18快车腰带店铺
	Name         string `json:"name"`                // 计划名
	EndTime      string `json:"endTime,omitempty"`   // 计划结束时间
}

// create new request
func NewKuaicheCampaignAddV2Request() (req *KuaicheCampaignAddV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.add.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignAddV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignAddV2Request) SetData(data *KuaicheCampaignAddV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignAddV2Request) GetData() *KuaicheCampaignAddV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignAddV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignAddV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignAddV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
