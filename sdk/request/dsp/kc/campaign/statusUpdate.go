package campaign

import "github.com/XiBao/jos/sdk"

type CampainStatusUpdateRequest struct {
	Request *sdk.Request
}

func NewCampainStatusUpdateRequest() (req *CampainStatusUpdateRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campain.status.update", Params: make(map[string]interface{}, 2)}
	req = &CampainStatusUpdateRequest{
		Request: &request,
	}
	return
}

func (req *CampainStatusUpdateRequest) SetStatus(status int) {
	req.Request.Params["status"] = status
}

func (req *CampainStatusUpdateRequest) GetStatus() int {
	status, found := req.Request.Params["status"]
	if found {
		return status.(int)
	}
	return 0
}

func (req *CampainStatusUpdateRequest) SetCompaignId(compaignId string) {
	req.Request.Params["compaignId"] = compaignId
}

func (req *CampainStatusUpdateRequest) GetCompaignId() string {
	compaignId, found := req.Request.Params["compaignId"]
	if found {
		return compaignId.(string)
	}
	return ""
}
