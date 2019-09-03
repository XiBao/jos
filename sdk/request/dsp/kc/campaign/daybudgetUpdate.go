package campaign

import "github.com/XiBao/jos/sdk"

type CampainDaybudgetUpdateRequest struct {
	Request *sdk.Request
}

func NewCampainDaybudgetUpdateRequest() (req *CampainDaybudgetUpdateRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campain.daybudget.update", Params: make(map[string]interface{}, 2)}
	req = &CampainDaybudgetUpdateRequest{
		Request: &request,
	}
	return
}

func (req *CampainDaybudgetUpdateRequest) SetCampaignId(CampaignId int) {
	req.Request.Params["campaignId"] = CampaignId
}

func (req *CampainDaybudgetUpdateRequest) GetCampaignId() int {
	CampaignId, found := req.Request.Params["campaignId"]
	if found {
		return CampaignId.(int)
	}
	return 0
}

func (req *CampainDaybudgetUpdateRequest) SetDayBudget(DayBudget uint64) {
	req.Request.Params["dayBudget"] = DayBudget
}

func (req *CampainDaybudgetUpdateRequest) GetDayBudget() uint64 {
	DayBudget, found := req.Request.Params["dayBudget"]
	if found {
		return DayBudget.(uint64)
	}
	return 0
}
