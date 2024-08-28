package campaign

import (
	"github.com/XiBao/jos/sdk"
)

type CampainShopAddRequest struct {
	Request *sdk.Request
}

func NewCampainShopAddRequest() (req *CampainShopAddRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campainShop.add", Params: make(map[string]interface{}, 2)}
	req = &CampainShopAddRequest{
		Request: &request,
	}
	return
}

func (req *CampainShopAddRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *CampainShopAddRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *CampainShopAddRequest) SetDayBudget(dayBudget int) {
	req.Request.Params["dayBudget"] = dayBudget
}

func (req *CampainShopAddRequest) GetDayBudget() int {
	dayBudget, found := req.Request.Params["dayBudget"]
	if found {
		return dayBudget.(int)
	}
	return 0
}
