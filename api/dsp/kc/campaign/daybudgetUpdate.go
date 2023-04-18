package campaign

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/campaign"
)

type DaybudgetUpdateRequest struct {
	api.BaseRequest
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"` //计划id
	DayBudget  uint64 `json:"day_budget,omitempty" codec:"day_budget,omitempty"` //预算
}

type DaybudgetUpdateResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *DaybudgetUpdateData `json:"jingdong_dsp_kc_campain_daybudget_update_responce,omitempty" codec:"jingdong_dsp_kc_campain_daybudget_update_responce,omitempty"`
}

type DaybudgetUpdateData struct {
	Result *DaybudgetUpdateResult `json:"updatecampaigndaybudget_result"`
}

type DaybudgetUpdateResult struct {
	CampaignId int    `json:"campaignId,omitempty" codec:"campaignId,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
}

// 修改计划日限额
func DaybudgetUpdate(req *DaybudgetUpdateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewCampainDaybudgetUpdateRequest()

	r.SetCampaignId(req.CampaignId)
	r.SetDayBudget(req.DayBudget)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response DaybudgetUpdateResponse
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
