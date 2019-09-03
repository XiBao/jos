package campaign

import "github.com/XiBao/jos/sdk"

type CampainDeleteRequest struct {
	Request *sdk.Request
}

func NewCampainDeleteRequest() (req *CampainDeleteRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campain.delete", Params: make(map[string]interface{}, 1)}
	req = &CampainDeleteRequest{
		Request: &request,
	}
	return

}

func (req *CampainDeleteRequest) SetCompaignId(compaignId string) {
	req.Request.Params["compaignId"] = compaignId
}

func (req *CampainDeleteRequest) GetCompaignId() string {
	compaignId, found := req.Request.Params["compaignId"]
	if found {
		return compaignId.(string)
	}
	return ""
}
