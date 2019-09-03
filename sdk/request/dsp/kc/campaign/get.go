package campaign

import "github.com/XiBao/jos/sdk"

type CampainGetRequest struct {
	Request *sdk.Request
}

func NewCampainGetRequest() (req *CampainGetRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campain.get", Params: make(map[string]interface{}, 1)}
	req = &CampainGetRequest{
		Request: &request,
	}
	return

}

func (req *CampainGetRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaignId"] = campaignId
}

func (req *CampainGetRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaignId"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}
