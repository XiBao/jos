package campaign

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type AddRequest struct {
	api.BaseRequest
	Name      string `json:"name,omitempty" codec:"name,omitempty"`             //计划名称
	DayBudget int    `json:"day_budget,omitempty" codec:"day_budget,omitempty"` //预算
}

type AddResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddData            `json:"jingdong_dsp_kc_campainShop_add_responce,omitempty" codec:"jingdong_dsp_kc_campainShop_add_responce,omitempty"`
}
type AddData struct {
	Result AddResult `json:"addcampain_result,omitempty" codec:"addcampain_result,omitempty"`
}

type AddResult struct {
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

// 新建计划
func Add(req *AddRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainShopAddRequest()

	r.SetName(req.Name)
	r.SetDayBudget(req.DayBudget)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}

	var response AddResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return false, errors.New(response.Data.Result.ErrorMsg)
	}

	return true, nil
}
