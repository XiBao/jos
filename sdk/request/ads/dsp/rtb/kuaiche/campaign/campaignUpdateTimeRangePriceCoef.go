package campaign

import (
	"github.com/XiBao/jos/sdk"
)

type KuaicheCampaignUpdateTimeRangePriceCoefRequest struct {
	Request *sdk.Request
}

type KuaicheCampaignUpdateTimeRangePriceCoefRequestData struct {
	Id        uint64 `json:"id"`        // 计划id
	DateRange string `json:"dateRange"` // 计划自定义预算
	DayBudget uint   `json:"dayBudget"` // 计划日预算
}

// create new request
func NewKuaicheCampaignUpdateTimeRangePriceCoefRequest() (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.updateTimeRangePriceCoef", Params: make(map[string]interface{}, 5)}
	req = &KuaicheCampaignUpdateTimeRangePriceCoefRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) GetId() uint64 {
	id, found := req.Request.Params["id"]
	if found {
		return id.(uint64)
	}
	return 0
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) SetTimeRangePriceCoef(timeRangePriceCoef string) {
	req.Request.Params["timeRangePriceCoef"] = timeRangePriceCoef
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) GetTimeRangePriceCoef() string {
	timeRangePriceCoef, found := req.Request.Params["timeRangePriceCoef"]
	if found {
		return timeRangePriceCoef.(string)
	}
	return ""
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) SetAccessPin(accessPin string) {
	req.Request.Params["accessPin"] = accessPin
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) GetAccessPin() string {
	accessPin, found := req.Request.Params["accessPin"]
	if found {
		return accessPin.(string)
	}
	return ""
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) SetAuthType(authType string) {
	req.Request.Params["authType"] = authType
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) GetAuthType() string {
	authType, found := req.Request.Params["authType"]
	if found {
		return authType.(string)
	}
	return ""
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) SetPlatformBusinessType(platformBusinessType string) {
	req.Request.Params["platformBusinessType"] = platformBusinessType
}

func (req *KuaicheCampaignUpdateTimeRangePriceCoefRequest) GetPlatformBusinessType() string {
	platformBusinessType, found := req.Request.Params["platformBusinessType"]
	if found {
		return platformBusinessType.(string)
	}
	return ""
}
